package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// Signature 计算微信签名
func Signature(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	h.Write([]byte(strings.Join(params, "")))
	return hex.EncodeToString(h.Sum(nil))
}

// MD5 计算MD5哈希值
func MD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// PKCS7Padding 实现PKCS7填充
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7UnPadding 实现PKCS7去填充
func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return nil
	}
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

// AESEncrypt AES加密
func AESEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)

	ciphertext := make([]byte, len(plaintext))

	mode := cipher.NewCBCEncrypter(block, key[:blockSize])
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

// AESDecrypt AES解密
func AESDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	if len(ciphertext) < blockSize {
		return nil, fmt.Errorf("密文太短")
	}

	if len(ciphertext)%blockSize != 0 {
		return nil, fmt.Errorf("密文不是块大小的倍数")
	}

	mode := cipher.NewCBCDecrypter(block, key[:blockSize])

	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext = PKCS7UnPadding(plaintext)

	return plaintext, nil
}

// Base64Encode Base64编码
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode Base64解码
func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
