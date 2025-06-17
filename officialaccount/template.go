package officialaccount

import (
	"encoding/json"
	"fmt"
)

// Template 模板消息
type Template struct {
	TemplateID      string `json:"template_id"`      // 模板ID
	Title           string `json:"title"`            // 模板标题
	PrimaryIndustry string `json:"primary_industry"` // 模板所属行业的一级行业
	DeputyIndustry  string `json:"deputy_industry"`  // 模板所属行业的二级行业
	Content         string `json:"content"`          // 模板内容
	Example         string `json:"example"`          // 模板示例
}

// TemplateMessage 发送的模板消息
type TemplateMessage struct {
	ToUser      string                  `json:"touser"`                // 接收者openid
	TemplateID  string                  `json:"template_id"`           // 模板ID
	URL         string                  `json:"url,omitempty"`         // 模板跳转链接
	MiniProgram *MiniProgram            `json:"miniprogram,omitempty"` // 跳小程序所需数据
	Data        map[string]TemplateData `json:"data"`                  // 模板数据
}

// MiniProgram 跳小程序所需数据
type MiniProgram struct {
	AppID    string `json:"appid"`    // 所需跳转到的小程序appid
	PagePath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径
}

// TemplateData 模板消息内容
type TemplateData struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// Industry 行业信息
type Industry struct {
	PrimaryIndustry   string `json:"primary_industry"`   // 帐号设置的主营行业
	SecondaryIndustry string `json:"secondary_industry"` // 帐号设置的副营行业
}

// SetIndustry 设置所属行业
// industryID1: 公众号模板消息所属行业编号
// industryID2: 公众号模板消息所属行业编号
func (oa *OfficialAccount) SetIndustry(industryID1, industryID2 string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=%s", accessToken)
	data := struct {
		IndustryID1 string `json:"industry_id1"`
		IndustryID2 string `json:"industry_id2"`
	}{
		IndustryID1: industryID1,
		IndustryID2: industryID2,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// GetIndustry 获取设置的行业信息
func (oa *OfficialAccount) GetIndustry() (*Industry, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=%s", accessToken)
	response, err := oa.httpGet(url)
	if err != nil {
		return nil, err
	}

	var industry Industry
	err = json.Unmarshal(response, &industry)
	if err != nil {
		return nil, err
	}

	return &industry, nil
}

// AddTemplate 获得模板ID
// shortID: 模板库中模板的编号，有"TM**"和"OPENTMTM**"等形式
func (oa *OfficialAccount) AddTemplate(shortID string) (string, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=%s", accessToken)
	data := struct {
		TemplateIDShort string `json:"template_id_short"`
	}{
		TemplateIDShort: shortID,
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return "", err
	}

	var result struct {
		TemplateID string `json:"template_id"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return "", err
	}

	return result.TemplateID, nil
}

// GetAllPrivateTemplate 获取所有私有模板列表
func (oa *OfficialAccount) GetAllPrivateTemplate() ([]Template, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=%s", accessToken)
	response, err := oa.httpGet(url)
	if err != nil {
		return nil, err
	}

	var result struct {
		TemplateList []Template `json:"template_list"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.TemplateList, nil
}

// DeletePrivateTemplate 删除私有模板
func (oa *OfficialAccount) DeletePrivateTemplate(templateID string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=%s", accessToken)
	data := struct {
		TemplateID string `json:"template_id"`
	}{
		TemplateID: templateID,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// SendTemplate 发送模板消息
func (oa *OfficialAccount) SendTemplate(message *TemplateMessage) (int64, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return 0, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken)
	response, err := oa.httpPost(url, message)
	if err != nil {
		return 0, err
	}

	var result struct {
		MsgID int64 `json:"msgid"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return 0, err
	}

	return result.MsgID, nil
}

// Example of using template message:
/*
message := &TemplateMessage{
    ToUser:     "OPENID",
    TemplateID: "TEMPLATE_ID",
    URL:        "http://example.com",
    Data: map[string]TemplateData{
        "first": {
            Value: "恭喜你购买成功！",
            Color: "#173177",
        },
        "keyword1": {
            Value: "巧克力",
            Color: "#173177",
        },
        "keyword2": {
            Value: "39.8元",
            Color: "#173177",
        },
        "keyword3": {
            Value: "2014年9月22日",
            Color: "#173177",
        },
        "remark": {
            Value: "欢迎再次购买！",
            Color: "#173177",
        },
    },
}

msgID, err := oa.SendTemplate(message)
if err != nil {
    // handle error
}
*/
