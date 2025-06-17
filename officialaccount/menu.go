package officialaccount

import (
	"encoding/json"
	"fmt"

	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// MenuButton 菜单按钮
type MenuButton struct {
	Type       string       `json:"type,omitempty"`
	Name       string       `json:"name,omitempty"`
	Key        string       `json:"key,omitempty"`
	URL        string       `json:"url,omitempty"`
	MediaID    string       `json:"media_id,omitempty"`
	AppID      string       `json:"appid,omitempty"`
	PagePath   string       `json:"pagepath,omitempty"`
	SubButtons []MenuButton `json:"sub_button,omitempty"`
}

// MenuStruct 菜单结构
type MenuStruct struct {
	Buttons []MenuButton `json:"button,omitempty"`
	MenuID  int64        `json:"menuid,omitempty"`
}

// SetMenu 设置菜单
func (menu *Menu) SetMenu(buttons []MenuButton) error {
	accessToken, err := menu.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", common.OfficialAccountMenuCreateURL, accessToken)

	menuStruct := &MenuStruct{
		Buttons: buttons,
	}

	response, err := util.HTTPPost(uri, menuStruct)
	if err != nil {
		return err
	}

	var result common.WechatError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return fmt.Errorf("设置菜单失败: %s", result.ErrMsg)
	}

	return nil
}

// GetMenu 获取菜单配置
func (menu *Menu) GetMenu() (*MenuStruct, error) {
	accessToken, err := menu.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", common.OfficialAccountMenuGetURL, accessToken)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	var result struct {
		common.WechatError
		Menu *MenuStruct `json:"menu"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("获取菜单失败: %s", result.ErrMsg)
	}

	return result.Menu, nil
}

// DeleteMenu 删除菜单
func (menu *Menu) DeleteMenu() error {
	accessToken, err := menu.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", common.OfficialAccountMenuDeleteURL, accessToken)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return err
	}

	var result common.WechatError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return fmt.Errorf("删除菜单失败: %s", result.ErrMsg)
	}

	return nil
}

// NewMenu 创建菜单
func newMenu(oa *OfficialAccount) *Menu {
	return &Menu{oa}
}
