package common

// 微信API地址
const (
	// WechatAPIURL 微信API基础URL
	WechatAPIURL = "https://api.weixin.qq.com"

	// OfficialAccount 公众号相关
	OfficialAccountUserInfoURL     = WechatAPIURL + "/cgi-bin/user/info"
	OfficialAccountMenuCreateURL   = WechatAPIURL + "/cgi-bin/menu/create"
	OfficialAccountMenuGetURL      = WechatAPIURL + "/cgi-bin/menu/get"
	OfficialAccountMenuDeleteURL   = WechatAPIURL + "/cgi-bin/menu/delete"
	OfficialAccountQRCodeCreateURL = WechatAPIURL + "/cgi-bin/qrcode/create"
	OfficialAccountQRCodeShowURL   = "https://mp.weixin.qq.com/cgi-bin/showqrcode"

	// MiniProgram 小程序相关
	MiniProgramCode2SessionURL   = WechatAPIURL + "/sns/jscode2session"
	MiniProgramGetAccessTokenURL = WechatAPIURL + "/cgi-bin/token"
	MiniProgramSendTemplateURL   = WechatAPIURL + "/cgi-bin/message/wxopen/template/send"
	MiniProgramCreateQRCodeURL   = WechatAPIURL + "/wxa/getwxacodeunlimit"
	MiniProgramGetPaidUnionIDURL = WechatAPIURL + "/wxa/getpaidunionid"

	// 支付相关
	WechatPayUnifiedOrderURL = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	WechatPayOrderQueryURL   = "https://api.mch.weixin.qq.com/pay/orderquery"
	WechatPayCloseOrderURL   = "https://api.mch.weixin.qq.com/pay/closeorder"
	WechatPayRefundURL       = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	WechatPayRefundQueryURL  = "https://api.mch.weixin.qq.com/pay/refundquery"
	WechatPayDownloadBillURL = "https://api.mch.weixin.qq.com/pay/downloadbill"
)

// 微信API返回码
const (
	// WechatSuccess 成功
	WechatSuccess = 0

	// WechatSystemBusy 系统繁忙
	WechatSystemBusy = -1

	// WechatInvalidToken 不合法的凭证
	WechatInvalidToken = 40001

	// WechatInvalidGrantType 不合法的grant_type
	WechatInvalidGrantType = 40002

	// WechatInvalidOpenID 不合法的OpenID
	WechatInvalidOpenID = 40003

	// WechatInvalidMediaType 不合法的媒体文件类型
	WechatInvalidMediaType = 40004

	// WechatInvalidMediaID 不合法的媒体文件ID
	WechatInvalidMediaID = 40007

	// WechatInvalidMessage 不合法的消息类型
	WechatInvalidMessage = 40008

	// WechatInvalidAppID 不合法的AppID
	WechatInvalidAppID = 40013

	// WechatAccessTokenExpired 访问令牌过期
	WechatAccessTokenExpired = 42001

	// WechatUserUnsubscribe 用户未关注公众号
	WechatUserUnsubscribe = 43004

	// WechatAPILimit API调用次数达到上限
	WechatAPILimit = 45009

	// WechatUserBlock 用户被拉黑
	WechatUserBlock = 48001

	// WechatAPINotAuthorized API未授权
	WechatAPINotAuthorized = 48004

	// WechatUserNotAuthorize 用户未授权
	WechatUserNotAuthorize = 50001
)
