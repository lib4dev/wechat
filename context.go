package wechat

import (
	"sync"
)

// WContext struct
type WContext struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	PayMchID       string
	PayNotifyURL   string
	PayKey         string

	//accessTokenLock 读写锁 同一个AppID一个
	accessTokenLock *sync.RWMutex

	//jsAPITicket 读写锁 同一个AppID一个
	jsAPITicketLock *sync.RWMutex
}

// SetJsAPITicketLock 设置jsAPITicket的lock
func (ctx *WContext) SetJsAPITicketLock(lock *sync.RWMutex) {
	ctx.jsAPITicketLock = lock
}

// GetJsAPITicketLock 获取jsAPITicket 的lock
func (ctx *WContext) GetJsAPITicketLock() *sync.RWMutex {
	return ctx.jsAPITicketLock
}
