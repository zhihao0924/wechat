package miniprogram

import (
	"encoding/json"
	"fmt"

	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// Customer 客服消息相关API
type Customer struct {
	*MiniProgram
}

// 客服消息相关API地址
const (
	// 发送客服消息
	sendCustomerMessageURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
	// 获取客服消息内的临时素材
	getTempMediaURL = "https://api.weixin.qq.com/cgi-bin/media/get"
	// 发送客服输入状态
	setTypingURL = "https://api.weixin.qq.com/cgi-bin/message/custom/typing"
)

// MessageType 消息类型
type MessageType string

const (
	// MessageTypeText 文本消息
	MessageTypeText MessageType = "text"
	// MessageTypeImage 图片消息
	MessageTypeImage MessageType = "image"
	// MessageTypeLink 图文链接消息
	MessageTypeLink MessageType = "link"
	// MessageTypeMiniProgramPage 小程序卡片消息
	MessageTypeMiniProgramPage MessageType = "miniprogrampage"
)

// CustomerMessage 客服消息
type CustomerMessage struct {
	ToUser          string                  `json:"touser"`                    // 用户的OpenID
	MsgType         MessageType             `json:"msgtype"`                   // 消息类型
	Text            *TextMessage            `json:"text,omitempty"`            // 文本消息
	Image           *ImageMessage           `json:"image,omitempty"`           // 图片消息
	Link            *LinkMessage            `json:"link,omitempty"`            // 图文链接消息
	MiniProgramPage *MiniProgramPageMessage `json:"miniprogrampage,omitempty"` // 小程序卡片消息
}

// TextMessage 文本消息
type TextMessage struct {
	Content string `json:"content"` // 文本消息内容
}

// ImageMessage 图片消息
type ImageMessage struct {
	MediaID string `json:"media_id"` // 发送的图片的媒体ID
}

// LinkMessage 图文链接消息
type LinkMessage struct {
	Title       string `json:"title"`       // 消息标题
	Description string `json:"description"` // 图文链接消息描述
	URL         string `json:"url"`         // 图文链接消息被点击后跳转的链接
	ThumbURL    string `json:"thumb_url"`   // 图文链接消息的图片链接，支持 JPG、PNG 格式，较好的效果为大图 640 X 320，小图 80 X 80
}

// MiniProgramPageMessage 小程序卡片消息
type MiniProgramPageMessage struct {
	Title        string `json:"title"`          // 消息标题
	PagePath     string `json:"pagepath"`       // 小程序的页面路径
	ThumbMediaID string `json:"thumb_media_id"` // 小程序消息卡片的封面图片素材id，必须是临时素材id
}

// Send 发送客服消息
func (c *Customer) Send(message *CustomerMessage) error {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", sendCustomerMessageURL, accessToken)

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
		return fmt.Errorf("发送客服消息失败: %s", result.ErrMsg)
	}

	return nil
}

// SendText 发送文本消息
func (c *Customer) SendText(toUser, content string) error {
	message := &CustomerMessage{
		ToUser:  toUser,
		MsgType: MessageTypeText,
		Text: &TextMessage{
			Content: content,
		},
	}

	return c.Send(message)
}

// SendImage 发送图片消息
func (c *Customer) SendImage(toUser, mediaID string) error {
	message := &CustomerMessage{
		ToUser:  toUser,
		MsgType: MessageTypeImage,
		Image: &ImageMessage{
			MediaID: mediaID,
		},
	}

	return c.Send(message)
}

// SendLink 发送图文链接消息
func (c *Customer) SendLink(toUser, title, description, url, thumbURL string) error {
	message := &CustomerMessage{
		ToUser:  toUser,
		MsgType: MessageTypeLink,
		Link: &LinkMessage{
			Title:       title,
			Description: description,
			URL:         url,
			ThumbURL:    thumbURL,
		},
	}

	return c.Send(message)
}

// SendMiniProgramPage 发送小程序卡片消息
func (c *Customer) SendMiniProgramPage(toUser, title, pagePath, thumbMediaID string) error {
	message := &CustomerMessage{
		ToUser:  toUser,
		MsgType: MessageTypeMiniProgramPage,
		MiniProgramPage: &MiniProgramPageMessage{
			Title:        title,
			PagePath:     pagePath,
			ThumbMediaID: thumbMediaID,
		},
	}

	return c.Send(message)
}

// GetTempMedia 获取客服消息内的临时素材
func (c *Customer) GetTempMedia(mediaID string) ([]byte, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s&media_id=%s", getTempMediaURL, accessToken, mediaID)

	return util.HTTPGetBytes(uri)
}

// SetTyping 下发客服当前输入状态给用户
func (c *Customer) SetTyping(toUser string, typing bool) error {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", setTypingURL, accessToken)

	data := struct {
		ToUser  string `json:"touser"`
		Command string `json:"command"`
	}{
		ToUser:  toUser,
		Command: "Typing",
	}
	if !typing {
		data.Command = "CancelTyping"
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
		return fmt.Errorf("设置输入状态失败: %s", result.ErrMsg)
	}

	return nil
}

func newCustomer(mp *MiniProgram) *Customer {
	return &Customer{mp}
}
