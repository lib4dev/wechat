package qrcode

import (
	"github.com/micro-plat/wechat/mp/base"
	"github.com/micro-plat/wechat/mp/core"
)

// ShortURL 将一条长链接转成短链接.
func ShortURL(clt *core.Context, longURL string) (shortURL string, err error) {
	return base.ShortURL(clt, longURL)
}
