package officialaccount

import (
	"encoding/xml"
	"time"
)

// MixMessage 存放所有微信发送过来的消息和事件
type MixMessage struct {
	CommonToken

	// 基本消息
	MsgID        int64   `xml:"MsgId"`
	Content      string  `xml:"Content"`
	PicURL       string  `xml:"PicUrl"`
	MediaID      string  `xml:"MediaId"`
	Format       string  `xml:"Format"`
	Recognition  string  `xml:"Recognition"`
	ThumbMediaID string  `xml:"ThumbMediaId"`
	LocationX    float64 `xml:"Location_X"`
	LocationY    float64 `xml:"Location_Y"`
	Scale        float64 `xml:"Scale"`
	Label        string  `xml:"Label"`
	Title        string  `xml:"Title"`
	Description  string  `xml:"Description"`
	URL          string  `xml:"Url"`

	// 事件相关
	Event     string `xml:"Event"`
	EventKey  string `xml:"EventKey"`
	Ticket    string `xml:"Ticket"`
	Latitude  string `xml:"Latitude"`
	Longitude string `xml:"Longitude"`
	Precision string `xml:"Precision"`

	// 菜单相关
	MenuID string `xml:"MenuId"`

	// 扫码相关
	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"`
		ScanResult string `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"`

	// 发送图片信息
	SendPicsInfo struct {
		Count   int32      `xml:"Count"`
		PicList []EventPic `xml:"PicList>item"`
	} `xml:"SendPicsInfo"`

	// 发送位置信息
	SendLocationInfo struct {
		LocationX float64 `xml:"Location_X"`
		LocationY float64 `xml:"Location_Y"`
		Scale     float64 `xml:"Scale"`
		Label     string  `xml:"Label"`
		Poiname   string  `xml:"Poiname"`
	}
}

// CommonToken 消息中通用的结构
type CommonToken struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
}

// EventPic 发送图片信息
type EventPic struct {
	PicMd5Sum string `xml:"PicMd5Sum"`
}

// ReplyMessage 回复消息的基础结构
type ReplyMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
}

// CDATA 定义XML CDATA数据类型
type CDATA struct {
	Text string `xml:",cdata"`
}

// ReplyText 文本消息回复
type ReplyText struct {
	ReplyMessage
	Content CDATA `xml:"Content"`
}

// ReplyImage 图片消息回复
type ReplyImage struct {
	ReplyMessage
	Image struct {
		MediaID CDATA `xml:"MediaId"`
	} `xml:"Image"`
}

// ReplyVoice 语音消息回复
type ReplyVoice struct {
	ReplyMessage
	Voice struct {
		MediaID CDATA `xml:"MediaId"`
	} `xml:"Voice"`
}

// ReplyVideo 视频消息回复
type ReplyVideo struct {
	ReplyMessage
	Video struct {
		MediaID     CDATA `xml:"MediaId"`
		Title       CDATA `xml:"Title,omitempty"`
		Description CDATA `xml:"Description,omitempty"`
	} `xml:"Video"`
}

// ReplyMusic 音乐消息回复
type ReplyMusic struct {
	ReplyMessage
	Music struct {
		Title        CDATA `xml:"Title,omitempty"`
		Description  CDATA `xml:"Description,omitempty"`
		MusicURL     CDATA `xml:"MusicUrl,omitempty"`
		HQMusicURL   CDATA `xml:"HQMusicUrl,omitempty"`
		ThumbMediaID CDATA `xml:"ThumbMediaId"`
	} `xml:"Music"`
}

// ReplyNews 图文消息回复
type ReplyNews struct {
	ReplyMessage
	ArticleCount int        `xml:"ArticleCount"`
	Articles     []NewsItem `xml:"Articles>item,omitempty"`
}

// NewsItem 图文消息项
type NewsItem struct {
	Title       CDATA `xml:"Title"`
	Description CDATA `xml:"Description"`
	PicURL      CDATA `xml:"PicUrl"`
	URL         CDATA `xml:"Url"`
}

// NewReplyText 创建文本回复
func NewReplyText(to, from, content string) *ReplyText {
	return &ReplyText{
		ReplyMessage: ReplyMessage{
			ToUserName:   CDATA{Text: to},
			FromUserName: CDATA{Text: from},
			CreateTime:   time.Now().Unix(),
			MsgType:      CDATA{Text: "text"},
		},
		Content: CDATA{Text: content},
	}
}

// NewReplyImage 创建图片回复
func NewReplyImage(to, from, mediaID string) *ReplyImage {
	reply := &ReplyImage{
		ReplyMessage: ReplyMessage{
			ToUserName:   CDATA{Text: to},
			FromUserName: CDATA{Text: from},
			CreateTime:   time.Now().Unix(),
			MsgType:      CDATA{Text: "image"},
		},
	}
	reply.Image.MediaID = CDATA{Text: mediaID}
	return reply
}

// NewReplyVoice 创建语音回复
func NewReplyVoice(to, from, mediaID string) *ReplyVoice {
	reply := &ReplyVoice{
		ReplyMessage: ReplyMessage{
			ToUserName:   CDATA{Text: to},
			FromUserName: CDATA{Text: from},
			CreateTime:   time.Now().Unix(),
			MsgType:      CDATA{Text: "voice"},
		},
	}
	reply.Voice.MediaID = CDATA{Text: mediaID}
	return reply
}

// NewReplyVideo 创建视频回复
func NewReplyVideo(to, from, mediaID, title, description string) *ReplyVideo {
	reply := &ReplyVideo{
		ReplyMessage: ReplyMessage{
			ToUserName:   CDATA{Text: to},
			FromUserName: CDATA{Text: from},
			CreateTime:   time.Now().Unix(),
			MsgType:      CDATA{Text: "video"},
		},
	}
	reply.Video.MediaID = CDATA{Text: mediaID}
	reply.Video.Title = CDATA{Text: title}
	reply.Video.Description = CDATA{Text: description}
	return reply
}

// NewReplyNews 创建图文回复
func NewReplyNews(to, from string, articles []NewsItem) *ReplyNews {
	reply := &ReplyNews{
		ReplyMessage: ReplyMessage{
			ToUserName:   CDATA{Text: to},
			FromUserName: CDATA{Text: from},
			CreateTime:   time.Now().Unix(),
			MsgType:      CDATA{Text: "news"},
		},
		ArticleCount: len(articles),
		Articles:     articles,
	}
	return reply
}
