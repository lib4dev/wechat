package mmpaymkttransfers

import "github.com/micro-plat/wechat/mch"

// 查询代金券批次信息.
func QueryCouponStock(clt *mch.Context, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(mch.APIBaseURL()+"/mmpaymkttransfers/query_coupon_stock", req)
}
