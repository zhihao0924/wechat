package miniprogram

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// Analysis 数据分析相关API
type Analysis struct {
	*MiniProgram
}

// 数据分析相关API地址
const (
	// 概况趋势
	dailySummaryURL = "https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend"
	// 访问趋势
	dailyVisitTrendURL   = "https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend"
	weeklyVisitTrendURL  = "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyvisittrend"
	monthlyVisitTrendURL = "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyvisittrend"
	// 访问分布
	visitDistributionURL = "https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution"
	// 访问页面
	visitPageURL = "https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage"
	// 用户画像
	userPortraitURL = "https://api.weixin.qq.com/datacube/getweanalysisappiduserportrait"
)

// DateRange 日期范围
type DateRange struct {
	BeginDate string `json:"begin_date"` // 开始日期，格式为 yyyymmdd
	EndDate   string `json:"end_date"`   // 结束日期，格式为 yyyymmdd
}

// DailySummary 概况趋势数据项
type DailySummary struct {
	RefDate    string `json:"ref_date"`    // 日期，格式为 yyyymmdd
	VisitTotal int    `json:"visit_total"` // 累计用户数
	SharePV    int    `json:"share_pv"`    // 转发次数
	ShareUV    int    `json:"share_uv"`    // 转发人数
}

// DailyVisitTrend 访问趋势数据项
type DailyVisitTrend struct {
	RefDate         string  `json:"ref_date"`          // 日期，格式为 yyyymmdd
	SessionCnt      int     `json:"session_cnt"`       // 打开次数
	VisitPV         int     `json:"visit_pv"`          // 访问次数
	VisitUV         int     `json:"visit_uv"`          // 访问人数
	VisitUVNew      int     `json:"visit_uv_new"`      // 新用户数
	StayTimeUV      float64 `json:"stay_time_uv"`      // 人均停留时长 (秒)
	StayTimeSession float64 `json:"stay_time_session"` // 次均停留时长 (秒)
	VisitDepth      float64 `json:"visit_depth"`       // 平均访问深度
}

// VisitDistribution 访问分布数据
type VisitDistribution struct {
	RefDate string `json:"ref_date"` // 日期，格式为 yyyymmdd
	// 访问来源分布
	AccessSourceList []struct {
		AccessSource string  `json:"access_source"` // 访问来源
		SessionCnt   int     `json:"session_cnt"`   // 打开次数
		SessionPct   float64 `json:"session_pct"`   // 访问占比
	} `json:"access_source_session_cnt"`
	// 访问时长分布
	AccessStayTimeList []struct {
		StayTime   string  `json:"stay_time_info"` // 访问时长，0-2s等区间
		SessionCnt int     `json:"session_cnt"`    // 打开次数
		SessionPct float64 `json:"session_pct"`    // 访问占比
	} `json:"access_staytime_info"`
	// 访问深度分布
	AccessDepthList []struct {
		VisitDepth string  `json:"visit_depth_info"` // 访问深度，1页，2页等
		SessionCnt int     `json:"session_cnt"`      // 打开次数
		SessionPct float64 `json:"session_pct"`      // 访问占比
	} `json:"access_depth_info"`
}

// VisitPage 访问页面数据项
type VisitPage struct {
	RefDate      string  `json:"ref_date"`      // 日期，格式为 yyyymmdd
	PagePath     string  `json:"page_path"`     // 页面路径
	PageVisitPV  int     `json:"page_visit_pv"` // 访问次数
	PageVisitUV  int     `json:"page_visit_uv"` // 访问人数
	PageStayTime float64 `json:"page_staytime"` // 平均停留时长
	EntryPagePV  int     `json:"entrypage_pv"`  // 进入页次数
	ExitPagePV   int     `json:"exitpage_pv"`   // 退出页次数
	PageSharePV  int     `json:"page_share_pv"` // 转发次数
	PageShareUV  int     `json:"page_share_uv"` // 转发人数
}

// UserPortrait 用户画像数据
type UserPortrait struct {
	RefDate string `json:"ref_date"` // 日期，格式为 yyyymmdd
	// 性别分布
	GenderList []struct {
		ID    int     `json:"id"`    // 性别，0未知，1男性，2女性
		Name  string  `json:"name"`  // 性别名称
		Value int     `json:"value"` // 用户数
		Ratio float64 `json:"ratio"` // 占比
	} `json:"visit_uv_gender"`
	// 年龄分布
	AgeList []struct {
		ID    int     `json:"id"`    // 年龄，0未知，1-6分别代表17岁以下、18-24岁、25-29岁、30-39岁、40-49岁、50岁以上
		Name  string  `json:"name"`  // 年龄段名称
		Value int     `json:"value"` // 用户数
		Ratio float64 `json:"ratio"` // 占比
	} `json:"visit_uv_age"`
	// 地域分布
	ProvinceList []struct {
		ID    int     `json:"id"`    // 省份ID
		Name  string  `json:"name"`  // 省份名称
		Value int     `json:"value"` // 用户数
		Ratio float64 `json:"ratio"` // 占比
	} `json:"visit_uv_province"`
	// 终端分布
	PlatformList []struct {
		ID    int     `json:"id"`    // 终端类型，1iOS，2安卓，3其他
		Name  string  `json:"name"`  // 终端名称
		Value int     `json:"value"` // 用户数
		Ratio float64 `json:"ratio"` // 占比
	} `json:"visit_uv_platform"`
}

// GetDailySummary 获取小程序概况趋势
func (a *Analysis) GetDailySummary(beginDate, endDate time.Time) ([]DailySummary, error) {
	return a.getData(dailySummaryURL, beginDate, endDate, []DailySummary{})
}

// GetDailyVisitTrend 获取小程序日访问趋势
func (a *Analysis) GetDailyVisitTrend(beginDate, endDate time.Time) ([]DailyVisitTrend, error) {
	return a.getData(dailyVisitTrendURL, beginDate, endDate, []DailyVisitTrend{})
}

// GetWeeklyVisitTrend 获取小程序周访问趋势
func (a *Analysis) GetWeeklyVisitTrend(beginDate, endDate time.Time) ([]DailyVisitTrend, error) {
	return a.getData(weeklyVisitTrendURL, beginDate, endDate, []DailyVisitTrend{})
}

// GetMonthlyVisitTrend 获取小程序月访问趋势
func (a *Analysis) GetMonthlyVisitTrend(beginDate, endDate time.Time) ([]DailyVisitTrend, error) {
	return a.getData(monthlyVisitTrendURL, beginDate, endDate, []DailyVisitTrend{})
}

// GetVisitDistribution 获取小程序访问分布
func (a *Analysis) GetVisitDistribution(date time.Time) (*VisitDistribution, error) {
	dateStr := date.Format("20060102")
	result := &VisitDistribution{}
	err := a.getDataForDate(visitDistributionURL, dateStr, result)
	return result, err
}

// GetVisitPage 获取小程序访问页面
func (a *Analysis) GetVisitPage(beginDate, endDate time.Time) ([]VisitPage, error) {
	return a.getData(visitPageURL, beginDate, endDate, []VisitPage{})
}

// GetUserPortrait 获取小程序用户画像
func (a *Analysis) GetUserPortrait(beginDate, endDate time.Time) (*UserPortrait, error) {
	beginDateStr := beginDate.Format("20060102")
	endDateStr := endDate.Format("20060102")
	result := &UserPortrait{}
	err := a.getDataForDateRange(userPortraitURL, beginDateStr, endDateStr, result)
	return result, err
}

// getData 获取数据通用方法
func (a *Analysis) getData(url string, beginDate, endDate time.Time, result interface{}) (interface{}, error) {
	beginDateStr := beginDate.Format("20060102")
	endDateStr := endDate.Format("20060102")

	var respData struct {
		common.WechatError
		List json.RawMessage `json:"list"`
	}

	err := a.getDataForDateRange(url, beginDateStr, endDateStr, &respData)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respData.List, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// getDataForDateRange 获取日期范围数据
func (a *Analysis) getDataForDateRange(url, beginDateStr, endDateStr string, result interface{}) error {
	accessToken, err := a.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", url, accessToken)

	data := &DateRange{
		BeginDate: beginDateStr,
		EndDate:   endDateStr,
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(response, result)
	if err != nil {
		return err
	}

	// 检查错误码
	if wechatErr, ok := result.(common.WechatError); ok && wechatErr.ErrCode != 0 {
		return fmt.Errorf("获取数据失败: %s", wechatErr.ErrMsg)
	}

	return nil
}

// getDataForDate 获取单日数据
func (a *Analysis) getDataForDate(url, dateStr string, result interface{}) error {
	return a.getDataForDateRange(url, dateStr, dateStr, result)
}

func newAnalysis(mp *MiniProgram) *Analysis {
	return &Analysis{mp}
}
