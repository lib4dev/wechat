package statistics

import (
	"github.com/micro-plat/wechat/mp"
	"github.com/micro-plat/wechat/mp/shakearound/device"
)

// 以设备为维度的数据统计接口
func Device(clt *mp.Context, deviceIdentifier *device.DeviceIdentifier, beginDate, endDate int64) (data []StatisticsBase, err error) {
	request := struct {
		DeviceIdentifier *device.DeviceIdentifier `json:"device_identifier,omitempty"`
		BeginDate        int64                    `json:"begin_date"`
		EndDate          int64                    `json:"end_date"`
	}{
		DeviceIdentifier: deviceIdentifier,
		BeginDate:        beginDate,
		EndDate:          endDate,
	}

	var result struct {
		mp.Error
		Data []StatisticsBase `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/statistics/device?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	data = result.Data
	return
}
