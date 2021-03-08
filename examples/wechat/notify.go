package main

import (
	"github.com/lib4dev/wechat/mch"
	"github.com/micro-plat/hydra"
)

func notifyServeHandler(conf *mch.PayConf) func(ctx hydra.IContext) interface{} {
	return func(ctx hydra.IContext) interface{} {
		return "success"
	}
}
