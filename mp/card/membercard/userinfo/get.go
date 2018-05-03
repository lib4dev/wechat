package userinfo

import (
	"github.com/micro-plat/wechat/mp"
	"github.com/micro-plat/wechat/mp/card/code"
)

type CustomField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type UserInfo struct {
	openid          string        `json:"openid"`
	nickname        string        `json:"nickname"`
	sex             string        `json:"sex"`
	CustomFieldList []CustomField `json:"custom_field_list"`
}

// 拉取会员信息（积分查询）接口
func Get(clt *mp.Context, id *code.CardItemIdentifier) (info *UserInfo, err error) {
	var result struct {
		mp.Error
		UserInfo
	}

	incompleteURL := "https://api.weixin.qq.com/card/membercard/userinfo/get?access_token="
	if err = clt.PostJSON(incompleteURL, id, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.UserInfo
	return
}
