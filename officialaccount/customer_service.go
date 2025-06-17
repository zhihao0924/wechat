package officialaccount

import (
	"encoding/json"
	"fmt"
	"time"
)

// KFAccount 客服账号
type KFAccount struct {
	KfAccount    string `json:"kf_account"`    // 完整客服账号，格式为：账号前缀@公众号微信号
	KfNick       string `json:"kf_nick"`       // 客服昵称
	KfID         string `json:"kf_id"`         // 客服编号
	KfHeadImgURL string `json:"kf_headimgurl"` // 客服头像URL
}

// KFOnline 在线客服
type KFOnline struct {
	KfAccount    string `json:"kf_account"`    // 完整客服账号，格式为：账号前缀@公众号微信号
	Status       int    `json:"status"`        // 客服在线状态，目前为：1、web在线，2、手机在线，3、PC在线
	AutoAccept   int    `json:"auto_accept"`   // 客服设置的最大自动接入数
	AcceptedCase int    `json:"accepted_case"` // 客服当前正在接待的会话数
}

// KFSession 客服会话
type KFSession struct {
	CreateTime int64  `json:"createtime"` // 会话创建时间
	OpenID     string `json:"openid"`     // 粉丝openid
}

// KFSessionList 客服会话列表
type KFSessionList struct {
	SessionList []KFSession `json:"sessionlist"` // 会话列表
}

// KFWaitCase 未接入会话
type KFWaitCase struct {
	LatestTime int64  `json:"latest_time"` // 粉丝的最后一条消息的时间
	OpenID     string `json:"openid"`      // 粉丝openid
}

// KFWaitCaseList 未接入会话列表
type KFWaitCaseList struct {
	Count        int          `json:"count"`        // 未接入会话数量
	WaitCaseList []KFWaitCase `json:"waitcaselist"` // 未接入会话列表
}

// KFMsgRecord 客服消息记录
type KFMsgRecord struct {
	Worker   string `json:"worker"`   // 客服账号
	OpenID   string `json:"openid"`   // 用户的openid
	OperCode int    `json:"opercode"` // 操作码，2002（客服发送信息），2003（客服接收消息）
	Text     string `json:"text"`     // 聊天记录
	Time     int64  `json:"time"`     // 操作时间，unix时间戳
}

// KFMsgRecordList 客服消息记录列表
type KFMsgRecordList struct {
	RecordList []KFMsgRecord `json:"recordlist"` // 消息记录列表
	Number     int           `json:"number"`     // 记录数目
	MsgID      int64         `json:"msgid"`      // 消息id，下次请求时用于分页
}

// AddKFAccount 添加客服账号
func (oa *OfficialAccount) AddKFAccount(account, nickname, password string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%s", accessToken)
	data := struct {
		KfAccount string `json:"kf_account"`
		Nickname  string `json:"nickname"`
		Password  string `json:"password"`
	}{
		KfAccount: account,
		Nickname:  nickname,
		Password:  password,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// UpdateKFAccount 修改客服账号
func (oa *OfficialAccount) UpdateKFAccount(account, nickname, password string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfaccount/update?access_token=%s", accessToken)
	data := struct {
		KfAccount string `json:"kf_account"`
		Nickname  string `json:"nickname"`
		Password  string `json:"password"`
	}{
		KfAccount: account,
		Nickname:  nickname,
		Password:  password,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// DeleteKFAccount 删除客服账号
func (oa *OfficialAccount) DeleteKFAccount(account, nickname, password string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfaccount/del?access_token=%s", accessToken)
	data := struct {
		KfAccount string `json:"kf_account"`
		Nickname  string `json:"nickname"`
		Password  string `json:"password"`
	}{
		KfAccount: account,
		Nickname:  nickname,
		Password:  password,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// GetKFList 获取所有客服账号
func (oa *OfficialAccount) GetKFList() ([]KFAccount, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=%s", accessToken)
	response, err := oa.httpGet(url)
	if err != nil {
		return nil, err
	}

	var result struct {
		KfList []KFAccount `json:"kf_list"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.KfList, nil
}

// GetOnlineKFList 获取在线客服账号
func (oa *OfficialAccount) GetOnlineKFList() ([]KFOnline, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=%s", accessToken)
	response, err := oa.httpGet(url)
	if err != nil {
		return nil, err
	}

	var result struct {
		KfOnlineList []KFOnline `json:"kf_online_list"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.KfOnlineList, nil
}

// UploadKFHeadImg 上传客服头像
// account: 完整客服账号，格式为：账号前缀@公众号微信号
// filename: 头像图片文件路径
func (oa *OfficialAccount) UploadKFHeadImg(account, filename string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=%s&kf_account=%s", accessToken, account)

	// 使用素材管理中的上传文件方法
	_, err = oa.uploadMedia(url, filename)
	return err
}

// CreateKFSession 创建会话
// account: 完整客服账号，格式为：账号前缀@公众号微信号
// openID: 粉丝的openid
func (oa *OfficialAccount) CreateKFSession(account, openID string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfsession/create?access_token=%s", accessToken)
	data := struct {
		KfAccount string `json:"kf_account"`
		OpenID    string `json:"openid"`
	}{
		KfAccount: account,
		OpenID:    openID,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// CloseKFSession 关闭会话
// account: 完整客服账号，格式为：账号前缀@公众号微信号
// openID: 粉丝的openid
func (oa *OfficialAccount) CloseKFSession(account, openID string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfsession/close?access_token=%s", accessToken)
	data := struct {
		KfAccount string `json:"kf_account"`
		OpenID    string `json:"openid"`
	}{
		KfAccount: account,
		OpenID:    openID,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// GetKFSession 获取客户会话状态
// openID: 粉丝的openid
func (oa *OfficialAccount) GetKFSession(openID string) (*KFSession, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfsession/getsession?access_token=%s&openid=%s", accessToken, openID)
	response, err := oa.httpGet(url)
	if err != nil {
		return nil, err
	}

	var session KFSession
	err = json.Unmarshal(response, &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

// GetKFSessionList 获取客服的会话列表
// account: 完整客服账号，格式为：账号前缀@公众号微信号
func (oa *OfficialAccount) GetKFSessionList(account string) (*KFSessionList, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfsession/getsessionlist?access_token=%s&kf_account=%s", accessToken, account)
	response, err := oa.httpGet(url)
	if err != nil {
		return nil, err
	}

	var sessionList KFSessionList
	err = json.Unmarshal(response, &sessionList)
	if err != nil {
		return nil, err
	}

	return &sessionList, nil
}

// GetKFWaitCaseList 获取未接入会话列表
func (oa *OfficialAccount) GetKFWaitCaseList() (*KFWaitCaseList, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfsession/getwaitcase?access_token=%s", accessToken)
	response, err := oa.httpGet(url)
	if err != nil {
		return nil, err
	}

	var waitCaseList KFWaitCaseList
	err = json.Unmarshal(response, &waitCaseList)
	if err != nil {
		return nil, err
	}

	return &waitCaseList, nil
}

// GetKFMsgRecords 获取客服聊天记录
// startTime, endTime: 查询时间范围
// msgID: 消息id顺序从小到大，从1开始
// number: 每次获取条数，最多10000条
func (oa *OfficialAccount) GetKFMsgRecords(startTime, endTime time.Time, msgID int64, number int) (*KFMsgRecordList, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/customservice/msgrecord/getmsglist?access_token=%s", accessToken)
	data := struct {
		StartTime int64 `json:"starttime"`
		EndTime   int64 `json:"endtime"`
		MsgID     int64 `json:"msgid"`
		Number    int   `json:"number"`
	}{
		StartTime: startTime.Unix(),
		EndTime:   endTime.Unix(),
		MsgID:     msgID,
		Number:    number,
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var recordList KFMsgRecordList
	err = json.Unmarshal(response, &recordList)
	if err != nil {
		return nil, err
	}

	return &recordList, nil
}
