package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zhihao0924/wechat/common"
)

// HTTPClient 是一个HTTP客户端接口
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	// DefaultHTTPClient 默认的HTTP客户端
	DefaultHTTPClient HTTPClient = &http.Client{
		Timeout: 30 * time.Second,
	}
)

// HTTPGet 发送GET请求
func HTTPGet(url string) ([]byte, error) {
	return HTTPGetWithClient(url, DefaultHTTPClient)
}

// HTTPGetWithClient 使用指定的HTTP客户端发送GET请求
func HTTPGetWithClient(url string, client HTTPClient) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP请求失败，状态码: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 检查微信API返回的错误
	var wechatErr common.WechatError
	if err := json.Unmarshal(body, &wechatErr); err == nil {
		if wechatErr.ErrCode != common.ErrCodeOK {
			return nil, &wechatErr
		}
	}

	return body, nil
}

// HTTPPost 发送POST请求
func HTTPPost(url string, data interface{}) ([]byte, error) {
	return HTTPPostWithClient(url, data, DefaultHTTPClient)
}

// HTTPPostWithClient 使用指定的HTTP客户端发送POST请求
func HTTPPostWithClient(url string, data interface{}, client HTTPClient) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP请求失败，状态码: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 检查微信API返回的错误
	var wechatErr common.WechatError
	if err := json.Unmarshal(body, &wechatErr); err == nil {
		if wechatErr.ErrCode != common.ErrCodeOK {
			return nil, &wechatErr
		}
	}

	return body, nil
}
