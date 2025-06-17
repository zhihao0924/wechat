package miniprogram

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/zhihao0924/wechat/util"
)

// Crypto 加密解密相关API
type Crypto struct {
	*MiniProgram
}

// WaterMark 水印
type WaterMark struct {
	AppID     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

// UserInfo 用户信息
type UserInfo struct {
	OpenID    string    `json:"openId"`
	NickName  string    `json:"nickName"`
	Gender    int       `json:"gender"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	AvatarURL string    `json:"avatarUrl"`
	UnionID   string    `json:"unionId"`
	WaterMark WaterMark `json:"watermark"`
}

// PhoneInfo 手机号信息
type PhoneInfo struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	WaterMark       WaterMark `json:"watermark"`
}

// DecryptUserInfo 解密用户信息
func (c *Crypto) DecryptUserInfo(sessionKey, encryptedData, iv string) (*UserInfo, error) {
	data, err := c.decrypt(sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}

	userInfo := new(UserInfo)
	err = json.Unmarshal(data, userInfo)
	if err != nil {
		return nil, fmt.Errorf("解析用户数据失败: %v", err)
	}

	// 校验水印
	if userInfo.WaterMark.AppID != c.Config.AppID {
		return nil, errors.New("水印AppID不匹配")
	}

	return userInfo, nil
}

// DecryptPhoneInfo 解密手机号信息
func (c *Crypto) DecryptPhoneInfo(sessionKey, encryptedData, iv string) (*PhoneInfo, error) {
	data, err := c.decrypt(sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}

	phoneInfo := new(PhoneInfo)
	err = json.Unmarshal(data, phoneInfo)
	if err != nil {
		return nil, fmt.Errorf("解析手机号数据失败: %v", err)
	}

	// 校验水印
	if phoneInfo.WaterMark.AppID != c.Config.AppID {
		return nil, errors.New("水印AppID不匹配")
	}

	return phoneInfo, nil
}

// decrypt 解密数据
func (c *Crypto) decrypt(sessionKey, encryptedData, iv string) ([]byte, error) {
	// Base64解码
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, fmt.Errorf("解码sessionKey失败: %v", err)
	}

	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, fmt.Errorf("解码encryptedData失败: %v", err)
	}

	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, fmt.Errorf("解码iv失败: %v", err)
	}

	// AES-128-CBC解密
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, fmt.Errorf("创建AES密码块失败: %v", err)
	}

	if len(ivBytes) != block.BlockSize() {
		return nil, errors.New("iv长度不正确")
	}

	// 解密
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	plainText := make([]byte, len(cipherText))
	mode.CryptBlocks(plainText, cipherText)

	// 去除PKCS#7填充
	plainText, err = util.PKCS7UnPadding(plainText)
	if err != nil {
		return nil, fmt.Errorf("去除PKCS#7填充失败: %v", err)
	}

	return plainText, nil
}

// EncryptData 加密数据
func (c *Crypto) EncryptData(sessionKey, plainText, iv string) (string, error) {
	// Base64解码
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return "", fmt.Errorf("解码sessionKey失败: %v", err)
	}

	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", fmt.Errorf("解码iv失败: %v", err)
	}

	// AES-128-CBC加密
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", fmt.Errorf("创建AES密码块失败: %v", err)
	}

	if len(ivBytes) != block.BlockSize() {
		return "", errors.New("iv长度不正确")
	}

	// PKCS#7填充
	plainTextBytes := []byte(plainText)
	plainTextBytes = util.PKCS7Padding(plainTextBytes, block.BlockSize())

	// 加密
	cipherText := make([]byte, len(plainTextBytes))
	mode := cipher.NewCBCEncrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, plainTextBytes)

	// Base64编码
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func newCrypto(mp *MiniProgram) *Crypto {
	return &Crypto{mp}
}
