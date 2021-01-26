package menu

import (
	"testing"

	"github.com/lib4dev/wechat/mp/menu"

	"github.com/lib4dev/wechat/mp"
)

func TestCreate(t *testing.T) {
	tk := mp.NewDefaultAccessTokenByURL("wx9e02ddcc88e13fd4", "6acb2bf99177524beba3d97d54df2de5", "http://192.168.5.71:9999/wx9e02ddcc88e13fd4/wechat/token/get")
	ctx := mp.NewContext(tk)
	mu := &menu.Menu{
		Buttons: []menu.Button{
			menu.Button{Type: menu.ButtonTypeView, Name: "搜索1", URL: "http://www.baidu.com"},
			menu.Button{Type: menu.ButtonTypeView, Name: "搜索2", URL: "http://www.baidu.com"},
			menu.Button{Type: menu.ButtonTypeView, Name: "搜索3", URL: "http://www.baidu.com"},
		},
	}
	err := menu.Create(ctx, mu)
	if err != nil {
		t.Error(err)
	}

}
