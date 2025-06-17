package officialaccount

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// ArticleItem 图文素材文章
type ArticleItem struct {
	Title              string `json:"title"`                           // 标题
	ThumbMediaID       string `json:"thumb_media_id"`                  // 图文消息的封面图片素材id
	Author             string `json:"author,omitempty"`                // 作者
	Digest             string `json:"digest,omitempty"`                // 图文消息的摘要
	Content            string `json:"content"`                         // 图文消息的具体内容
	URL                string `json:"url,omitempty"`                   // 图文页的URL
	ContentSourceURL   string `json:"content_source_url,omitempty"`    // 图文消息的原文地址
	ShowCoverPic       int    `json:"show_cover_pic"`                  // 是否显示封面，0为false，即不显示，1为true，即显示
	NeedOpenComment    int    `json:"need_open_comment,omitempty"`     // 是否打开评论，0不打开，1打开
	OnlyFansCanComment int    `json:"only_fans_can_comment,omitempty"` // 是否粉丝才可评论，0所有人可评论，1粉丝才可评论
}

// MediaResponse 上传媒体文件返回
type MediaResponse struct {
	Type      string `json:"type"`       // 媒体文件类型
	MediaID   string `json:"media_id"`   // 媒体文件ID
	CreatedAt int64  `json:"created_at"` // 媒体文件上传时间戳
	URL       string `json:"url"`        // 新增的永久素材的图片URL
}

// MaterialCount 素材总数
type MaterialCount struct {
	VoiceCount int `json:"voice_count"` // 语音总数量
	VideoCount int `json:"video_count"` // 视频总数量
	ImageCount int `json:"image_count"` // 图片总数量
	NewsCount  int `json:"news_count"`  // 图文总数量
}

// MaterialList 素材列表
type MaterialList struct {
	TotalCount int `json:"total_count"` // 该类型的素材的总数
	ItemCount  int `json:"item_count"`  // 本次调用获取的素材的数量
	Item       []struct {
		MediaID    string       `json:"media_id"`    // 素材ID
		Name       string       `json:"name"`        // 素材名称
		UpdateTime int64        `json:"update_time"` // 更新时间
		URL        string       `json:"url"`         // 素材URL
		Content    *ArticleItem `json:"content"`     // 图文素材内容
	} `json:"item"` // 素材列表
}

// UploadTempMedia 上传临时素材
// mediaType: 媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb）
// filename: 文件路径
func (oa *OfficialAccount) UploadTempMedia(mediaType, filename string) (*MediaResponse, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", accessToken, mediaType)

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建multipart表单
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("media", filepath.Base(filename))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// 发送请求
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var mediaResp MediaResponse
	err = json.Unmarshal(respBody, &mediaResp)
	if err != nil {
		return nil, err
	}

	return &mediaResp, nil
}

// GetTempMedia 获取临时素材
// mediaID: 媒体文件ID
// filename: 保存文件路径
func (oa *OfficialAccount) GetTempMedia(mediaID, filename string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s", accessToken, mediaID)

	// 发送请求
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查是否返回了错误信息
	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		var errResp struct {
			ErrCode int    `json:"errcode"`
			ErrMsg  string `json:"errmsg"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err == nil && errResp.ErrCode != 0 {
			return fmt.Errorf("微信API错误: %d %s", errResp.ErrCode, errResp.ErrMsg)
		}
	}

	// 创建文件
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// 写入文件
	_, err = io.Copy(out, resp.Body)
	return err
}

// AddMaterial 新增永久素材
// mediaType: 媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb）
// filename: 文件路径
// title, introduction: 视频素材需要title和introduction
func (oa *OfficialAccount) AddMaterial(mediaType, filename, title, introduction string) (*MediaResponse, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s", accessToken, mediaType)

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建multipart表单
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("media", filepath.Base(filename))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	// 如果是视频素材，需要额外的描述信息
	if mediaType == "video" {
		description := struct {
			Title        string `json:"title"`
			Introduction string `json:"introduction"`
		}{
			Title:        title,
			Introduction: introduction,
		}
		descBytes, err := json.Marshal(description)
		if err != nil {
			return nil, err
		}
		err = writer.WriteField("description", string(descBytes))
		if err != nil {
			return nil, err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// 发送请求
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var mediaResp MediaResponse
	err = json.Unmarshal(respBody, &mediaResp)
	if err != nil {
		return nil, err
	}

	return &mediaResp, nil
}

// AddNews 新增永久图文素材
func (oa *OfficialAccount) AddNews(articles []ArticleItem) (string, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=%s", accessToken)
	data := struct {
		Articles []ArticleItem `json:"articles"`
	}{
		Articles: articles,
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return "", err
	}

	var result struct {
		MediaID string `json:"media_id"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return "", err
	}

	return result.MediaID, nil
}

// GetMaterial 获取永久素材
// mediaID: 要获取的素材的media_id
func (oa *OfficialAccount) GetMaterial(mediaID string) ([]byte, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=%s", accessToken)
	data := struct {
		MediaID string `json:"media_id"`
	}{
		MediaID: mediaID,
	}

	return oa.httpPost(url, data)
}

// DeleteMaterial 删除永久素材
// mediaID: 要删除的素材的media_id
func (oa *OfficialAccount) DeleteMaterial(mediaID string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%s", accessToken)
	data := struct {
		MediaID string `json:"media_id"`
	}{
		MediaID: mediaID,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// GetMaterialCount 获取素材总数
func (oa *OfficialAccount) GetMaterialCount() (*MaterialCount, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=%s", accessToken)
	response, err := oa.httpGet(url)
	if err != nil {
		return nil, err
	}

	var count MaterialCount
	err = json.Unmarshal(response, &count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

// BatchGetMaterial 获取素材列表
// materialType: 素材的类型，图片（image）、视频（video）、语音 （voice）、图文（news）
// offset: 从全部素材的该偏移位置开始返回，0表示从第一个素材返回
// count: 返回素材的数量，取值在1到20之间
func (oa *OfficialAccount) BatchGetMaterial(materialType string, offset, count int) (*MaterialList, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%s", accessToken)
	data := struct {
		Type   string `json:"type"`
		Offset int    `json:"offset"`
		Count  int    `json:"count"`
	}{
		Type:   materialType,
		Offset: offset,
		Count:  count,
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var materialList MaterialList
	err = json.Unmarshal(response, &materialList)
	if err != nil {
		return nil, err
	}

	return &materialList, nil
}
