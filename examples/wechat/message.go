package main

import (
	"github.com/lib4dev/wechat/mp/message"
	"github.com/micro-plat/hydra"

	"github.com/lib4dev/wechat/mp"
)

//处理微信消息
func recvMessage(cnf *mp.WConf, msg *mp.MixedMsg, ctx hydra.IContext) *mp.Reply {
	switch msg.EventType {
	case mp.EventSubscribe:
		return &mp.Reply{MsgType: mp.MsgTypeText, MsgData: message.NewText("欢迎关注公众号")}
	}
	return nil
}
