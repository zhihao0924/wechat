package miniprogram

import (
	"encoding/json"
	"fmt"

	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// Template 模板消息相关API
type Template struct {
	*MiniProgram
}

// TemplateMessage 模板消息数据
type TemplateMessage struct {
	ToUser          string                 `json:"touser"`           // 接收者（用户）的 openid
	TemplateID      string                 `json:"template_id"`      // 所需下发的模板消息的id
	Page            string                 `json:"page,omitempty"`   // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	FormID          string                 `json:"form_id"`          // 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
	Data            map[string]interface{} `json:"data"`             // 模板内容，不填则下发空模板。具体格式请参考示例。
	EmphasisKeyword string                 `json:"emphasis_keyword"` // 模板需要放大的关键词，不填则默认无放大
}

// TemplateItem 模板数据项
type TemplateItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// Send 发送模板消息
func (tpl *Template) Send(message *TemplateMessage) error {
	accessToken, err := tpl.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", common.MiniProgramSendTemplateURL, accessToken)

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
		return fmt.Errorf("发送模板消息失败: %s", result.ErrMsg)
	}

	return nil
}

// SendTemplate 快速发送模板消息
func (tpl *Template) SendTemplate(toUser, templateID, formID string, data map[string]interface{}, page ...string) error {
	message := &TemplateMessage{
		ToUser:     toUser,
		TemplateID: templateID,
		FormID:     formID,
		Data:       data,
	}

	if len(page) > 0 {
		message.Page = page[0]
	}

	return tpl.Send(message)
}

// CreateTemplateData 创建模板数据
func (tpl *Template) CreateTemplateData() map[string]interface{} {
	return make(map[string]interface{})
}

// AddTemplateData 添加模板数据项
func (tpl *Template) AddTemplateData(data map[string]interface{}, key, value string, color ...string) {
	item := TemplateItem{
		Value: value,
	}
	if len(color) > 0 {
		item.Color = color[0]
	}
	data[key] = item
}

// Example:
// template := miniProgram.Template
// data := template.CreateTemplateData()
// template.AddTemplateData(data, "keyword1", "订单已支付", "#173177")
// template.AddTemplateData(data, "keyword2", "￥88.00", "#173177")
// err := template.SendTemplate("openid", "template_id", "form_id", data, "pages/index/index")

func newTemplate(mp *MiniProgram) *Template {
	return &Template{mp}
}
