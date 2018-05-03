package mmpaymkttransfers

import "github.com/micro-plat/wechat/mch"

// 发放代金券.
//  请求需要双向证书
func SendCoupon(clt *mch.Context, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(mch.APIBaseURL()+"/mmpaymkttransfers/send_coupon", req)
}
