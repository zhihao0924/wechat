package auth

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// AccessToken 微信接口调用凭证
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	CreateTime  int64  `json:"-"` // token创建时间
}

// AccessTokenHandle AccessToken接口

type AccessTokenHandle interface {
	GetAccessToken() (string, error)
}

// DefaultAccessToken 默认AccessToken 获取
type DefaultAccessToken struct {
	appID           string
	appSecret       string
	accessToken     *AccessToken
	accessTokenLock *sync.RWMutex
}

// NewDefaultAccessToken 实例化一个默认的AccessToken获取器
func NewDefaultAccessToken(appID, appSecret string) AccessTokenHandle {
	return &DefaultAccessToken{
		appID:           appID,
		appSecret:       appSecret,
		accessTokenLock: new(sync.RWMutex),
	}
}

// GetAccessToken 获取AccessToken
func (ak *DefaultAccessToken) GetAccessToken() (string, error) {
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()

	if ak.accessToken == nil || ak.isExpired() {
		if err := ak.refreshAccessToken(); err != nil {
			return "", err
		}
	}
	return ak.accessToken.AccessToken, nil
}

// isExpired 判断token是否过期
func (ak *DefaultAccessToken) isExpired() bool {
	if ak.accessToken == nil {
		return true
	}
	return time.Now().Unix() >= (ak.accessToken.CreateTime + ak.accessToken.ExpiresIn - 200)
}

// refreshAccessToken 刷新AccessToken
func (ak *DefaultAccessToken) refreshAccessToken() error {
	url := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s",
		common.WechatAPIURL+"/cgi-bin/token",
		ak.appID,
		ak.appSecret,
	)

	body, err := util.HTTPGet(url)
	if err != nil {
		return err
	}

	var result struct {
		common.WechatError
		AccessToken
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return fmt.Errorf("get access_token error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
	}

	result.AccessToken.CreateTime = time.Now().Unix()
	ak.accessToken = &result.AccessToken
	return nil
}

// GetAccessTokenFromServer 从微信服务器获取新的access_token
func GetAccessTokenFromServer(appID, appSecret string) (*AccessToken, error) {
	url := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s",
		common.WechatAPIURL+"/cgi-bin/token",
		appID,
		appSecret,
	)

	body, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}

	var result struct {
		common.WechatError
		AccessToken
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("get access_token error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
	}

	result.AccessToken.CreateTime = time.Now().Unix()
	return &result.AccessToken, nil
}
