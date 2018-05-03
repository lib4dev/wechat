package device

import "github.com/micro-plat/wechat/mp"

// 删除设备
func Delete(clt *mp.Context, bssid string) (err error) {
	request := struct {
		BSSID string `json:"bssid"`
	}{
		BSSID: bssid,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/bizwifi/device/delete?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
