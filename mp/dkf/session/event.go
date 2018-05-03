package session

import (
	"github.com/micro-plat/wechat/mp"
)

const (
	EventTypeKfCreateSession mp.EventType = "kf_create_session" // 接入会话
	EventTypeKfCloseSession  mp.EventType = "kf_close_session"  // 关闭会话
	EventTypeKfSwitchSession mp.EventType = "kf_switch_session" // 转接会话
)

type KfCreateSessionEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.MsgHeader
	EventType mp.EventType `xml:"Event"     json:"Event"`
	KfAccount string       `xml:"KfAccount" json:"KfAccount"`
}

func GetKfCreateSessionEvent(msg *mp.MixedMsg) *KfCreateSessionEvent {
	return &KfCreateSessionEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		KfAccount: msg.KfAccount,
	}
}

type KfCloseSessionEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.MsgHeader
	EventType mp.EventType `xml:"Event"     json:"Event"`
	KfAccount string       `xml:"KfAccount" json:"KfAccount"`
}

func GetKfCloseSessionEvent(msg *mp.MixedMsg) *KfCloseSessionEvent {
	return &KfCloseSessionEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		KfAccount: msg.KfAccount,
	}
}

type KfSwitchSessionEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.MsgHeader
	EventType     mp.EventType `xml:"Event"         json:"Event"`
	FromKfAccount string       `xml:"FromKfAccount" json:"FromKfAccount"`
	ToKfAccount   string       `xml:"ToKfAccount"   json:"ToKfAccount"`
}

func GetKfSwitchSessionEvent(msg *mp.MixedMsg) *KfSwitchSessionEvent {
	return &KfSwitchSessionEvent{
		MsgHeader:     msg.MsgHeader,
		EventType:     msg.EventType,
		FromKfAccount: msg.FromKfAccount,
		ToKfAccount:   msg.ToKfAccount,
	}
}
