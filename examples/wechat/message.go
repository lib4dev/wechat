package main

import (
	"github.com/micro-plat/hydra/context"

	"github.com/micro-plat/wechat"
	"github.com/silenceper/wechat/message"
)

//处理微信消息
func recvMessage(context *wechat.WContext, ctx *context.Context, msg *message.MixMessage) *message.Reply {
	ctx.Log.Info("-----recv.message-----")
	ctx.Log.Info(string(msg.Raw))
	return nil
}
