package menu

import (
	"github.com/micro-plat/wechat/mp/core"
)

// 创建自定义菜单.
func Create(clt *core.Context, menu *Menu) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token="

	var result core.Error
	if err = clt.PostJSON(incompleteURL, menu, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 查询自定义菜单.
func Get(clt *core.Context) (menu *Menu, conditionalMenus []Menu, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/menu/get?access_token="

	var result struct {
		core.Error
		Menu             Menu   `json:"menu"`
		ConditionalMenus []Menu `json:"conditionalmenu"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	menu = &result.Menu
	conditionalMenus = result.ConditionalMenus
	return
}

// 删除自定义菜单.
func Delete(clt *core.Context) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token="

	var result core.Error
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
