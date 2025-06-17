package miniprogram

import (
	"encoding/json"
	"fmt"

	"github.com/zhihao0924/wechat/auth"
	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// MiniProgram 微信小程序实例
type MiniProgram struct {
	Config *Config

	// 各种API服务
	Auth      *Auth
	QRCode    *QRCode
	Template  *Template
	Analysis  *Analysis
	Crypto    *Crypto
	Subscribe *Subscribe
	Customer  *Customer
	Security  *Security
	Live      *Live
	Plugin    *Plugin
	Operation *Operation
	// 可以根据需要添加更多服务
}

// NewMiniProgram 实例化小程序API
func NewMiniProgram(cfg *Config) *MiniProgram {
	mp := &MiniProgram{
		Config: cfg,
	}

	// 初始化各种服务
	mp.Auth = newAuth(mp)
	mp.QRCode = newQRCode(mp)
	mp.Template = newTemplate(mp)
	mp.Analysis = newAnalysis(mp)
	mp.Crypto = newCrypto(mp)
	mp.Subscribe = newSubscribe(mp)
	mp.Customer = newCustomer(mp)
	mp.Security = newSecurity(mp)
	mp.Live = newLive(mp)
	mp.Plugin = newPlugin(mp)
	mp.Operation = newOperation(mp)

	return mp
}

// GetAccessToken 获取access_token
func (mp *MiniProgram) GetAccessToken() (string, error) {
	return mp.Config.GetAccessTokenHandle().GetAccessToken()
}

// GetContext 获取上下文
func (mp *MiniProgram) GetContext() *MiniProgram {
	return mp
}

// Auth 登录认证相关API
type Auth struct {
	*MiniProgram
}

// Code2SessionResult 登录凭证校验返回结果
type Code2SessionResult struct {
	common.WechatError
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

// Code2Session 登录凭证校验
func (auth *Auth) Code2Session(jsCode string) (*Code2SessionResult, error) {
	uri := fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		common.MiniProgramCode2SessionURL,
		auth.Config.AppID,
		auth.Config.AppSecret,
		jsCode)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	result := &Code2SessionResult{}
	err = json.Unmarshal(response, result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, common.NewWechatError(result.ErrCode, result.ErrMsg)
	}

	return result, nil
}

// GetPaidUnionID 获取支付后的用户UnionID
func (auth *Auth) GetPaidUnionID(openID, transactionID string) (string, error) {
	accessToken, err := auth.GetAccessToken()
	if err != nil {
		return "", err
	}

	uri := fmt.Sprintf("%s?access_token=%s&openid=%s&transaction_id=%s",
		common.MiniProgramGetPaidUnionIDURL,
		accessToken,
		openID,
		transactionID)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return "", err
	}

	var result struct {
		common.WechatError
		UnionID string `json:"unionid"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return "", err
	}

	if result.ErrCode != 0 {
		return "", common.NewWechatError(result.ErrCode, result.ErrMsg)
	}

	return result.UnionID, nil
}

// QRCode 小程序码相关API
type QRCode struct {
	*MiniProgram
}

// Template 模板消息相关API
type Template struct {
	*MiniProgram
}

// Analysis 数据分析相关API
type Analysis struct {
	*MiniProgram
}

// 初始化各种服务
func newAuth(mp *MiniProgram) *Auth {
	return &Auth{mp}
}

func newQRCode(mp *MiniProgram) *QRCode {
	return &QRCode{mp}
}

func newTemplate(mp *MiniProgram) *Template {
	return &Template{mp}
}

func newAnalysis(mp *MiniProgram) *Analysis {
	return &Analysis{mp}
}
