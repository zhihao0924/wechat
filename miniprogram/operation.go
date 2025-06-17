package miniprogram

import (
	"encoding/json"
	"fmt"

	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// Operation 运维中心相关API
type Operation struct {
	*MiniProgram
}

// 运维中心相关API地址
const (
	// 获取域名配置
	getDomainInfoURL = "https://api.weixin.qq.com/wxa/getwxadevinfo"
	// 修改域名配置
	modifyDomainURL = "https://api.weixin.qq.com/wxa/modify_domain"
	// 设置业务域名
	setWebviewDomainURL = "https://api.weixin.qq.com/wxa/setwebviewdomain"
	// 获取性能数据
	getPerformanceDataURL = "https://api.weixin.qq.com/wxa/business/performance/boot"
	// 获取服务状态
	getServerStatusURL = "https://api.weixin.qq.com/wxa/getwxadevinfo"
)

// DomainInfo 域名配置信息
type DomainInfo struct {
	RequestDomain   []string `json:"requestdomain"`   // request合法域名
	WSRequestDomain []string `json:"wsrequestdomain"` // socket合法域名
	UploadDomain    []string `json:"uploaddomain"`    // uploadFile合法域名
	DownloadDomain  []string `json:"downloaddomain"`  // downloadFile合法域名
	BizDomain       []string `json:"bizdomain"`       // 业务域名
}

// PerformanceData 性能数据
type PerformanceData struct {
	TimeRange string  `json:"time_range"` // 时间范围
	PageCount int     `json:"page_count"` // 页面访问次数
	AvgTime   float64 `json:"avg_time"`   // 平均耗时（毫秒）
	ErrorRate float64 `json:"error_rate"` // 错误率
}

// ServerStatus 服务状态信息
type ServerStatus struct {
	Status      int    `json:"status"`      // 服务状态
	Description string `json:"description"` // 状态描述
	ErrorCount  int    `json:"error_count"` // 错误次数
	LastError   string `json:"last_error"`  // 最后一次错误信息
}

// GetDomainInfo 获取域名配置
func (o *Operation) GetDomainInfo() (*DomainInfo, error) {
	accessToken, err := o.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", getDomainInfoURL, accessToken)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	var result struct {
		common.WechatError
		DomainInfo
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("获取域名配置失败: %s", result.ErrMsg)
	}

	return &result.DomainInfo, nil
}

// ModifyDomain 修改域名配置
func (o *Operation) ModifyDomain(info *DomainInfo, action string) error {
	accessToken, err := o.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", modifyDomainURL, accessToken)

	data := struct {
		Action string `json:"action"`
		DomainInfo
	}{
		Action:     action, // add 添加, delete 删除, set 覆盖, get 获取
		DomainInfo: *info,
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
		return fmt.Errorf("修改域名配置失败: %s", result.ErrMsg)
	}

	return nil
}

// SetWebviewDomain 设置业务域名
func (o *Operation) SetWebviewDomain(domains []string, action string) error {
	accessToken, err := o.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", setWebviewDomainURL, accessToken)

	data := struct {
		Action  string   `json:"action"`
		Domains []string `json:"webviewdomain"`
	}{
		Action:  action, // add 添加, delete 删除, set 覆盖, get 获取
		Domains: domains,
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
		return fmt.Errorf("设置业务域名失败: %s", result.ErrMsg)
	}

	return nil
}

// GetPerformanceData 获取性能数据
func (o *Operation) GetPerformanceData(startTime, endTime string) ([]PerformanceData, error) {
	accessToken, err := o.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", getPerformanceDataURL, accessToken)

	data := struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}{
		StartTime: startTime,
		EndTime:   endTime,
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return nil, err
	}

	var result struct {
		common.WechatError
		DataList []PerformanceData `json:"data_list"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("获取性能数据失败: %s", result.ErrMsg)
	}

	return result.DataList, nil
}

// GetServerStatus 获取服务状态
func (o *Operation) GetServerStatus() (*ServerStatus, error) {
	accessToken, err := o.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", getServerStatusURL, accessToken)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	var result struct {
		common.WechatError
		ServerStatus
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("获取服务状态失败: %s", result.ErrMsg)
	}

	return &result.ServerStatus, nil
}

func newOperation(mp *MiniProgram) *Operation {
	return &Operation{mp}
}
