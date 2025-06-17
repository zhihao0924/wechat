package officialaccount

import (
	"fmt"
)

// SendKFTextMessage 发送客服文本消息
func (oa *OfficialAccount) SendKFTextMessage(openID, content, kfAccount string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	msg := struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Text    struct {
			Content string `json:"content"`
		} `json:"text"`
		CustomService *struct {
			KfAccount string `json:"kf_account"`
		} `json:"customservice,omitempty"`
	}{
		ToUser:  openID,
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
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

// SendKFImageMessage 发送客服图片消息
func (oa *OfficialAccount) SendKFImageMessage(openID, mediaID, kfAccount string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	msg := struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Image   struct {
			MediaID string `json:"media_id"`
		} `json:"image"`
		CustomService *struct {
			KfAccount string `json:"kf_account"`
		} `json:"customservice,omitempty"`
	}{
		ToUser:  openID,
		MsgType: "image",
		Image: struct {
			MediaID string `json:"media_id"`
		}{
			MediaID: mediaID,
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

// SendKFVoiceMessage 发送客服语音消息
func (oa *OfficialAccount) SendKFVoiceMessage(openID, mediaID, kfAccount string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	msg := struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Voice   struct {
			MediaID string `json:"media_id"`
		} `json:"voice"`
		CustomService *struct {
			KfAccount string `json:"kf_account"`
		} `json:"customservice,omitempty"`
	}{
		ToUser:  openID,
		MsgType: "voice",
		Voice: struct {
			MediaID string `json:"media_id"`
		}{
			MediaID: mediaID,
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

// SendKFVideoMessage 发送客服视频消息
func (oa *OfficialAccount) SendKFVideoMessage(openID, mediaID, thumbMediaID, title, description, kfAccount string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	msg := struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Video   struct {
			MediaID      string `json:"media_id"`
			ThumbMediaID string `json:"thumb_media_id"`
			Title        string `json:"title,omitempty"`
			Description  string `json:"description,omitempty"`
		} `json:"video"`
		CustomService *struct {
			KfAccount string `json:"kf_account"`
		} `json:"customservice,omitempty"`
	}{
		ToUser:  openID,
		MsgType: "video",
		Video: struct {
			MediaID      string `json:"media_id"`
			ThumbMediaID string `json:"thumb_media_id"`
			Title        string `json:"title,omitempty"`
			Description  string `json:"description,omitempty"`
		}{
			MediaID:      mediaID,
			ThumbMediaID: thumbMediaID,
			Title:        title,
			Description:  description,
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

// SendKFMusicMessage 发送客服音乐消息
func (oa *OfficialAccount) SendKFMusicMessage(openID, title, description, musicURL, hqMusicURL, thumbMediaID, kfAccount string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	msg := struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Music   struct {
			Title        string `json:"title,omitempty"`
			Description  string `json:"description,omitempty"`
			MusicURL     string `json:"musicurl"`
			HQMusicURL   string `json:"hqmusicurl"`
			ThumbMediaID string `json:"thumb_media_id"`
		} `json:"music"`
		CustomService *struct {
			KfAccount string `json:"kf_account"`
		} `json:"customservice,omitempty"`
	}{
		ToUser:  openID,
		MsgType: "music",
		Music: struct {
			Title        string `json:"title,omitempty"`
			Description  string `json:"description,omitempty"`
			MusicURL     string `json:"musicurl"`
			HQMusicURL   string `json:"hqmusicurl"`
			ThumbMediaID string `json:"thumb_media_id"`
		}{
			Title:        title,
			Description:  description,
			MusicURL:     musicURL,
			HQMusicURL:   hqMusicURL,
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

// Article 图文消息文章
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

// SendKFNewsMessage 发送客服图文消息
func (oa *OfficialAccount) SendKFNewsMessage(openID string, articles []Article, kfAccount string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	msg := struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		News    struct {
			Articles []Article `json:"articles"`
		} `json:"news"`
		CustomService *struct {
			KfAccount string `json:"kf_account"`
		} `json:"customservice,omitempty"`
	}{
		ToUser:  openID,
		MsgType: "news",
		News: struct {
			Articles []Article `json:"articles"`
		}{
			Articles: articles,
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

// SendKFMPNewsMessage 发送客服图文消息（点击跳转到图文消息页面）
func (oa *OfficialAccount) SendKFMPNewsMessage(openID, mediaID, kfAccount string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	msg := struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		MPNews  struct {
			MediaID string `json:"media_id"`
		} `json:"mpnews"`
		CustomService *struct {
			KfAccount string `json:"kf_account"`
		} `json:"customservice,omitempty"`
	}{
		ToUser:  openID,
		MsgType: "mpnews",
		MPNews: struct {
			MediaID string `json:"media_id"`
		}{
			MediaID: mediaID,
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

// MenuListItem 客服菜单项
type MenuListItem struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// SendKFMenuMessage 发送客服菜单消息
func (oa *OfficialAccount) SendKFMenuMessage(openID string, headContent string, list []MenuListItem, tailContent string, kfAccount string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	msg := struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		MsgMenu struct {
			HeadContent string         `json:"head_content"`
			List        []MenuListItem `json:"list"`
			TailContent string         `json:"tail_content"`
		} `json:"msgmenu"`
		CustomService *struct {
			KfAccount string `json:"kf_account"`
		} `json:"customservice,omitempty"`
	}{
		ToUser:  openID,
		MsgType: "msgmenu",
		MsgMenu: struct {
			HeadContent string         `json:"head_content"`
			List        []MenuListItem `json:"list"`
			TailContent string         `json:"tail_content"`
		}{
			HeadContent: headContent,
			List:        list,
			TailContent: tailContent,
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

// SendKFCardMessage 发送客服卡券消息
func (oa *OfficialAccount) SendKFCardMessage(openID, cardID, kfAccount string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	msg := struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		WxCard  struct {
			CardID string `json:"card_id"`
		} `json:"wxcard"`
		CustomService *struct {
			KfAccount string `json:"kf_account"`
		} `json:"customservice,omitempty"`
	}{
		ToUser:  openID,
		MsgType: "wxcard",
		WxCard: struct {
			CardID string `json:"card_id"`
		}{
			CardID: cardID,
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
