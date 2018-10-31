package mch

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/wechat/util"
)

type PayConf struct {
	AppId  string
	MchId  string
	ApiKey string

	SubAppId string
	SubMchId string
}

type NoitfyServer struct {
	PayConf
	handler Handler
}

func NewNotifyByConf(conf PayConf, handler HandlerFunc) *NoitfyServer {
	if conf.ApiKey == "" {
		panic("empty apiKey")
	}
	if handler == nil {
		panic("nil Handler")
	}

	return &NoitfyServer{
		PayConf: conf,
		handler: handler,
	}
}

// NewNoitfyServer 创建一个新的 Server.
//  appId:        可选; 公众号的 appid, 如果设置了值则该 Server 只能处理 appid 为该值的消息(事件)
//  mchId:        可选; 商户号 mch_id, 如果设置了值则该 Server 只能处理 mch_id 为该值的消息(事件)
//  apiKey:       必选; 商户的签名 key
//  handler:      必选; 处理微信服务器推送过来的消息(事件)的 Handler
//  errorHandler: 可选; 用于处理 Server 在处理消息(事件)过程中产生的错误, 如果没有设置则默认使用 DefaultErrorHandler

func NewNoitfyServerHandler(appId, mchId, apiKey string, handler HandlerFunc) func(container component.IContainer) *NoitfyServer {
	return func(container component.IContainer) (u *NoitfyServer) {
		return NewNotifyByConf(PayConf{AppId: appId, MchId: mchId, ApiKey: apiKey}, handler)
	}
}

// NewSubMchServer 创建一个新的 Server.
func NewSubMchServer(appId, mchId, apiKey string, subAppId, subMchId string, handler HandlerFunc) *NoitfyServer {
	return NewNotifyByConf(PayConf{AppId: appId, MchId: mchId, ApiKey: apiKey, SubAppId: subAppId, SubMchId: subMchId}, handler)
}

// Handle 处理微信服务器的回调请求, query 参数可以为 nil.
func (srv *NoitfyServer) Handle(ctx *context.Context) (r interface{}) {
	switch strings.ToUpper(ctx.Request.GetMethod()) {
	case "POST":
		body, err := ctx.Request.GetBody()
		if err != nil {
			return err
		}
		msg, err := util.DecodeXMLToMap(strings.NewReader(body))
		if err != nil {
			return err
		}

		returnCode := msg["return_code"]
		if returnCode != "" && returnCode != ReturnCodeSuccess {
			err = &Error{
				ReturnCode: returnCode,
				ReturnMsg:  msg["return_msg"],
			}
			return err
		}

		resultCode := msg["result_code"]
		if resultCode != "" && resultCode != ResultCodeSuccess {
			err = &BizError{
				ResultCode:  resultCode,
				ErrCode:     msg["err_code"],
				ErrCodeDesc: msg["err_code_des"],
			}
			return err
		}

		if srv.AppId != "" {
			wantAppId := srv.AppId
			haveAppId := msg["appid"]
			if haveAppId != "" && !util.SecureCompareString(haveAppId, wantAppId) {
				err = fmt.Errorf("appid mismatch, have: %s, want: %s", haveAppId, wantAppId)
				return err
			}
		}
		if srv.MchId != "" {
			wantMchId := srv.MchId
			haveMchId := msg["mch_id"]
			if haveMchId != "" && !util.SecureCompareString(haveMchId, wantMchId) {
				err = fmt.Errorf("mch_id mismatch, have: %s, want: %s", haveMchId, wantMchId)
				return err
			}
		}

		if srv.SubAppId != "" {
			wantSubAppId := srv.SubAppId
			haveSubAppId := msg["sub_appid"]
			if haveSubAppId != "" && !util.SecureCompareString(haveSubAppId, wantSubAppId) {
				err = fmt.Errorf("sub_appid mismatch, have: %s, want: %s", haveSubAppId, wantSubAppId)
				return err
			}
		}
		if srv.SubMchId != "" {
			wantSubMchId := srv.SubMchId
			haveSubMchId := msg["sub_mch_id"]
			if haveSubMchId != "" && !util.SecureCompareString(haveSubMchId, wantSubMchId) {
				err = fmt.Errorf("sub_mch_id mismatch, have: %s, want: %s", haveSubMchId, wantSubMchId)
				return err
			}
		}

		// 认证签名
		if haveSignature := msg["sign"]; haveSignature != "" {
			var wantSignature string
			switch signType := msg["sign_type"]; signType {
			case "", SignType_MD5:
				wantSignature = Sign2(msg, srv.ApiKey, md5.New())
			case SignType_HMAC_SHA256:
				wantSignature = Sign2(msg, srv.ApiKey, hmac.New(sha256.New, []byte(srv.ApiKey)))
			default:
				err = fmt.Errorf("unsupported notification sign_type: %s", signType)
				return err
			}
			if !util.SecureCompareString(haveSignature, wantSignature) {
				err = fmt.Errorf("sign mismatch,\nhave: %s,\nwant: %s", haveSignature, wantSignature)
				return err
			}
		} else {
			if _, ok := msg["req_info"]; !ok { // 退款结果通知没有 sign 字段
				return ErrNotFoundSign
			}
		}
		srv.handler.ServeMsg(&srv.PayConf, msg, ctx)
	default:
	}
	return nil
}
