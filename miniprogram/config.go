package miniprogram

import (
	"github.com/zhihao0924/wechat/auth"
)

// Config 小程序配置
type Config struct {
	// AppID 小程序的唯一标识
	AppID string
	// AppSecret 小程序的appsecret
	AppSecret string
	// AccessTokenHandle access token获取器
	AccessTokenHandle auth.AccessTokenHandle
}

// SetAccessTokenHandle 设置access token获取器
func (cfg *Config) SetAccessTokenHandle(accessTokenHandle auth.AccessTokenHandle) {
	cfg.AccessTokenHandle = accessTokenHandle
}

// GetAccessTokenHandle 获取access token处理器
func (cfg *Config) GetAccessTokenHandle() auth.AccessTokenHandle {
	if cfg.AccessTokenHandle == nil {
		cfg.AccessTokenHandle = auth.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret)
	}
	return cfg.AccessTokenHandle
}
