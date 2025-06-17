// Package wechat 提供了微信服务号、公众号和小程序API的Go语言封装
package wechat

import (
	"github.com/zhihao0924/wechat/miniprogram"
	"github.com/zhihao0924/wechat/officialaccount"
)

// Wechat 结构体是SDK的主入口
type Wechat struct {
	// 可以在这里添加一些全局配置
}

// NewWechat 创建一个新的Wechat实例
func NewWechat() *Wechat {
	return &Wechat{}
}

// GetOfficialAccount 获取公众号/服务号实例
func (wc *Wechat) GetOfficialAccount(cfg *officialaccount.Config) *officialaccount.OfficialAccount {
	return officialaccount.NewOfficialAccount(cfg)
}

// GetMiniProgram 获取小程序实例
func (wc *Wechat) GetMiniProgram(cfg *miniprogram.Config) *miniprogram.MiniProgram {
	return miniprogram.NewMiniProgram(cfg)
}
