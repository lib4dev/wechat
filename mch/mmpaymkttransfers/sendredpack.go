package mmpaymkttransfers

import "github.com/micro-plat/wechat/mch"

// 红包发放.
//  NOTE: 请求需要双向证书
func SendRedPack(clt *mch.Context, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(mch.APIBaseURL()+"/mmpaymkttransfers/sendredpack", req)
}
