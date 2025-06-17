package miniprogram

import (
	"encoding/json"
	"fmt"

	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// Live 直播相关API
type Live struct {
	*MiniProgram
}

// 直播相关API地址
const (
	// 获取直播房间列表
	getLiveRoomListURL = "https://api.weixin.qq.com/wxa/business/getliveinfo"
	// 获取直播间回放
	getLiveReplayURL = "https://api.weixin.qq.com/wxa/business/getliveinfo"
	// 获取直播间商品列表
	getLiveGoodsURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/getapproved"
	// 商品添加并提审
	addGoodsURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/add"
	// 撤回商品审核
	resetAuditGoodsURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/resetaudit"
	// 重新提交商品审核
	resubmitAuditGoodsURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/audit"
	// 删除商品
	deleteGoodsURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/delete"
	// 更新商品
	updateGoodsURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/update"
	// 获取商品状态
	getGoodsStatusURL = "https://api.weixin.qq.com/wxa/business/getgoodswarehouse"
)

// LiveRoom 直播间信息
type LiveRoom struct {
	Name       string `json:"name"`        // 直播间名称
	RoomID     int    `json:"roomid"`      // 直播间ID
	CoverImg   string `json:"cover_img"`   // 直播间背景图
	ShareImg   string `json:"share_img"`   // 直播间分享图
	StartTime  int64  `json:"start_time"`  // 直播开始时间
	EndTime    int64  `json:"end_time"`    // 直播结束时间
	AnchorName string `json:"anchor_name"` // 主播昵称
	Status     int    `json:"live_status"` // 直播间状态
}

// LiveGoods 直播商品信息
type LiveGoods struct {
	GoodsID      int     `json:"goods_id"`        // 商品ID
	CoverImg     string  `json:"cover_img"`       // 商品图片
	Name         string  `json:"name"`            // 商品名称
	Price        float64 `json:"price"`           // 商品价格
	URL          string  `json:"url"`             // 商品详情页链接
	Status       int     `json:"audit_status"`    // 审核状态
	ThirdPartyID string  `json:"third_party_tag"` // 第三方商品ID
}

// GetRoomList 获取直播间列表
func (l *Live) GetRoomList(start, limit int) ([]LiveRoom, int, error) {
	accessToken, err := l.GetAccessToken()
	if err != nil {
		return nil, 0, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", getLiveRoomListURL, accessToken)

	data := struct {
		Start int `json:"start"`
		Limit int `json:"limit"`
	}{
		Start: start,
		Limit: limit,
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return nil, 0, err
	}

	var result struct {
		common.WechatError
		Total    int        `json:"total"`
		RoomInfo []LiveRoom `json:"room_info"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, 0, err
	}

	if result.ErrCode != 0 {
		return nil, 0, fmt.Errorf("获取直播间列表失败: %s", result.ErrMsg)
	}

	return result.RoomInfo, result.Total, nil
}

// GetReplay 获取直播间回放
func (l *Live) GetReplay(roomID int, start, limit int) ([]string, int, error) {
	accessToken, err := l.GetAccessToken()
	if err != nil {
		return nil, 0, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", getLiveReplayURL, accessToken)

	data := struct {
		Action string `json:"action"`
		RoomID int    `json:"room_id"`
		Start  int    `json:"start"`
		Limit  int    `json:"limit"`
	}{
		Action: "get_replay",
		RoomID: roomID,
		Start:  start,
		Limit:  limit,
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return nil, 0, err
	}

	var result struct {
		common.WechatError
		Total  int      `json:"total"`
		Videos []string `json:"live_replay"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, 0, err
	}

	if result.ErrCode != 0 {
		return nil, 0, fmt.Errorf("获取直播间回放失败: %s", result.ErrMsg)
	}

	return result.Videos, result.Total, nil
}

// GetGoods 获取直播间商品列表
func (l *Live) GetGoods(status int, start, limit int) ([]LiveGoods, int, error) {
	accessToken, err := l.GetAccessToken()
	if err != nil {
		return nil, 0, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", getLiveGoodsURL, accessToken)

	data := struct {
		Status int `json:"status"`
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	}{
		Status: status,
		Offset: start,
		Limit:  limit,
	}

	response, err := util.HTTPPost(uri, data)
	if err != nil {
		return nil, 0, err
	}

	var result struct {
		common.WechatError
		Total int         `json:"total"`
		Goods []LiveGoods `json:"goods"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, 0, err
	}

	if result.ErrCode != 0 {
		return nil, 0, fmt.Errorf("获取直播间商品列表失败: %s", result.ErrMsg)
	}

	return result.Goods, result.Total, nil
}

// AddGoods 添加商品并提交审核
func (l *Live) AddGoods(goods *LiveGoods) (int, error) {
	accessToken, err := l.GetAccessToken()
	if err != nil {
		return 0, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", addGoodsURL, accessToken)

	response, err := util.HTTPPost(uri, goods)
	if err != nil {
		return 0, err
	}

	var result struct {
		common.WechatError
		GoodsID int `json:"goodsId"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return 0, err
	}

	if result.ErrCode != 0 {
		return 0, fmt.Errorf("添加商品失败: %s", result.ErrMsg)
	}

	return result.GoodsID, nil
}

// ResetAuditGoods 撤回商品审核
func (l *Live) ResetAuditGoods(goodsID int) error {
	accessToken, err := l.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", resetAuditGoodsURL, accessToken)

	data := struct {
		GoodsID int `json:"goodsId"`
	}{
		GoodsID: goodsID,
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
		return fmt.Errorf("撤回商品审核失败: %s", result.ErrMsg)
	}

	return nil
}

// ResubmitAuditGoods 重新提交商品审核
func (l *Live) ResubmitAuditGoods(goodsID int) error {
	accessToken, err := l.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", resubmitAuditGoodsURL, accessToken)

	data := struct {
		GoodsID int `json:"goodsId"`
	}{
		GoodsID: goodsID,
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
		return fmt.Errorf("重新提交商品审核失败: %s", result.ErrMsg)
	}

	return nil
}

// DeleteGoods 删除商品
func (l *Live) DeleteGoods(goodsID int) error {
	accessToken, err := l.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", deleteGoodsURL, accessToken)

	data := struct {
		GoodsID int `json:"goodsId"`
	}{
		GoodsID: goodsID,
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
		return fmt.Errorf("删除商品失败: %s", result.ErrMsg)
	}

	return nil
}

// UpdateGoods 更新商品
func (l *Live) UpdateGoods(goods *LiveGoods) error {
	accessToken, err := l.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", updateGoodsURL, accessToken)

	response, err := util.HTTPPost(uri, goods)
	if err != nil {
		return err
	}

	var result common.WechatError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return fmt.Errorf("更新商品失败: %s", result.ErrMsg)
	}

	return nil
}

func newLive(mp *MiniProgram) *Live {
	return &Live{mp}
}
