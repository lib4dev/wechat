package main

import (
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/wechat/mp/message"

	"github.com/micro-plat/wechat/mp"
	"github.com/micro-plat/wechat/mp/core"
)

//处理微信消息
func recvMessage(cnf *mp.WConf, ctx *context.Context, msg *core.MixedMsg) *core.Reply {
	ctx.Log.Info("-----recv.message-----")
	switch msg.EventType {
	case core.EventSubscribe:
		return &core.Reply{MsgType: core.MsgTypeText, MsgData: message.NewText("欢迎关注公众号")}
	}
	return nil
}
