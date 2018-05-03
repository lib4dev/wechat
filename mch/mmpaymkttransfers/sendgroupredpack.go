package mmpaymkttransfers

import "github.com/micro-plat/wechat/mch"

// 发放裂变红包.
//  NOTE: 请求需要双向证书
func SendGroupRedPack(clt *mch.Context, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(mch.APIBaseURL()+"/mmpaymkttransfers/sendgroupredpack", req)
}
