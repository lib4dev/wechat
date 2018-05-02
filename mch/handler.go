package mch

import "github.com/micro-plat/hydra/context"

type Handler interface {
	ServeMsg(*PayConf, map[string]string, *context.Context)
}

// HandlerFunc ---------------------------------------------------------------------------------------------------------

var _ Handler = HandlerFunc(nil)

type HandlerFunc func(*PayConf, map[string]string, *context.Context)

// ServeMsg 实现 Handler 接口
func (fn HandlerFunc) ServeMsg(c *PayConf, m map[string]string, ctx *context.Context) { fn(c, m, ctx) }
