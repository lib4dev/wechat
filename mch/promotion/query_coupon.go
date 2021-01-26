package promotion

import "github.com/lib4dev/wechat/mch"

// 查询代金券信息.
func QueryCoupon(clt *mch.Context, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(mch.APIBaseURL()+"/promotion/query_coupon", req)
}
