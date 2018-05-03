package code

import "github.com/micro-plat/wechat/mp"

// 设置卡券失效接口.
func Unavailable(clt *mp.Context, id *CardItemIdentifier) (err error) {
	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/card/code/unavailable?access_token="
	if err = clt.PostJSON(incompleteURL, id, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
