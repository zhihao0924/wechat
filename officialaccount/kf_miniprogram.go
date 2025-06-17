package officialaccount

import (
	"fmt"
)

// SendKFMiniProgramMessage 发送客服小程序卡片消息
func (oa *OfficialAccount) SendKFMiniProgramMessage(openID, title, appID, pagePath, thumbMediaID, kfAccount string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	msg := struct {
		ToUser          string `json:"touser"`
		MsgType         string `json:"msgtype"`
		MiniProgramPage struct {
			Title        string `json:"title"`
			AppID        string `json:"appid"`
			PagePath     string `json:"pagepath"`
			ThumbMediaID string `json:"thumb_media_id"`
		} `json:"miniprogrampage"`
		CustomService *struct {
			KfAccount string `json:"kf_account"`
		} `json:"customservice,omitempty"`
	}{
		ToUser:  openID,
		MsgType: "miniprogrampage",
		MiniProgramPage: struct {
			Title        string `json:"title"`
			AppID        string `json:"appid"`
			PagePath     string `json:"pagepath"`
			ThumbMediaID string `json:"thumb_media_id"`
		}{
			Title:        title,
			AppID:        appID,
			PagePath:     pagePath,
			ThumbMediaID: thumbMediaID,
		},
	}

	// 如果指定了客服账号，则添加到消息中
	if kfAccount != "" {
		msg.CustomService = &struct {
			KfAccount string `json:"kf_account"`
		}{
			KfAccount: kfAccount,
		}
	}

	_, err = oa.httpPost(url, msg)
	return err
}

// SendKFMiniProgramLinkMessage 发送客服小程序链接消息
func (oa *OfficialAccount) SendKFMiniProgramLinkMessage(openID, title, description, url, thumbURL, kfAccount string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	apiURL := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	msg := struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Link    struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			ThumbURL    string `json:"thumb_url"`
		} `json:"link"`
		CustomService *struct {
			KfAccount string `json:"kf_account"`
		} `json:"customservice,omitempty"`
	}{
		ToUser:  openID,
		MsgType: "link",
		Link: struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			ThumbURL    string `json:"thumb_url"`
		}{
			Title:       title,
			Description: description,
			URL:         url,
			ThumbURL:    thumbURL,
		},
	}

	// 如果指定了客服账号，则添加到消息中
	if kfAccount != "" {
		msg.CustomService = &struct {
			KfAccount string `json:"kf_account"`
		}{
			KfAccount: kfAccount,
		}
	}

	_, err = oa.httpPost(apiURL, msg)
	return err
}

// SendKFMiniProgramLocationMessage 发送客服小程序位置消息
func (oa *OfficialAccount) SendKFMiniProgramLocationMessage(openID string, title string, latitude, longitude float64, kfAccount string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	msg := struct {
		ToUser   string `json:"touser"`
		MsgType  string `json:"msgtype"`
		Location struct {
			Title     string  `json:"title"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
		CustomService *struct {
			KfAccount string `json:"kf_account"`
		} `json:"customservice,omitempty"`
	}{
		ToUser:  openID,
		MsgType: "location",
		Location: struct {
			Title     string  `json:"title"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Title:     title,
			Latitude:  latitude,
			Longitude: longitude,
		},
	}

	// 如果指定了客服账号，则添加到消息中
	if kfAccount != "" {
		msg.CustomService = &struct {
			KfAccount string `json:"kf_account"`
		}{
			KfAccount: kfAccount,
		}
	}

	_, err = oa.httpPost(url, msg)
	return err
}

// SendKFMiniProgramTemplateMessage 发送客服小程序模板消息
func (oa *OfficialAccount) SendKFMiniProgramTemplateMessage(openID, templateID, page string, formID string, data map[string]struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}, emphasisKeyword string, kfAccount string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	msg := struct {
		ToUser        string `json:"touser"`
		MsgType       string `json:"msgtype"`
		WeappTemplate struct {
			TemplateID string `json:"template_id"`
			Page       string `json:"page"`
			FormID     string `json:"form_id"`
			Data       map[string]struct {
				Value string `json:"value"`
				Color string `json:"color,omitempty"`
			} `json:"data"`
			EmphasisKeyword string `json:"emphasis_keyword,omitempty"`
		} `json:"weapptemplate"`
		CustomService *struct {
			KfAccount string `json:"kf_account"`
		} `json:"customservice,omitempty"`
	}{
		ToUser:  openID,
		MsgType: "weapptemplate",
		WeappTemplate: struct {
			TemplateID string `json:"template_id"`
			Page       string `json:"page"`
			FormID     string `json:"form_id"`
			Data       map[string]struct {
				Value string `json:"value"`
				Color string `json:"color,omitempty"`
			} `json:"data"`
			EmphasisKeyword string `json:"emphasis_keyword,omitempty"`
		}{
			TemplateID:      templateID,
			Page:            page,
			FormID:          formID,
			Data:            data,
			EmphasisKeyword: emphasisKeyword,
		},
	}

	// 如果指定了客服账号，则添加到消息中
	if kfAccount != "" {
		msg.CustomService = &struct {
			KfAccount string `json:"kf_account"`
		}{
			KfAccount: kfAccount,
		}
	}

	_, err = oa.httpPost(url, msg)
	return err
}
