package miniprogram

import (
	"encoding/json"
	"fmt"

	"github.com/zhihao0924/wechat/common"
	"github.com/zhihao0924/wechat/util"
)

// QRCode 小程序码相关API
type QRCode struct {
	*MiniProgram
}

// QRCodeParams 获取小程序码参数
type QRCodeParams struct {
	Scene     string `json:"scene"`                // 最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~
	Page      string `json:"page,omitempty"`       // 必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index，根路径前不要填加 /，不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
	Width     int    `json:"width,omitempty"`      // 二维码的宽度，单位 px，最小 280px，最大 1280px
	AutoColor bool   `json:"auto_color,omitempty"` // 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false
	LineColor struct {
		R string `json:"r"` // 红色，0-255
		G string `json:"g"` // 绿色，0-255
		B string `json:"b"` // 蓝色，0-255
	} `json:"line_color,omitempty"` // auto_color 为 false 时生效，使用 rgb 设置颜色
	IsHyaline bool `json:"is_hyaline,omitempty"` // 是否需要透明底色，为 true 时，生成透明底色的小程序码
}

// GetQRCode 获取小程序码，适用于需要的码数量极多的业务场景
// 通过该接口生成的小程序码，永久有效，数量暂无限制
func (qr *QRCode) GetQRCode(params *QRCodeParams) ([]byte, error) {
	accessToken, err := qr.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", common.MiniProgramCreateQRCodeURL, accessToken)

	response, err := util.HTTPPost(uri, params)
	if err != nil {
		return nil, err
	}

	// 判断是否返回了错误信息
	var result common.WechatError
	err = json.Unmarshal(response, &result)
	if err == nil && result.ErrCode != 0 {
		return nil, fmt.Errorf("获取小程序码失败: %s", result.ErrMsg)
	}

	// 如果没有错误，返回的是图片二进制数据
	return response, nil
}

// GetQRCodeToFile 获取小程序码并保存到文件
func (qr *QRCode) GetQRCodeToFile(params *QRCodeParams, filename string) error {
	data, err := qr.GetQRCode(params)
	if err != nil {
		return err
	}

	return util.WriteFile(filename, data)
}

// CreateQRCode 创建小程序码
// scene: 场景值，用于区分不同的码
// page: 小程序页面路径，可选
// width: 码的宽度，可选
// autoColor: 是否自动配置线条颜色，可选
// lineColor: 线条颜色，可选，autoColor为false时生效
// isHyaline: 是否需要透明底色，可选
func (qr *QRCode) CreateQRCode(scene, page string, options ...interface{}) ([]byte, error) {
	params := &QRCodeParams{
		Scene: scene,
		Page:  page,
		Width: 430, // 默认430
	}

	// 处理可选参数
	for i, option := range options {
		switch i {
		case 0: // width
			if width, ok := option.(int); ok {
				params.Width = width
			}
		case 1: // autoColor
			if autoColor, ok := option.(bool); ok {
				params.AutoColor = autoColor
			}
		case 2: // lineColor
			if color, ok := option.(struct{ R, G, B string }); ok {
				params.LineColor = color
			}
		case 3: // isHyaline
			if isHyaline, ok := option.(bool); ok {
				params.IsHyaline = isHyaline
			}
		}
	}

	return qr.GetQRCode(params)
}

// CreateQRCodeToFile 创建小程序码并保存到文件
func (qr *QRCode) CreateQRCodeToFile(scene, page, filename string, options ...interface{}) error {
	data, err := qr.CreateQRCode(scene, page, options...)
	if err != nil {
		return err
	}

	return util.WriteFile(filename, data)
}

func newQRCode(mp *MiniProgram) *QRCode {
	return &QRCode{mp}
}
