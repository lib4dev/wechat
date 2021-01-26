package mch

import "github.com/micro-plat/hydra"

type Handler interface {
	ServeMsg(*PayConf, map[string]string, hydra.IContext)
}

// HandlerFunc ---------------------------------------------------------------------------------------------------------

var _ Handler = HandlerFunc(nil)

type HandlerFunc func(*PayConf, map[string]string, hydra.IContext)

// ServeMsg 实现 Handler 接口
func (fn HandlerFunc) ServeMsg(c *PayConf, m map[string]string, ctx hydra.IContext) { fn(c, m, ctx) }
