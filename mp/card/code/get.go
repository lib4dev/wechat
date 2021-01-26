package code

import "github.com/lib4dev/wechat/mp"

// 查询code.
func Get(clt *mp.Context, id *CardItemIdentifier) (info *CardItem, err error) {
	var result struct {
		mp.Error
		CardItem
	}

	incompleteURL := "https://api.weixin.qq.com/card/code/get?access_token="
	if err = clt.PostJSON(incompleteURL, id, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	result.CardItem.Code = id.Code
	info = &result.CardItem
	return
}
