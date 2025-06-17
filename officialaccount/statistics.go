package officialaccount

import (
	"encoding/json"
	"fmt"
	"time"
)

// UserSummary 用户增减数据
type UserSummary struct {
	RefDate    string `json:"ref_date"`    // 数据的日期
	UserSource int    `json:"user_source"` // 用户的渠道，0代表其他合计，1代表公众号搜索，17代表名片分享，30代表扫描二维码
	NewUser    int    `json:"new_user"`    // 新增的用户数量
	CancelUser int    `json:"cancel_user"` // 取消关注的用户数量
}

// UserCumulate 累计用户数据
type UserCumulate struct {
	RefDate      string `json:"ref_date"`      // 数据的日期
	CumulateUser int    `json:"cumulate_user"` // 总用户量
}

// ArticleSummary 图文群发每日数据
type ArticleSummary struct {
	RefDate          string `json:"ref_date"`            // 数据的日期
	MsgID            string `json:"msg_id"`              // 消息ID
	Title            string `json:"title"`               // 图文消息的标题
	IntPageReadUser  int    `json:"int_page_read_user"`  // 图文页（点击群发图文卡片进入的页面）的阅读人数
	IntPageReadCount int    `json:"int_page_read_count"` // 图文页的阅读次数
	OriPageReadUser  int    `json:"ori_page_read_user"`  // 原文页（点击图文页"阅读原文"进入的页面）的阅读人数
	OriPageReadCount int    `json:"ori_page_read_count"` // 原文页的阅读次数
	ShareUser        int    `json:"share_user"`          // 分享的人数
	ShareCount       int    `json:"share_count"`         // 分享的次数
	AddToFavUser     int    `json:"add_to_fav_user"`     // 收藏的人数
	AddToFavCount    int    `json:"add_to_fav_count"`    // 收藏的次数
}

// ArticleTotal 图文群发总数据
type ArticleTotal struct {
	RefDate string `json:"ref_date"` // 数据的日期
	MsgID   string `json:"msg_id"`   // 消息ID
	Title   string `json:"title"`    // 图文消息的标题
	Details []struct {
		StatDate         string `json:"stat_date"`           // 统计的日期
		TargetUser       int    `json:"target_user"`         // 送达人数
		IntPageReadUser  int    `json:"int_page_read_user"`  // 图文页的阅读人数
		IntPageReadCount int    `json:"int_page_read_count"` // 图文页的阅读次数
		OriPageReadUser  int    `json:"ori_page_read_user"`  // 原文页的阅读人数
		OriPageReadCount int    `json:"ori_page_read_count"` // 原文页的阅读次数
		ShareUser        int    `json:"share_user"`          // 分享的人数
		ShareCount       int    `json:"share_count"`         // 分享的次数
		AddToFavUser     int    `json:"add_to_fav_user"`     // 收藏的人数
		AddToFavCount    int    `json:"add_to_fav_count"`    // 收藏的次数
	} `json:"details"`
}

// UserRead 图文统计数据
type UserRead struct {
	RefDate          string `json:"ref_date"`            // 数据的日期
	UserSource       int    `json:"user_source"`         // 用户的渠道
	IntPageReadUser  int    `json:"int_page_read_user"`  // 图文页的阅读人数
	IntPageReadCount int    `json:"int_page_read_count"` // 图文页的阅读次数
	OriPageReadUser  int    `json:"ori_page_read_user"`  // 原文页的阅读人数
	OriPageReadCount int    `json:"ori_page_read_count"` // 原文页的阅读次数
	ShareUser        int    `json:"share_user"`          // 分享的人数
	ShareCount       int    `json:"share_count"`         // 分享的次数
	AddToFavUser     int    `json:"add_to_fav_user"`     // 收藏的人数
	AddToFavCount    int    `json:"add_to_fav_count"`    // 收藏的次数
}

// UserShare 图文分享转发数据
type UserShare struct {
	RefDate    string `json:"ref_date"`    // 数据的日期
	ShareScene int    `json:"share_scene"` // 分享的场景，1代表好友转发，2代表朋友圈，3代表腾讯微博，255代表其他
	ShareCount int    `json:"share_count"` // 分享的次数
	ShareUser  int    `json:"share_user"`  // 分享的人数
}

// UpstreamMsg 消息发送概况数据
type UpstreamMsg struct {
	RefDate  string `json:"ref_date"`  // 数据的日期
	MsgType  int    `json:"msg_type"`  // 消息类型，1代表文字，2代表图片，3代表语音，4代表视频，6代表第三方应用消息
	MsgUser  int    `json:"msg_user"`  // 上行发送了消息的用户数
	MsgCount int    `json:"msg_count"` // 上行发送了消息的消息总数
}

// UpstreamMsgDist 消息分布数据
type UpstreamMsgDist struct {
	RefDate       string `json:"ref_date"`       // 数据的日期
	CountInterval int    `json:"count_interval"` // 当日发送消息量分布的区间，0代表0，1代表1-5，2代表6-10，3代表10次以上
	MsgUser       int    `json:"msg_user"`       // 上行发送了消息的用户数
}

// InterfaceSummary 接口分析数据
type InterfaceSummary struct {
	RefDate       string `json:"ref_date"`        // 数据的日期
	CallbackCount int    `json:"callback_count"`  // 通过服务器配置地址获得消息后，被动回复用户消息的次数
	FailCount     int    `json:"fail_count"`      // 上述动作的失败次数
	TotalTimeCost int    `json:"total_time_cost"` // 总耗时，除以callback_count即为平均耗时
	MaxTimeCost   int    `json:"max_time_cost"`   // 最大耗时
}

// InterfaceSummaryHour 接口分析分时数据
type InterfaceSummaryHour struct {
	RefDate       string `json:"ref_date"`        // 数据的日期
	RefHour       int    `json:"ref_hour"`        // 数据的小时
	CallbackCount int    `json:"callback_count"`  // 通过服务器配置地址获得消息后，被动回复用户消息的次数
	FailCount     int    `json:"fail_count"`      // 上述动作的失败次数
	TotalTimeCost int    `json:"total_time_cost"` // 总耗时，除以callback_count即为平均耗时
	MaxTimeCost   int    `json:"max_time_cost"`   // 最大耗时
}

// formatDate 格式化日期为 "2006-01-02" 格式
func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// GetUserSummary 获取用户增减数据
// beginDate, endDate: 获取数据的起始日期和结束日期，格式为 time.Time
func (oa *OfficialAccount) GetUserSummary(beginDate, endDate time.Time) ([]UserSummary, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/datacube/getusersummary?access_token=%s", accessToken)
	data := struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: formatDate(beginDate),
		EndDate:   formatDate(endDate),
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result struct {
		List []UserSummary `json:"list"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.List, nil
}

// GetUserCumulate 获取累计用户数据
// beginDate, endDate: 获取数据的起始日期和结束日期，格式为 time.Time
func (oa *OfficialAccount) GetUserCumulate(beginDate, endDate time.Time) ([]UserCumulate, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/datacube/getusercumulate?access_token=%s", accessToken)
	data := struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: formatDate(beginDate),
		EndDate:   formatDate(endDate),
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result struct {
		List []UserCumulate `json:"list"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.List, nil
}

// GetArticleSummary 获取图文群发每日数据
// beginDate, endDate: 获取数据的起始日期和结束日期，格式为 time.Time
func (oa *OfficialAccount) GetArticleSummary(beginDate, endDate time.Time) ([]ArticleSummary, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/datacube/getarticlesummary?access_token=%s", accessToken)
	data := struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: formatDate(beginDate),
		EndDate:   formatDate(endDate),
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result struct {
		List []ArticleSummary `json:"list"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.List, nil
}

// GetArticleTotal 获取图文群发总数据
// beginDate, endDate: 获取数据的起始日期和结束日期，格式为 time.Time
func (oa *OfficialAccount) GetArticleTotal(beginDate, endDate time.Time) ([]ArticleTotal, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/datacube/getarticletotal?access_token=%s", accessToken)
	data := struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: formatDate(beginDate),
		EndDate:   formatDate(endDate),
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result struct {
		List []ArticleTotal `json:"list"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.List, nil
}

// GetUserRead 获取图文统计数据
// beginDate, endDate: 获取数据的起始日期和结束日期，格式为 time.Time
func (oa *OfficialAccount) GetUserRead(beginDate, endDate time.Time) ([]UserRead, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/datacube/getuserread?access_token=%s", accessToken)
	data := struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: formatDate(beginDate),
		EndDate:   formatDate(endDate),
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result struct {
		List []UserRead `json:"list"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.List, nil
}

// GetUserShare 获取图文分享转发数据
// beginDate, endDate: 获取数据的起始日期和结束日期，格式为 time.Time
func (oa *OfficialAccount) GetUserShare(beginDate, endDate time.Time) ([]UserShare, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/datacube/getusershare?access_token=%s", accessToken)
	data := struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: formatDate(beginDate),
		EndDate:   formatDate(endDate),
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result struct {
		List []UserShare `json:"list"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.List, nil
}
