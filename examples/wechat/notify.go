package main

import (
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/wechat/mch"
)

func orderNotify(conf *mch.PayConf, msg map[string]string, ctx *context.Context) {
	ctx.Log.Info("-----recv.notify-----")
}
