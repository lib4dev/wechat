package mmpaymkttransfers

import "github.com/micro-plat/wechat/mch"

// 查询企业付款.
//  NOTE: 请求需要双向证书
func GetTransferInfo(clt *mch.Context, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(mch.APIBaseURL()+"/mmpaymkttransfers/gettransferinfo", req)
}
