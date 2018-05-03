package message

import (
	"encoding/xml"

	"github.com/micro-plat/wechat/mp"
)

//Image 图片消息
type Image struct {
	CommonToken

	Image struct {
		MediaID string `xml:"MediaId"`
	} `xml:"Image"`
}

//NewImage 回复图片消息
func NewImage(mediaID string) *Image {
	image := new(Image)
	image.Image.MediaID = mediaID
	return image
}

// CommonToken 消息中通用的结构
type CommonToken struct {
	XMLName      xml.Name     `xml:"xml"`
	ToUserName   string       `xml:"ToUserName"`
	FromUserName string       `xml:"FromUserName"`
	CreateTime   int64        `xml:"CreateTime"`
	MsgType      mp.MsgType `xml:"MsgType"`
}

//SetToUserName set ToUserName
func (msg *CommonToken) SetToUserName(toUserName string) {
	msg.ToUserName = toUserName
}

//SetFromUserName set FromUserName
func (msg *CommonToken) SetFromUserName(fromUserName string) {
	msg.FromUserName = fromUserName
}

//SetCreateTime set createTime
func (msg *CommonToken) SetCreateTime(createTime int64) {
	msg.CreateTime = createTime
}

//SetMsgType set MsgType
func (msg *CommonToken) SetMsgType(msgType mp.MsgType) {
	msg.MsgType = msgType
}
