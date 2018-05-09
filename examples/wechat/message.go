package main

import (
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/wechat/mp/message"

	"github.com/micro-plat/wechat/mp"
)

//处理微信消息
func recvMessage(cnf *mp.WConf, msg *mp.MixedMsg, ctx *context.Context) *mp.Reply {
	switch msg.EventType {
	case mp.EventSubscribe:
		return &mp.Reply{MsgType: mp.MsgTypeText, MsgData: message.NewText("欢迎关注公众号")}
	}
	return nil
}
