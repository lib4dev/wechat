package poi

import (
	"github.com/lib4dev/wechat/mp"
)

// CategoryList 获取门店类目表.
func CategoryList(clt *mp.Context) (list []string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/api_getwxcategory?access_token="

	var result struct {
		mp.Error
		CategoryList []string `json:"category_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.CategoryList
	return
}
