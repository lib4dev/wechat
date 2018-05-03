package card

import "github.com/micro-plat/wechat/mp"

type Color struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// 获取卡券最新的颜色列表.
func GetColors(clt *mp.Context) (colors []Color, err error) {
	var result struct {
		mp.Error
		Colors []Color `json:"colors"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/getcolors?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	colors = result.Colors
	return
}
