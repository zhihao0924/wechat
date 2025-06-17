package officialaccount

import (
	"encoding/json"

	"github.com/zhihao0924/wechat/auth"
	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// OfficialAccount 公众号/服务号实例
type OfficialAccount struct {
	Config *Config

	// 各种API服务
	User    *User
	Menu    *Menu
	Message *Message
	QRCode  *QRCode
	// 可以根据需要添加更多服务
}

// NewOfficialAccount 实例化公众号/服务号API
func NewOfficialAccount(cfg *Config) *OfficialAccount {
	oa := &OfficialAccount{
		Config: cfg,
	}

	// 初始化各种服务
	oa.User = newUser(oa)
	oa.Menu = newMenu(oa)
	oa.Message = newMessage(oa)
	oa.QRCode = newQRCode(oa)

	return oa
}

// GetAccessToken 获取access_token
func (oa *OfficialAccount) GetAccessToken() (string, error) {
	return oa.Config.GetAccessTokenHandle().GetAccessToken()
}

// GetContext 获取上下文
func (oa *OfficialAccount) GetContext() *OfficialAccount {
	return oa
}

// User 用户管理相关API
type User struct {
	*OfficialAccount
}

func newUser(oa *OfficialAccount) *User {
	return &User{oa}
}

// GetUserInfo 获取用户基本信息
func (user *User) GetUserInfo(openID string, lang ...string) (*UserInfo, error) {
	language := "zh_CN"
	if len(lang) > 0 {
		language = lang[0]
	}

	accessToken, err := user.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := common.OfficialAccountUserInfoURL + "?access_token=" + accessToken + "&openid=" + openID + "&lang=" + language

	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	userInfo := &UserInfo{}
	err = json.Unmarshal(response, userInfo)
	if err != nil {
		return nil, err
	}

	if userInfo.ErrCode != 0 {
		return nil, common.NewWechatError(userInfo.ErrCode, userInfo.ErrMsg)
	}

	return userInfo, nil
}

// UserInfo 用户基本信息
type UserInfo struct {
	common.WechatError

	Subscribe      int    `json:"subscribe"`       // 用户是否订阅该公众号标识
	OpenID         string `json:"openid"`          // 用户的标识
	Nickname       string `json:"nickname"`        // 用户的昵称
	Sex            int    `json:"sex"`             // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City           string `json:"city"`            // 用户所在城市
	Country        string `json:"country"`         // 用户所在国家
	Province       string `json:"province"`        // 用户所在省份
	Language       string `json:"language"`        // 用户的语言，简体中文为zh_CN
	Headimgurl     string `json:"headimgurl"`      // 用户头像
	SubscribeTime  int64  `json:"subscribe_time"`  // 用户关注时间
	UnionID        string `json:"unionid"`         // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段
	Remark         string `json:"remark"`          // 公众号运营者对粉丝的备注
	GroupID        int    `json:"groupid"`         // 用户所在的分组ID
	TagIDList      []int  `json:"tagid_list"`      // 用户被打上的标签ID列表
	SubscribeScene string `json:"subscribe_scene"` // 返回用户关注的渠道来源
	QrScene        int    `json:"qr_scene"`        // 二维码扫码场景
	QrSceneStr     string `json:"qr_scene_str"`    // 二维码扫码场景描述
}

// Menu 自定义菜单
type Menu struct {
	*OfficialAccount
}

func newMenu(oa *OfficialAccount) *Menu {
	return &Menu{oa}
}

// Message 消息管理
type Message struct {
	*OfficialAccount
}

func newMessage(oa *OfficialAccount) *Message {
	return &Message{oa}
}

// QRCode 二维码
type QRCode struct {
	*OfficialAccount
}

func newQRCode(oa *OfficialAccount) *QRCode {
	return &QRCode{oa}
}
