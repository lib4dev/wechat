package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/lib4dev/wechat/mch"
	"github.com/lib4dev/wechat/mp"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/app"
)

//AppWXConf 应用程序配置
type AppWXConf struct {
	WX []WXConfig `json:"wx" valid:"required"`
}

//WXConfig 微信公众号配置
type WXConfig struct {
	AppID           string `json:"appid" valid:"ascii,required"`
	AppSecret       string `json:"secret" valid:"ascii,required"`
	Token           string `json:"token" valid:"ascii"`
	EncodingAESKey  string `json:"aes-key" valid:"ascii"`
	ServeURL        string `json:"msg-notify-url" valid:"ascii,required"`
	EnablePayNotify bool   `json:"enable-pay-notify"` //启动支付结果充值
}
type AppPayConf struct {
	WX []PayConfig `json:"wx" valid:"required"`
}
type PayConfig struct {
	PayNotifyURL string `json:"pay-notify-url" valid:"ascii,required"` //支付 - 接受微信支付结果通知的接口地址
	PayMchID     string `json:"mchid" valid:"ascii,required"`          //支付 - 商户 ID
	PayKey       string `json:"pay-key" valid:"ascii,required"`        //支付 - 商户后台设置的支付 key
	SubAppId     string `json:"sub-appid" valid:"ascii"`               //支付 - 商户后台设置的支付 key
	SubMchId     string `json:"sub-mchid" valid:"ascii"`               //支付 - 商户后台设置的支付 key
}

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func bindConf() {
	hydra.Conf.API("9999").Sub("app", `{
		"wx":[{
			"appid": "wx9e02ddcc88e13fd4",
			"secret": "6acb2bf99177524beba3d97d54df2de5",
			"token":"oTSvVuXdjb9Xx1FPi6bz",
			"aes-key": "D3mgxDexQDuqHm1MIWsyvhLMd3Y303cmf05JgjD9ZWY",
			"msg-notify-url": "/"
		}]		
	}`)
}

//bind 检查并缓存配置文件，初始化微信服务器用于接收微信通知
func bind(App *hydra.MicroApp) {
	bindConf()

	App.OnStarting(func(appConf app.IAPPConf) error {

		var wxConf AppWXConf
		if _, err := appConf.GetServerConf().GetSubObject("app", &wxConf); err != nil {
			return fmt.Errorf("获取vueconf配置失败:%v", err)
		}

		if b, err := govalidator.ValidateStruct(&wxConf); !b || len(wxConf.WX) == 0 {
			err = fmt.Errorf("app 配置文件有误:%v", err)
			return err
		}

		var payConf AppPayConf
		if _, err := appConf.GetServerConf().GetSubObject("app", &payConf); err != nil {
			return err
		}

		for i, wx := range wxConf.WX {
			//创建微信处理服务
			ctx := &mp.WConf{
				AppID:          wx.AppID,
				AppSecret:      wx.AppSecret,
				Token:          wx.Token,
				EncodingAESKey: wx.EncodingAESKey,
			}
			App.Micro(wx.ServeURL, mp.NewMessageSeverHandler(ctx, recvMessageHandler))
			if !wx.EnablePayNotify {
				continue
			}

			//验证支付配置是否正确
			if b, err := govalidator.ValidateStruct(&payConf.WX[i]); !b || len(payConf.WX) == 0 {
				err = fmt.Errorf("app 配置文件有误:%v", err)
				return err
			}
			pc := payConf.WX[i]
			App.Micro(pc.PayNotifyURL, notifyServeHandler(&mch.PayConf{
				AppId:    wx.AppID,
				ApiKey:   pc.PayKey,
				MchId:    pc.PayMchID,
				SubAppId: pc.SubAppId,
				SubMchId: pc.SubMchId,
			}))
		}

		return nil
	})
}
