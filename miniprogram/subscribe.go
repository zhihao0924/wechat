package miniprogram

import (
	"encoding/json"
	"fmt"

	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// Subscribe 订阅消息相关API
type Subscribe struct {
	*MiniProgram
}

// 订阅消息相关API地址
const (
	// 发送订阅消息
	sendSubscribeMessageURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"
	// 获取当前帐号所设置的类目信息
	getCategoryURL = "https://api.weixin.qq.com/wxaapi/newtmpl/getcategory"
	// 获取模板标题列表
	getPubTemplateTitlesURL = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles"
	// 获取模板标题下的关键词库
	getPubTemplateKeywordsURL = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords"
	// 添加模板
	addTemplateURL = "https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate"
	// 获取个人模板列表
	getTemplateListURL = "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate"
	// 删除模板
	deleteTemplateURL = "https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate"
)

// SubscribeMessage 订阅消息
type SubscribeMessage struct {
	ToUser           string                       `json:"touser"`                      // 接收者（用户）的 openid
	TemplateID       string                       `json:"template_id"`                 // 所需下发的订阅模板id
	Page             string                       `json:"page,omitempty"`              // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Data             map[string]SubscribeDataItem `json:"data"`                        // 模板内容
	MiniprogramState string                       `json:"miniprogram_state,omitempty"` // 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string                       `json:"lang,omitempty"`              // 进入小程序查看"的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

// SubscribeDataItem 订阅消息数据项
type SubscribeDataItem struct {
	Value string `json:"value"`
}

// Category 类目信息
type Category struct {
	ID   int    `json:"id"`   // 类目ID
	Name string `json:"name"` // 类目名称
}

// TemplateTitle 模板标题
type TemplateTitle struct {
	TID        int    `json:"tid"`        // 模板标题 id
	Title      string `json:"title"`      // 模板标题
	Type       int    `json:"type"`       // 模板类型，2 为一次性订阅，3 为长期订阅
	CategoryID string `json:"categoryId"` // 模板所属类目 id
}

// TemplateKeyword 模板关键词
type TemplateKeyword struct {
	KID     int    `json:"kid"`     // 关键词 id
	Name    string `json:"name"`    // 关键词内容
	Example string `json:"example"` // 关键词内容对应的示例
	Rule    string `json:"rule"`    // 参数类型
}

// Template 个人模板
type Template struct {
	PriTmplID string `json:"priTmplId"` // 模板ID
	Title     string `json:"title"`     // 模板标题
	Content   string `json:"content"`   // 模板内容
	Example   string `json:"example"`   // 模板示例
	Type      int    `json:"type"`      // 模板类型，2 为一次性订阅，3 为长期订阅
}

// Send 发送订阅消息
func (s *Subscribe) Send(message *SubscribeMessage) error {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", sendSubscribeMessageURL, accessToken)

	response, err := util.HTTPPost(uri, message)
	if err != nil {
		return err
	}

	var result common.WechatError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return fmt.Errorf("发送订阅消息失败: %s", result.ErrMsg)
	}

	return nil
}

// SendSubscribe 快速发送订阅消息
func (s *Subscribe) SendSubscribe(toUser, templateID string, data map[string]SubscribeDataItem, page ...string) error {
	message := &SubscribeMessage{
		ToUser:     toUser,
		TemplateID: templateID,
		Data:       data,
	}

	if len(page) > 0 {
		message.Page = page[0]
	}

	return s.Send(message)
}

// CreateSubscribeData 创建订阅消息数据
func (s *Subscribe) CreateSubscribeData() map[string]SubscribeDataItem {
	return make(map[string]SubscribeDataItem)
}

// AddSubscribeData 添加订阅消息数据项
func (s *Subscribe) AddSubscribeData(data map[string]SubscribeDataItem, key, value string) {
	data[key] = SubscribeDataItem{
		Value: value,
	}
}

// GetCategory 获取当前帐号所设置的类目信息
func (s *Subscribe) GetCategory() ([]Category, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", getCategoryURL, accessToken)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	var result struct {
		common.WechatError
		Data []Category `json:"data"`
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("获取类目信息失败: %s", result.ErrMsg)
	}

	return result.Data, nil
}

// GetPubTemplateTitles 获取模板标题列表
func (s *Subscribe) GetPubTemplateTitles(categoryID string, start, limit int) ([]TemplateTitle, int, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return nil, 0, err
	}

	uri := fmt.Sprintf("%s?access_token=%s&ids=%s&start=%d&limit=%d",
		getPubTemplateTitlesURL, accessToken, categoryID, start, limit)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, 0, err
	}

	var result struct {
		common.WechatError
		Data       []TemplateTitle `json:"data"`
		TotalCount int             `json:"count"`
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, 0, err
	}

	if result.ErrCode != 0 {
		return nil, 0, fmt.Errorf("获取模板标题列表失败: %s", result.ErrMsg)
	}

	return result.Data, result.TotalCount, nil
}

// GetPubTemplateKeywords 获取模板标题下的关键词库
func (s *Subscribe) GetPubTemplateKeywords(tid int) ([]TemplateKeyword, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s&tid=%d", getPubTemplateKeywordsURL, accessToken, tid)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	var result struct {
		common.WechatError
		Data []TemplateKeyword `json:"data"`
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("获取模板关键词列表失败: %s", result.ErrMsg)
	}

	return result.Data, nil
}

// AddTemplate 添加模板
func (s *Subscribe) AddTemplate(tid int, kidList []int, sceneDesc string) (string, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return "", err
	}

	uri := fmt.Sprintf("%s?access_token=%s", addTemplateURL, accessToken)

	data := struct {
		TID       int    `json:"tid"`
		KidList   []int  `json:"kidList"`
		SceneDesc string `json:"sceneDesc"`
	}{
		TID:       tid,
		KidList:   kidList,
		SceneDesc: sceneDesc,
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return "", err
	}

	var result struct {
		common.WechatError
		PriTmplID string `json:"priTmplId"`
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return "", err
	}

	if result.ErrCode != 0 {
		return "", fmt.Errorf("添加模板失败: %s", result.ErrMsg)
	}

	return result.PriTmplID, nil
}

// GetTemplateList 获取个人模板列表
func (s *Subscribe) GetTemplateList() ([]Template, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", getTemplateListURL, accessToken)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	var result struct {
		common.WechatError
		Data []Template `json:"data"`
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("获取个人模板列表失败: %s", result.ErrMsg)
	}

	return result.Data, nil
}

// DeleteTemplate 删除模板
func (s *Subscribe) DeleteTemplate(priTmplID string) error {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", deleteTemplateURL, accessToken)

	data := struct {
		PriTmplID string `json:"priTmplId"`
	}{
		PriTmplID: priTmplID,
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
		return fmt.Errorf("删除模板失败: %s", result.ErrMsg)
	}

	return nil
}

func newSubscribe(mp *MiniProgram) *Subscribe {
	return &Subscribe{mp}
}
