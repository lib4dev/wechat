package mmpaymkttransfers

import "github.com/micro-plat/wechat/mch"

// 红包查询接口.
//  NOTE: 请求需要双向证书
func GetRedPackInfo(clt *mch.Context, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(mch.APIBaseURL()+"/mmpaymkttransfers/gethbinfo", req)
}
