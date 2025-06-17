package common

import (
	"fmt"
)

// WechatError 定义微信API返回的错误
type WechatError struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// Error 实现error接口
func (e *WechatError) Error() string {
	return fmt.Sprintf("微信错误: 错误码=%d, 错误信息=%s", e.ErrCode, e.ErrMsg)
}

// NewWechatError 创建一个新的WechatError
func NewWechatError(code int, msg string) *WechatError {
	return &WechatError{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

// IsWechatError 判断错误是否为WechatError类型
func IsWechatError(err error) (*WechatError, bool) {
	if wechatErr, ok := err.(*WechatError); ok {
		return wechatErr, true
	}
	return nil, false
}
