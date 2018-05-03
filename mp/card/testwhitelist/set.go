package testwhitelist

import "github.com/micro-plat/wechat/mp"

type SetParameters struct {
	OpenIdList   []string `json:"openid,omitempty"`   // 测试的openid列表
	UserNameList []string `json:"username,omitempty"` // 测试的微信号列表
}

// 设置测试白名单
func Set(clt *mp.Context, para *SetParameters) (err error) {
	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/card/testwhitelist/set?access_token="
	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
