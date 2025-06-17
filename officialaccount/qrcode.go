package officialaccount

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// QRCode 二维码
type QRCode struct {
	*OfficialAccount
}

// QRCodeRequest 二维码请求参数
type QRCodeRequest struct {
	ExpireSeconds int64  `json:"expire_seconds,omitempty"`
	ActionName    string `json:"action_name"`
	ActionInfo    struct {
		Scene struct {
			SceneID  int    `json:"scene_id,omitempty"`
			SceneStr string `json:"scene_str,omitempty"`
		} `json:"scene"`
	} `json:"action_info"`
}

// QRCodeResult 二维码响应结果
type QRCodeResult struct {
	common.WechatError
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	URL           string `json:"url"`
}

// NewQRCode 实例化二维码生成接口
func newQRCode(oa *OfficialAccount) *QRCode {
	return &QRCode{oa}
}

// CreateTemporary 创建临时二维码
func (qr *QRCode) CreateTemporary(sceneID int, expireSeconds int64) (*QRCodeResult, error) {
	accessToken, err := qr.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", common.OfficialAccountQRCodeCreateURL, accessToken)

	req := &QRCodeRequest{
		ExpireSeconds: expireSeconds,
		ActionName:    "QR_SCENE",
	}
	req.ActionInfo.Scene.SceneID = sceneID

	response, err := util.HTTPPost(uri, req)
	if err != nil {
		return nil, err
	}

	var result QRCodeResult
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("创建临时二维码失败: %s", result.ErrMsg)
	}

	return &result, nil
}

// CreateTemporaryStr 创建临时字符串二维码
func (qr *QRCode) CreateTemporaryStr(sceneStr string, expireSeconds int64) (*QRCodeResult, error) {
	accessToken, err := qr.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", common.OfficialAccountQRCodeCreateURL, accessToken)

	req := &QRCodeRequest{
		ExpireSeconds: expireSeconds,
		ActionName:    "QR_STR_SCENE",
	}
	req.ActionInfo.Scene.SceneStr = sceneStr

	response, err := util.HTTPPost(uri, req)
	if err != nil {
		return nil, err
	}

	var result QRCodeResult
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("创建临时字符串二维码失败: %s", result.ErrMsg)
	}

	return &result, nil
}

// CreatePermanent 创建永久二维码
func (qr *QRCode) CreatePermanent(sceneID int) (*QRCodeResult, error) {
	if sceneID > 100000 {
		return nil, fmt.Errorf("永久二维码场景ID不能大于100000")
	}

	accessToken, err := qr.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", common.OfficialAccountQRCodeCreateURL, accessToken)

	req := &QRCodeRequest{
		ActionName: "QR_LIMIT_SCENE",
	}
	req.ActionInfo.Scene.SceneID = sceneID

	response, err := util.HTTPPost(uri, req)
	if err != nil {
		return nil, err
	}

	var result QRCodeResult
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("创建永久二维码失败: %s", result.ErrMsg)
	}

	return &result, nil
}

// CreatePermanentStr 创建永久字符串二维码
func (qr *QRCode) CreatePermanentStr(sceneStr string) (*QRCodeResult, error) {
	accessToken, err := qr.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", common.OfficialAccountQRCodeCreateURL, accessToken)

	req := &QRCodeRequest{
		ActionName: "QR_LIMIT_STR_SCENE",
	}
	req.ActionInfo.Scene.SceneStr = sceneStr

	response, err := util.HTTPPost(uri, req)
	if err != nil {
		return nil, err
	}

	var result QRCodeResult
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("创建永久字符串二维码失败: %s", result.ErrMsg)
	}

	return &result, nil
}

// GetQRCodeURL 获取二维码图片URL
func (qr *QRCode) GetQRCodeURL(ticket string) string {
	return common.OfficialAccountQRCodeShowURL + "?ticket=" + ticket
}

// GetQRCodeTicket 通过场景ID获取二维码ticket
func (qr *QRCode) GetQRCodeTicket(sceneID int, expireSeconds ...int64) (string, error) {
	var expire int64 = 2592000 // 默认30天
	if len(expireSeconds) > 0 && expireSeconds[0] > 0 {
		expire = expireSeconds[0]
	}

	result, err := qr.CreateTemporary(sceneID, expire)
	if err != nil {
		return "", err
	}

	return result.Ticket, nil
}

// GetQRCodeTicketByStr 通过场景字符串获取二维码ticket
func (qr *QRCode) GetQRCodeTicketByStr(sceneStr string, expireSeconds ...int64) (string, error) {
	var expire int64 = 2592000 // 默认30天
	if len(expireSeconds) > 0 && expireSeconds[0] > 0 {
		expire = expireSeconds[0]
	}

	result, err := qr.CreateTemporaryStr(sceneStr, expire)
	if err != nil {
		return "", err
	}

	return result.Ticket, nil
}
