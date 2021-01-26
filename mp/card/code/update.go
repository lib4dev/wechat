package code

import "github.com/lib4dev/wechat/mp"

// 更改Code接口.
func Update(clt *mp.Context, id *CardItemIdentifier, newCode string) (err error) {
	request := struct {
		*CardItemIdentifier
		NewCode string `json:"new_code,omitempty"`
	}{
		CardItemIdentifier: id,
		NewCode:            newCode,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/card/code/update?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
