package menu

import (
	"testing"

	"github.com/lib4dev/wechat/mp"
)

func TestCreate(t *testing.T) {
	ctx := mp.NewContext(mp.NewDefaultAccessToken("wx9e02ddcc88e13fd4", "6acb2bf99177524beba3d97d54df2de5"))

	//ctx := mp.NewContext(mp.NewAccessToken("6acb2bf99177524beba3d97d54df2de5"))
	menu := &Menu{
		Buttons: []Button{
			Button{Type: ButtonTypeView, Name: "搜索", URL: "http://www.baidu.com"},
			Button{Type: ButtonTypeView, Name: "搜索", URL: "http://www.baidu.com"},
			Button{Type: ButtonTypeView, Name: "搜索", URL: "http://www.baidu.com"},
		},
	}
	err := Create(ctx, menu)
	if err != nil {
		t.Error(err)
	}

}
