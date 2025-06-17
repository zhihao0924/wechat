package officialaccount

import (
	"encoding/json"
	"fmt"
)

// Card 卡券基础信息
type Card struct {
	CardType string        `json:"card_type"`
	CardID   string        `json:"card_id,omitempty"`
	General  *GeneralCard  `json:"general_card,omitempty"`
	Groupon  *GrouponCard  `json:"groupon,omitempty"`
	Gift     *GiftCard     `json:"gift,omitempty"`
	Cash     *CashCard     `json:"cash,omitempty"`
	Discount *DiscountCard `json:"discount,omitempty"`
	Member   *MemberCard   `json:"member_card,omitempty"`
}

// BaseInfo 卡券基础信息
type BaseInfo struct {
	LogoURL     string `json:"logo_url"`    // 卡券的商户logo
	CodeType    string `json:"code_type"`   // 码型：CODE_TYPE_TEXT 文本；CODE_TYPE_BARCODE 一维码；CODE_TYPE_QRCODE 二维码；CODE_TYPE_ONLY_QRCODE 仅显示二维码；CODE_TYPE_ONLY_BARCODE 仅显示一维码；CODE_TYPE_NONE 不显示任何码型
	BrandName   string `json:"brand_name"`  // 商户名字
	Title       string `json:"title"`       // 卡券名
	Color       string `json:"color"`       // 券颜色
	Notice      string `json:"notice"`      // 卡券使用提醒
	Description string `json:"description"` // 卡券使用说明
	SKU         struct {
		Quantity int `json:"quantity"` // 卡券库存的数量
	} `json:"sku"`
	DateInfo struct {
		Type           string `json:"type"`                       // 使用时间的类型：DATE_TYPE_FIX_TIME_RANGE 表示固定日期区间；DATE_TYPE_FIX_TERM 表示固定时长（自领取后按天算）；DATE_TYPE_PERMANENT 表示永久有效（会员卡类型专用）
		BeginTimestamp int64  `json:"begin_timestamp,omitempty"`  // 起用时间
		EndTimestamp   int64  `json:"end_timestamp,omitempty"`    // 结束时间
		FixedTerm      int    `json:"fixed_term,omitempty"`       // 自领取后多少天内有效
		FixedBeginTerm int    `json:"fixed_begin_term,omitempty"` // 自领取后多少天开始生效
	} `json:"date_info"`
	UseCustomCode        bool     `json:"use_custom_code,omitempty"`         // 是否自定义Code码
	BindOpenID           bool     `json:"bind_openid,omitempty"`             // 是否指定用户领取
	ServicePhone         string   `json:"service_phone,omitempty"`           // 客服电话
	LocationIDList       []string `json:"location_id_list,omitempty"`        // 门店位置ID
	UseAllLocations      bool     `json:"use_all_locations,omitempty"`       // 是否支持全部门店
	CenterTitle          string   `json:"center_title,omitempty"`            // 卡券顶部居中的按钮
	CenterSubTitle       string   `json:"center_sub_title,omitempty"`        // 显示在入口下方的提示语
	CenterURL            string   `json:"center_url,omitempty"`              // 顶部居中的url
	CustomURLName        string   `json:"custom_url_name,omitempty"`         // 自定义跳转外链的入口名字
	CustomURL            string   `json:"custom_url,omitempty"`              // 自定义跳转的URL
	CustomURLSubTitle    string   `json:"custom_url_sub_title,omitempty"`    // 显示在入口右侧的提示语
	PromotionURLName     string   `json:"promotion_url_name,omitempty"`      // 营销场景的自定义入口名称
	PromotionURL         string   `json:"promotion_url,omitempty"`           // 入口跳转外链的地址链接
	PromotionURLSubTitle string   `json:"promotion_url_sub_title,omitempty"` // 显示在营销入口右侧的提示语
	GetLimit             int      `json:"get_limit,omitempty"`               // 每人可领券的数量限制
	CanShare             bool     `json:"can_share,omitempty"`               // 卡券领取页面是否可分享
	CanGiveFriend        bool     `json:"can_give_friend,omitempty"`         // 卡券是否可转赠
}

// GeneralCard 通用优惠券信息
type GeneralCard struct {
	BaseInfo      BaseInfo `json:"base_info"`
	DefaultDetail string   `json:"default_detail"` // 优惠券详情
}

// GrouponCard 团购券信息
type GrouponCard struct {
	BaseInfo   BaseInfo `json:"base_info"`
	DealDetail string   `json:"deal_detail"` // 团购券专用，团购详情
}

// GiftCard 礼品券信息
type GiftCard struct {
	BaseInfo BaseInfo `json:"base_info"`
	Gift     string   `json:"gift"` // 礼品券专用，表示礼品名字
}

// CashCard 代金券信息
type CashCard struct {
	BaseInfo   BaseInfo `json:"base_info"`
	LeastCost  int      `json:"least_cost,omitempty"` // 代金券专用，表示起用金额（单位为分）
	ReduceCost int      `json:"reduce_cost"`          // 代金券专用，表示减免金额（单位为分）
}

// DiscountCard 折扣券信息
type DiscountCard struct {
	BaseInfo BaseInfo `json:"base_info"`
	Discount int      `json:"discount"` // 折扣券专用，表示打折额度（百分比）
}

// MemberCard 会员卡信息
type MemberCard struct {
	BaseInfo         BaseInfo     `json:"base_info"`
	BackgroundPicURL string       `json:"background_pic_url,omitempty"` // 会员卡背景图
	PrerogativeDesc  string       `json:"prerogative_desc"`             // 会员卡特权说明
	AutoActivate     bool         `json:"auto_activate,omitempty"`      // 是否自动激活
	WXActivate       bool         `json:"wx_activate,omitempty"`        // 是否一键开卡
	SupplyBonus      bool         `json:"supply_bonus,omitempty"`       // 显示积分
	BonusURL         string       `json:"bonus_url,omitempty"`          // 设置跳转外链查看积分详情
	SupplyBalance    bool         `json:"supply_balance,omitempty"`     // 是否支持储值
	BalanceURL       string       `json:"balance_url,omitempty"`        // 设置跳转外链查看余额详情
	CustomField1     *CustomField `json:"custom_field1,omitempty"`      // 自定义会员信息类目1
	CustomField2     *CustomField `json:"custom_field2,omitempty"`      // 自定义会员信息类目2
	CustomField3     *CustomField `json:"custom_field3,omitempty"`      // 自定义会员信息类目3
	ActivateURL      string       `json:"activate_url,omitempty"`       // 激活会员卡的url
	CustomCellURL    string       `json:"custom_cell_url,omitempty"`    // 自定义会员信息类目跳转外链的地址链接
}

// CustomField 自定义会员信息类目
type CustomField struct {
	NameType string `json:"name_type"`      // 会员信息类目名称
	URL      string `json:"url,omitempty"`  // 自定义跳转外链的地址链接
	Tips     string `json:"tips,omitempty"` // 自定义提示语
}

// CardResponse 创建卡券返回结果
type CardResponse struct {
	CardID string `json:"card_id"` // 卡券ID
}

// CreateCard 创建卡券
func (oa *OfficialAccount) CreateCard(card *Card) (*CardResponse, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/create?access_token=%s", accessToken)
	data := struct {
		Card *Card `json:"card"`
	}{
		Card: card,
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result CardResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetCard 查询卡券详情
func (oa *OfficialAccount) GetCard(cardID string) (*Card, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/get?access_token=%s", accessToken)
	data := struct {
		CardID string `json:"card_id"`
	}{
		CardID: cardID,
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result struct {
		Card *Card `json:"card"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.Card, nil
}

// BatchGetCards 批量查询卡券列表
func (oa *OfficialAccount) BatchGetCards(offset, count int, statusList []string) ([]Card, int, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, 0, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/batchget?access_token=%s", accessToken)
	data := struct {
		Offset     int      `json:"offset"`
		Count      int      `json:"count"`
		StatusList []string `json:"status_list,omitempty"`
	}{
		Offset:     offset,
		Count:      count,
		StatusList: statusList,
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, 0, err
	}

	var result struct {
		CardIDList []string `json:"card_id_list"`
		TotalNum   int      `json:"total_num"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, 0, err
	}

	// 获取每张卡券的详细信息
	cards := make([]Card, 0, len(result.CardIDList))
	for _, cardID := range result.CardIDList {
		card, err := oa.GetCard(cardID)
		if err != nil {
			return nil, 0, err
		}
		cards = append(cards, *card)
	}

	return cards, result.TotalNum, nil
}

// UpdateCard 更新卡券
func (oa *OfficialAccount) UpdateCard(card *Card) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/update?access_token=%s", accessToken)
	data := struct {
		Card *Card `json:"card"`
	}{
		Card: card,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// DeleteCard 删除卡券
func (oa *OfficialAccount) DeleteCard(cardID string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/delete?access_token=%s", accessToken)
	data := struct {
		CardID string `json:"card_id"`
	}{
		CardID: cardID,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// CardQRCode 卡券二维码
type CardQRCode struct {
	Ticket        string `json:"ticket"`         // 获取的二维码ticket，凭借此ticket调用通过ticket换取二维码接口可以在有效时间内换取二维码
	ExpireSeconds int    `json:"expire_seconds"` // 二维码的有效时间
	URL           string `json:"url"`            // 二维码图片解析后的地址
}

// CreateCardQRCode 创建卡券二维码
func (oa *OfficialAccount) CreateCardQRCode(cardID string, outerStr string, expireSeconds int) (*CardQRCode, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/qrcode/create?access_token=%s", accessToken)
	data := struct {
		ActionName    string `json:"action_name"`
		ExpireSeconds int    `json:"expire_seconds"`
		ActionInfo    struct {
			Card struct {
				CardID   string `json:"card_id"`
				OuterStr string `json:"outer_str,omitempty"`
			} `json:"card"`
		} `json:"action_info"`
	}{
		ActionName:    "QR_CARD",
		ExpireSeconds: expireSeconds,
		ActionInfo: struct {
			Card struct {
				CardID   string `json:"card_id"`
				OuterStr string `json:"outer_str,omitempty"`
			} `json:"card"`
		}{
			Card: struct {
				CardID   string `json:"card_id"`
				OuterStr string `json:"outer_str,omitempty"`
			}{
				CardID:   cardID,
				OuterStr: outerStr,
			},
		},
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result CardQRCode
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CardCode 卡券Code码
type CardCode struct {
	Code   string `json:"code"`    // 卡券Code码
	CardID string `json:"card_id"` // 卡券ID
}

// ConsumeCardCode 核销卡券Code码
func (oa *OfficialAccount) ConsumeCardCode(code string) (*CardCode, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/code/consume?access_token=%s", accessToken)
	data := struct {
		Code string `json:"code"`
	}{
		Code: code,
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result CardCode
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DecryptCardCode 解密卡券Code码
func (oa *OfficialAccount) DecryptCardCode(encryptCode string) (*CardCode, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/code/decrypt?access_token=%s", accessToken)
	data := struct {
		EncryptCode string `json:"encrypt_code"`
	}{
		EncryptCode: encryptCode,
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result CardCode
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetCardCode 获取卡券Code码信息
func (oa *OfficialAccount) GetCardCode(code string, checkConsume bool) (*CardCode, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/code/get?access_token=%s", accessToken)
	data := struct {
		Code         string `json:"code"`
		CheckConsume bool   `json:"check_consume,omitempty"`
	}{
		Code:         code,
		CheckConsume: checkConsume,
	}

	response, err := oa.httpPost(url, data)
	if err != nil {
		return nil, err
	}

	var result CardCode
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateCardCode 更新卡券Code码
func (oa *OfficialAccount) UpdateCardCode(code, newCode string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/code/update?access_token=%s", accessToken)
	data := struct {
		Code    string `json:"code"`
		NewCode string `json:"new_code"`
	}{
		Code:    code,
		NewCode: newCode,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// DisableCardCode 设置卡券Code失效
func (oa *OfficialAccount) DisableCardCode(code, reason string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/code/unavailable?access_token=%s", accessToken)
	data := struct {
		Code   string `json:"code"`
		Reason string `json:"reason,omitempty"`
	}{
		Code:   code,
		Reason: reason,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// ActivateMemberCard 激活会员卡
func (oa *OfficialAccount) ActivateMemberCard(params map[string]interface{}) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/membercard/activate?access_token=%s", accessToken)
	_, err = oa.httpPost(url, params)
	return err
}

// UpdateMemberCard 更新会员卡信息
func (oa *OfficialAccount) UpdateMemberCard(params map[string]interface{}) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/membercard/updateuser?access_token=%s", accessToken)
	_, err = oa.httpPost(url, params)
	return err
}

// ModifyStock 修改库存
func (oa *OfficialAccount) ModifyStock(cardID string, increase, reduce int) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/modifystock?access_token=%s", accessToken)
	data := struct {
		CardID   string `json:"card_id"`
		Increase int    `json:"increase_stock_value,omitempty"`
		Reduce   int    `json:"reduce_stock_value,omitempty"`
	}{
		CardID:   cardID,
		Increase: increase,
		Reduce:   reduce,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// GetCardStatistics 获取卡券统计数据
func (oa *OfficialAccount) GetCardStatistics(beginDate, endDate string, condSource int) ([]byte, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/datacube/getcardbizuininfo?access_token=%s", accessToken)
	data := struct {
		BeginDate  string `json:"begin_date"`
		EndDate    string `json:"end_date"`
		CondSource int    `json:"cond_source"`
	}{
		BeginDate:  beginDate,
		EndDate:    endDate,
		CondSource: condSource,
	}

	return oa.httpPost(url, data)
}

// AddCardWhiteList 添加卡券测试白名单
func (oa *OfficialAccount) AddCardWhiteList(openids []string, usernames []string) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/testwhitelist/set?access_token=%s", accessToken)
	data := struct {
		Openid   []string `json:"openid,omitempty"`
		Username []string `json:"username,omitempty"`
	}{
		Openid:   openids,
		Username: usernames,
	}

	_, err = oa.httpPost(url, data)
	return err
}

// GetInvoiceAuthURL 获取授权页链接
func (oa *OfficialAccount) GetInvoiceAuthURL(params map[string]interface{}) (string, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/invoice/getauthurl?access_token=%s", accessToken)
	resp, err := oa.httpPost(url, params)
	if err != nil {
		return "", err
	}

	var result struct {
		URL string `json:"url"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return "", err
	}

	return result.URL, nil
}

// SetCardEntrance 设置卡券自定义入口
func (oa *OfficialAccount) SetCardEntrance(params map[string]interface{}) error {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/landingpage/create?access_token=%s", accessToken)
	_, err = oa.httpPost(url, params)
	return err
}

// GetCardJSConfig 获取卡券JSAPI配置
func (oa *OfficialAccount) GetCardJSConfig(params map[string]interface{}) ([]byte, error) {
	accessToken, err := oa.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/card/mpnews/gethtml?access_token=%s", accessToken)
	return oa.httpPost(url, params)
}
