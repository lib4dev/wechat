package qrcode

import (
	"github.com/lib4dev/wechat/mp"
	"github.com/lib4dev/wechat/mp/base"
)

// ShortURL 将一条长链接转成短链接.
func ShortURL(clt *mp.Context, longURL string) (shortURL string, err error) {
	return base.ShortURL(clt, longURL)
}
