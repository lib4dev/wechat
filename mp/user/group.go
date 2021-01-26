package user

import (
	"github.com/lib4dev/wechat/mp"
)

// GroupId 查询用户所在分组.
func GroupId(clt *mp.Context, openId string) (groupId int64, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/groups/getid?access_token="

	var request = struct {
		OpenId string `json:"openid"`
	}{
		OpenId: openId,
	}
	var result struct {
		mp.Error
		GroupId int64 `json:"groupid"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	groupId = result.GroupId
	return
}

// MoveToGroup 移动用户分组.
func MoveToGroup(clt *mp.Context, openId string, toGroupId int64) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token="

	var request = struct {
		OpenId    string `json:"openid"`
		ToGroupId int64  `json:"to_groupid"`
	}{
		OpenId:    openId,
		ToGroupId: toGroupId,
	}
	var result mp.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// BatchMoveToGroup 批量移动用户分组.
func BatchMoveToGroup(clt *mp.Context, openIdList []string, toGroupId int64) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/groups/members/batchupdate?access_token="

	if len(openIdList) <= 0 {
		return
	}

	var request = struct {
		OpenIdList []string `json:"openid_list,omitempty"`
		ToGroupId  int64    `json:"to_groupid"`
	}{
		OpenIdList: openIdList,
		ToGroupId:  toGroupId,
	}
	var result mp.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
