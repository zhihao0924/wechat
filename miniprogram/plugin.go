package miniprogram

import (
	"encoding/json"
	"fmt"

	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// Plugin 插件相关API
type Plugin struct {
	*MiniProgram
}

// 插件相关API地址
const (
	// 申请使用插件
	applyPluginURL = "https://api.weixin.qq.com/wxa/plugin"
	// 查询已添加的插件
	getPluginListURL = "https://api.weixin.qq.com/wxa/plugin"
	// 修改插件使用申请的状态
	setDevPluginApplyStatusURL = "https://api.weixin.qq.com/wxa/devplugin"
	// 查询插件开发者的权限
	getDevPluginPermissionURL = "https://api.weixin.qq.com/wxa/devplugin"
)

// PluginInfo 插件信息
type PluginInfo struct {
	AppID       string `json:"appid"`       // 插件AppID
	Status      int    `json:"status"`      // 插件状态
	Nickname    string `json:"nickname"`    // 插件昵称
	HeadImgURL  string `json:"headimgurl"`  // 插件头像
	Description string `json:"description"` // 插件描述
	Version     string `json:"version"`     // 插件版本号
}

// PluginDevApplyInfo 插件开发者收到的插件使用申请信息
type PluginDevApplyInfo struct {
	AppID    string `json:"appid"`    // 申请者的AppID
	Status   int    `json:"status"`   // 申请的状态
	Nickname string `json:"nickname"` // 申请者的昵称
	HeadImg  string `json:"headimg"`  // 申请者的头像
	Reason   string `json:"reason"`   // 申请者填写的申请理由
}

// ApplyPlugin 申请使用插件
func (p *Plugin) ApplyPlugin(pluginAppID string, reason string) error {
	accessToken, err := p.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", applyPluginURL, accessToken)

	data := struct {
		Action      string `json:"action"`
		PluginAppID string `json:"plugin_appid"`
		Reason      string `json:"reason"`
	}{
		Action:      "apply",
		PluginAppID: pluginAppID,
		Reason:      reason,
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return err
	}

	var result common.WechatError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return fmt.Errorf("申请使用插件失败: %s", result.ErrMsg)
	}

	return nil
}

// GetPluginList 查询已添加的插件
func (p *Plugin) GetPluginList() ([]PluginInfo, error) {
	accessToken, err := p.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", getPluginListURL, accessToken)

	data := struct {
		Action string `json:"action"`
	}{
		Action: "list",
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return nil, err
	}

	var result struct {
		common.WechatError
		PluginList []PluginInfo `json:"plugin_list"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("查询已添加的插件失败: %s", result.ErrMsg)
	}

	return result.PluginList, nil
}

// UnbindPlugin 解除插件绑定
func (p *Plugin) UnbindPlugin(pluginAppID string) error {
	accessToken, err := p.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", applyPluginURL, accessToken)

	data := struct {
		Action      string `json:"action"`
		PluginAppID string `json:"plugin_appid"`
	}{
		Action:      "unbind",
		PluginAppID: pluginAppID,
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return err
	}

	var result common.WechatError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return fmt.Errorf("解除插件绑定失败: %s", result.ErrMsg)
	}

	return nil
}

// GetDevApplyList 查询插件使用方的申请列表（插件开发者使用）
func (p *Plugin) GetDevApplyList(page, num int) ([]PluginDevApplyInfo, error) {
	accessToken, err := p.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", getDevPluginPermissionURL, accessToken)

	data := struct {
		Action string `json:"action"`
		Page   int    `json:"page"`
		Num    int    `json:"num"`
	}{
		Action: "dev_apply_list",
		Page:   page,
		Num:    num,
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return nil, err
	}

	var result struct {
		common.WechatError
		ApplyList []PluginDevApplyInfo `json:"apply_list"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("查询插件使用方的申请列表失败: %s", result.ErrMsg)
	}

	return result.ApplyList, nil
}

// SetDevPluginApplyStatus 修改插件使用申请的状态（插件开发者使用）
func (p *Plugin) SetDevPluginApplyStatus(appID string, status int) error {
	accessToken, err := p.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", setDevPluginApplyStatusURL, accessToken)

	data := struct {
		Action string `json:"action"`
		AppID  string `json:"appid"`
		Status int    `json:"status"`
	}{
		Action: "dev_agree",
		AppID:  appID,
		Status: status,
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return err
	}

	var result common.WechatError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return fmt.Errorf("修改插件使用申请的状态失败: %s", result.ErrMsg)
	}

	return nil
}

// GetPluginDevPermission 查询插件开发者的权限（插件开发者使用）
func (p *Plugin) GetPluginDevPermission() ([]string, error) {
	accessToken, err := p.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", getDevPluginPermissionURL, accessToken)

	data := struct {
		Action string `json:"action"`
	}{
		Action: "dev_permission",
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return nil, err
	}

	var result struct {
		common.WechatError
		Permissions []string `json:"permissions"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("查询插件开发者的权限失败: %s", result.ErrMsg)
	}

	return result.Permissions, nil
}

func newPlugin(mp *MiniProgram) *Plugin {
	return &Plugin{mp}
}
