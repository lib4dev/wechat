package wechat

import (
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"runtime/debug"
	"strconv"

	"github.com/micro-plat/hydra/context"

	"github.com/silenceper/wechat/message"
	"github.com/silenceper/wechat/util"
)

//Server struct
type Server struct {
	*WContext

	messageHandler func(*WContext, *context.Context, *message.MixMessage) *message.Reply

	requestRawXMLMsg []byte

	isSafeMode bool
	random     []byte
	nonce      string
	timestamp  int64
}

//NewServer init
func NewServer(c *WContext) *Server {
	srv := new(Server)
	srv.WContext = c
	return srv
}

//NewSeverHandler init
func NewSeverHandler(c *WContext, handler func(*WContext, *context.Context, *message.MixMessage) *message.Reply) func() *Server {
	return func() *Server {
		s := NewServer(c)
		s.messageHandler = handler
		return s
	}
}

//Handle 处理微信的请求消息
func (srv *Server) Handle(name string, engine string, service string, ctx *context.Context) (r interface{}) {
	if !srv.Validate(ctx) {
		return fmt.Errorf("请求校验失败")
	}
	echostr, exists := ctx.Request.QueryString.Get("echostr")
	if exists {
		ctx.Response.ShouldContent(echostr)
		return nil
	}
	response, msg, err := srv.handleRequest(ctx)
	if err != nil {
		return err
	}

	rspMsg, err := srv.buildResponse(msg, response)
	if err != nil {
		return err
	}
	return srv.send(rspMsg, ctx)
}

//Validate 校验请求是否合法
func (srv *Server) Validate(ctx *context.Context) bool {
	timestamp := ctx.Request.GetString("timestamp")
	nonce := ctx.Request.GetString("nonce")
	signature := ctx.Request.GetString("signature")
	return signature == util.Signature(srv.Token, timestamp, nonce)
}

//HandleRequest 处理微信的请求
func (srv *Server) handleRequest(ctx *context.Context) (reply *message.Reply, mixMsg *message.MixMessage, err error) {
	//set isSafeMode
	srv.isSafeMode = false
	encryptType := ctx.Request.GetString("encrypt_type")
	if encryptType == "aes" {
		srv.isSafeMode = true
	}

	var msg interface{}
	msg, err = srv.getMessage(ctx)
	if err != nil {
		return
	}
	mixMessage, success := msg.(*message.MixMessage)
	if !success {
		err = errors.New("消息类型转换失败")
	}
	reply = srv.messageHandler(srv.WContext, ctx, mixMessage)
	return reply, mixMessage, nil
}

//getMessage 解析微信返回的消息
func (srv *Server) getMessage(ctx *context.Context) (interface{}, error) {
	var rawXMLMsgBytes []byte
	if srv.isSafeMode {
		var encryptedXMLMsg message.EncryptedXMLMsg
		body, err := ctx.Request.Ext.GetBody()
		if err != nil {
			return nil, err
		}
		if err = xml.Unmarshal([]byte(body), &encryptedXMLMsg); err != nil {
			return nil, fmt.Errorf("从body中解析xml失败,err=%v", err)
		}

		//验证消息签名
		timestamp := ctx.Request.GetString("timestamp")
		srv.timestamp, err = strconv.ParseInt(timestamp, 10, 32)
		if err != nil {
			return nil, err
		}
		nonce := ctx.Request.GetString("nonce")
		srv.nonce = nonce
		msgSignature := ctx.Request.GetString("msg_signature")
		msgSignatureGen := util.Signature(srv.Token, timestamp, nonce, encryptedXMLMsg.EncryptedMsg)
		if msgSignature != msgSignatureGen {
			return nil, fmt.Errorf("消息不合法，验证签名失败")
		}

		//解密
		srv.random, rawXMLMsgBytes, err = util.DecryptMsg(srv.AppID, encryptedXMLMsg.EncryptedMsg, srv.EncodingAESKey)
		if err != nil {
			return nil, fmt.Errorf("消息解密失败, err=%v", err)
		}
	} else {

		body, err := ctx.Request.Ext.GetBody()
		if err != nil {
			return nil, fmt.Errorf("从body中解析xml失败, err=%v", err)
		}
		rawXMLMsgBytes = []byte(body)
	}
	srv.requestRawXMLMsg = rawXMLMsgBytes

	return srv.parseRequestMessage(rawXMLMsgBytes)
}

func (srv *Server) parseRequestMessage(rawXMLMsgBytes []byte) (msg *message.MixMessage, err error) {
	msg = &message.MixMessage{}
	err = xml.Unmarshal(rawXMLMsgBytes, msg)
	msg.Raw = rawXMLMsgBytes
	return
}

//SetMessageHandler 设置用户自定义的回调方法
func (srv *Server) SetMessageHandler(handler func(*WContext, *context.Context, *message.MixMessage) *message.Reply) {
	srv.messageHandler = handler
}

func (srv *Server) buildResponse(requestMsg *message.MixMessage, reply *message.Reply) (msgData interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("panic error: %v\n%s", e, debug.Stack())
		}
	}()
	if reply == nil {
		return
	}
	msgType := reply.MsgType
	switch msgType {
	case message.MsgTypeText:
	case message.MsgTypeImage:
	case message.MsgTypeVoice:
	case message.MsgTypeVideo:
	case message.MsgTypeMusic:
	case message.MsgTypeNews:
	case message.MsgTypeTransfer:
	default:
		err = message.ErrUnsupportReply
		return
	}

	msgData = reply.MsgData
	value := reflect.ValueOf(msgData)
	//msgData must be a ptr
	kind := value.Kind().String()
	if "ptr" != kind {
		return nil, message.ErrUnsupportReply
	}

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(requestMsg.FromUserName)
	value.MethodByName("SetToUserName").Call(params)

	params[0] = reflect.ValueOf(requestMsg.ToUserName)
	value.MethodByName("SetFromUserName").Call(params)

	params[0] = reflect.ValueOf(msgType)
	value.MethodByName("SetMsgType").Call(params)

	params[0] = reflect.ValueOf(util.GetCurrTs())
	value.MethodByName("SetCreateTime").Call(params)
	return msgData, nil
}

//Send 将自定义的消息发送
func (srv *Server) send(responseMsg interface{}, ctx *context.Context) (err error) {
	replyMsg := responseMsg
	if srv.isSafeMode {
		//安全模式下对消息进行加密
		responseRawXMLMsg, err := xml.Marshal(replyMsg)
		if err != nil {
			return err
		}
		var encryptedMsg []byte
		encryptedMsg, err = util.EncryptMsg(srv.random, responseRawXMLMsg, srv.AppID, srv.EncodingAESKey)
		if err != nil {
			return err
		}
		//TODO 如果获取不到timestamp nonce 则自己生成
		timestamp := srv.timestamp
		timestampStr := strconv.FormatInt(timestamp, 10)
		msgSignature := util.Signature(srv.Token, timestampStr, srv.nonce, string(encryptedMsg))
		replyMsg = message.ResponseEncryptedXMLMsg{
			EncryptedMsg: string(encryptedMsg),
			MsgSignature: msgSignature,
			Timestamp:    timestamp,
			Nonce:        srv.nonce,
		}
	}
	if replyMsg != nil {
		ctx.Response.XML(replyMsg)
	}
	return
}
