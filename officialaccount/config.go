package officialaccount

import (
	"github.com/zhihao0924/wechat/auth"
)

// Config 公众号配置
type Config struct {
	// AppID 公众号的唯一标识
	AppID string
	// AppSecret 公众号的appsecret
	AppSecret string
	// Token 用于生成签名校验请求合法性
	Token string
	// EncodingAESKey 消息加解密密钥
	EncodingAESKey string
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
