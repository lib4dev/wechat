package promotion

import "github.com/micro-plat/wechat/mch"

// 企业付款.
//  NOTE: 请求需要双向证书
func Transfers(clt *mch.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(mch.APIBaseURL()+"/mmpaymkttransfers/promotion/transfers", req)
}
