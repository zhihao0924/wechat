package miniprogram

import (
	"encoding/json"
	"fmt"

	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// Security 内容安全相关API
type Security struct {
	*MiniProgram
}

// 内容安全相关API地址
const (
	// 文本内容安全检测
	msgSecCheckURL = "https://api.weixin.qq.com/wxa/msg_sec_check"
	// 图片内容安全检测
	imgSecCheckURL = "https://api.weixin.qq.com/wxa/img_sec_check"
	// 音频内容安全检测
	mediaCheckAsyncURL = "https://api.weixin.qq.com/wxa/media_check_async"
)

// MediaType 媒体类型
type MediaType int

const (
	// MediaTypeAudio 音频
	MediaTypeAudio MediaType = 1
	// MediaTypeImage 图片
	MediaTypeImage MediaType = 2
)

// SecurityResult 安全检测结果
type SecurityResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Result  struct {
		Suggest string `json:"suggest"`
		Label   int    `json:"label"`
	} `json:"result"`
	TraceID string `json:"trace_id"`
}

// MediaCheckResult 媒体文件异步检测结果
type MediaCheckResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	TraceID string `json:"trace_id"`
}

// MsgSecCheck 检查一段文本是否含有违法违规内容
func (s *Security) MsgSecCheck(content string) (*SecurityResult, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", msgSecCheckURL, accessToken)

	data := struct {
		Content string `json:"content"`
	}{
		Content: content,
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return nil, err
	}

	var result SecurityResult
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return &result, fmt.Errorf("文本内容安全检测失败: %s", result.ErrMsg)
	}

	return &result, nil
}

// ImgSecCheck 校验一张图片是否含有违法违规内容
func (s *Security) ImgSecCheck(filePath string) (*SecurityResult, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", imgSecCheckURL, accessToken)

	response, err := util.PostFile("media", filePath, uri)
	if err != nil {
		return nil, err
	}

	var result SecurityResult
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return &result, fmt.Errorf("图片内容安全检测失败: %s", result.ErrMsg)
	}

	return &result, nil
}

// ImgSecCheckBytes 校验图片字节数据是否含有违法违规内容
func (s *Security) ImgSecCheckBytes(imageData []byte) (*SecurityResult, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", imgSecCheckURL, accessToken)

	response, err := util.PostFileBytes("media", "image.jpg", imageData, uri)
	if err != nil {
		return nil, err
	}

	var result SecurityResult
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return &result, fmt.Errorf("图片内容安全检测失败: %s", result.ErrMsg)
	}

	return &result, nil
}

// MediaCheckAsync 异步校验图片/音频是否含有违法违规内容
func (s *Security) MediaCheckAsync(mediaURL string, mediaType MediaType) (*MediaCheckResult, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", mediaCheckAsyncURL, accessToken)

	data := struct {
		MediaURL  string    `json:"media_url"`
		MediaType MediaType `json:"media_type"`
	}{
		MediaURL:  mediaURL,
		MediaType: mediaType,
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return nil, err
	}

	var result MediaCheckResult
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return &result, fmt.Errorf("媒体文件异步检测失败: %s", result.ErrMsg)
	}

	return &result, nil
}

// IsRisky 判断内容是否有风险
func (s *Security) IsRisky(result *SecurityResult) bool {
	if result == nil {
		return true
	}
	return result.Result.Suggest != "pass"
}

// GetRiskLevel 获取风险等级
func (s *Security) GetRiskLevel(result *SecurityResult) string {
	if result == nil {
		return "unknown"
	}

	switch result.Result.Suggest {
	case "pass":
		return "safe"
	case "risky":
		return "risky"
	case "review":
		return "review"
	default:
		return "unknown"
	}
}

// GetRiskType 获取风险类型
func (s *Security) GetRiskType(result *SecurityResult) string {
	if result == nil {
		return "unknown"
	}

	switch result.Result.Label {
	case 100:
		return "正常"
	case 10001:
		return "广告"
	case 20001:
		return "时政"
	case 20002:
		return "色情"
	case 20003:
		return "辱骂"
	case 20006:
		return "违法犯罪"
	case 20008:
		return "欺诈"
	case 20012:
		return "低俗"
	case 20013:
		return "版权"
	case 21000:
		return "其他"
	default:
		return fmt.Sprintf("未知类型(%d)", result.Result.Label)
	}
}

func newSecurity(mp *MiniProgram) *Security {
	return &Security{mp}
}
