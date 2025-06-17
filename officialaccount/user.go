package officialaccount

import (
	"encoding/json"
	"fmt"
)

// User 用户信息
type User struct {
	Subscribe     int    `json:"subscribe"`      // 用户是否订阅该公众号标识
	OpenID        string `json:"openid"`         // 用户的标识，对当前公众号唯一
	Nickname      string `json:"nickname"`       // 用户的昵称
	Sex           int    `json:"sex"`            // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language      string `json:"language"`       // 用户的语言，简体中文为zh_CN
	City          string `json:"city"`           // 用户所在城市
	Province      string `json:"province"`       // 用户所在省份
	Country       string `json:"country"`        // 用户所在国家
	HeadImgURL    string `json:"headimgurl"`     // 用户头像
	SubscribeTime int64  `json:"subscribe_time"` // 用户关注时间，为时间戳
	UnionID       string `json:"unionid"`        // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段
	Remark        string `json:"remark"`         // 公众号运营者对粉丝的备注
	GroupID       int    `json:"groupid"`        // 用户所在的分组ID
	TagIDList     []int  `json:"tagid_list"`     // 用户被打上的标签ID列表
}

// UserList 用户列表
type UserList struct {
	Total int      `json:"total"` // 关注该公众账号的总用户数
	Count int      `json:"count"` // 拉取的OPENID个数，最大值为10000
	Data  struct { // 列表数据，OPENID的列表
		OpenID []string `json:"openid"`
	} `json:"data"`
	NextOpenID string `json:"next_openid"` // 拉取列表的最后一个用户的OPENID
}

// Tag 用户标签
type Tag struct {
	ID    int    `json:"id"`    // 标签ID
	Name  string `json:"name"`  // 标签名称
	Count int    `json:"count"` // 此标签下粉丝数
}

// GetUser 获取用户基本信息
func (oa *OfficialAccount) GetUser(openID string, lang string) (*User, error) {
	if lang == "" {
		lang = "zh_CN"
	}

	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=%s",
		accessToken, openID, lang)

	response, err := oa.httpGet(url)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(response, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserList 获取用户列表
func (oa *OfficialAccount) GetUserList(nextOpenID string) (*UserList, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s",
		accessToken, nextOpenID)

	response, err := oa.httpGet(url)
	if err != nil {
		return nil, err
	}

	var userList UserList
	err = json.Unmarshal(response, &userList)
	if err != nil {
		return nil, err
	}

	return &userList, nil
}

// UpdateUserRemark 设置用户备注名
func (oa *OfficialAccount) UpdateUserRemark(openID, remark string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=%s", accessToken)
	data := struct {
		OpenID string `json:"openid"`
		Remark string `json:"remark"`
	}{
		OpenID: openID,
		Remark: remark,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// CreateTag 创建标签
func (oa *OfficialAccount) CreateTag(name string) (*Tag, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/create?access_token=%s", accessToken)
	data := struct {
		Tag struct {
			Name string `json:"name"`
		} `json:"tag"`
	}{
		Tag: struct {
			Name string `json:"name"`
		}{
			Name: name,
		},
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result struct {
		Tag Tag `json:"tag"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return &result.Tag, nil
}

// GetTags 获取公众号已创建的标签
func (oa *OfficialAccount) GetTags() ([]Tag, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s", accessToken)
	response, err := oa.httpGet(url)
	if err != nil {
		return nil, err
	}

	var result struct {
		Tags []Tag `json:"tags"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.Tags, nil
}

// UpdateTag 编辑标签
func (oa *OfficialAccount) UpdateTag(tagID int, name string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/update?access_token=%s", accessToken)
	data := struct {
		Tag struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"tag"`
	}{
		Tag: struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{
			ID:   tagID,
			Name: name,
		},
	}

	_, err = oa.httpPost(url, data)
	return err
}

// DeleteTag 删除标签
func (oa *OfficialAccount) DeleteTag(tagID int) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=%s", accessToken)
	data := struct {
		Tag struct {
			ID int `json:"id"`
		} `json:"tag"`
	}{
		Tag: struct {
			ID int `json:"id"`
		}{
			ID: tagID,
		},
	}

	_, err = oa.httpPost(url, data)
	return err
}

// GetTagUsers 获取标签下粉丝列表
func (oa *OfficialAccount) GetTagUsers(tagID int, nextOpenID string) (*UserList, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=%s", accessToken)
	data := struct {
		TagID      int    `json:"tagid"`
		NextOpenID string `json:"next_openid"`
	}{
		TagID:      tagID,
		NextOpenID: nextOpenID,
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var userList UserList
	err = json.Unmarshal(response, &userList)
	if err != nil {
		return nil, err
	}

	return &userList, nil
}

// BatchTagging 批量为用户打标签
func (oa *OfficialAccount) BatchTagging(tagID int, openIDList []string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=%s", accessToken)
	data := struct {
		OpenIDList []string `json:"openid_list"`
		TagID      int      `json:"tagid"`
	}{
		OpenIDList: openIDList,
		TagID:      tagID,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// BatchUntagging 批量为用户取消标签
func (oa *OfficialAccount) BatchUntagging(tagID int, openIDList []string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=%s", accessToken)
	data := struct {
		OpenIDList []string `json:"openid_list"`
		TagID      int      `json:"tagid"`
	}{
		OpenIDList: openIDList,
		TagID:      tagID,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// GetUserTags 获取用户身上的标签列表
func (oa *OfficialAccount) GetUserTags(openID string) ([]int, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=%s", accessToken)
	data := struct {
		OpenID string `json:"openid"`
	}{
		OpenID: openID,
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result struct {
		TagIDList []int `json:"tagid_list"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.TagIDList, nil
}
