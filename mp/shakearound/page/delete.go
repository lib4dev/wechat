package page

import (
	"github.com/lib4dev/wechat/mp"
)

// 删除页面
func Delete(clt *mp.Context, pageIds []int64) (err error) {
	request := struct {
		PageIds []int64 `json:"page_ids,omitempty"`
	}{
		PageIds: pageIds,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/delete?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
