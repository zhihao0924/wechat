package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/zhihao0924/wechat/miniprogram"
)

func main() {
	// 创建小程序配置
	config := &miniprogram.Config{
		AppID:     "your_app_id",     // 替换为你的小程序AppID
		AppSecret: "your_app_secret", // 替换为你的小程序AppSecret
	}

	// 创建小程序实例
	mp := miniprogram.NewMiniProgram(config)

	// 获取加密解密服务
	crypto := mp.Crypto

	// 模拟数据（实际开发中这些数据来自小程序客户端）
	sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
	encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
	iv := "r7BXXKkLb8qrSNn05n0qiA=="

	// 1. 解密用户信息
	fmt.Println("=== 解密用户信息 ===")

	userInfo, err := crypto.DecryptUserInfo(sessionKey, encryptedData, iv)
	if err != nil {
		log.Printf("解密用户信息失败: %v", err)
	} else {
		fmt.Println("用户信息解密成功:")
		fmt.Printf("  - OpenID: %s\n", userInfo.OpenID)
		fmt.Printf("  - 昵称: %s\n", userInfo.NickName)
		fmt.Printf("  - 性别: %s\n", getGender(userInfo.Gender))
		fmt.Printf("  - 城市: %s\n", userInfo.City)
		fmt.Printf("  - 省份: %s\n", userInfo.Province)
		fmt.Printf("  - 国家: %s\n", userInfo.Country)
		fmt.Printf("  - 头像: %s\n", userInfo.AvatarURL)
		fmt.Printf("  - UnionID: %s\n", userInfo.UnionID)
		fmt.Printf("  - 水印AppID: %s\n", userInfo.WaterMark.AppID)
		fmt.Printf("  - 水印时间戳: %d\n", userInfo.WaterMark.Timestamp)
	}

	// 2. 解密手机号信息
	fmt.Println("\n=== 解密手机号信息 ===")

	// 注意：这里使用的是示例数据，实际开发中应该使用从小程序获取的真实加密数据
	phoneEncryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="

	phoneInfo, err := crypto.DecryptPhoneInfo(sessionKey, phoneEncryptedData, iv)
	if err != nil {
		log.Printf("解密手机号信息失败: %v", err)
	} else {
		fmt.Println("手机号信息解密成功:")
		fmt.Printf("  - 手机号: %s\n", phoneInfo.PhoneNumber)
		fmt.Printf("  - 纯手机号: %s\n", phoneInfo.PurePhoneNumber)
		fmt.Printf("  - 国家代码: %s\n", phoneInfo.CountryCode)
		fmt.Printf("  - 水印AppID: %s\n", phoneInfo.WaterMark.AppID)
		fmt.Printf("  - 水印时间戳: %d\n", phoneInfo.WaterMark.Timestamp)
	}

	// 3. 自定义数据加密
	fmt.Println("\n=== 自定义数据加密 ===")

	// 准备要加密的数据
	plainText := "Hello, 这是需要加密的数据!"
	customIV := generateRandomIV() // 在实际应用中应该使用加密安全的随机数生成器

	fmt.Printf("原始数据: %s\n", plainText)
	fmt.Printf("使用的IV: %s\n", customIV)

	encryptedText, err := crypto.EncryptData(sessionKey, plainText, customIV)
	if err != nil {
		log.Printf("加密数据失败: %v", err)
	} else {
		fmt.Printf("加密后的数据: %s\n", encryptedText)

		// 4. 解密刚才加密的数据
		fmt.Println("\n=== 解密自定义数据 ===")

		decryptedData, err := crypto.decrypt(sessionKey, encryptedText, customIV)
		if err != nil {
			log.Printf("解密数据失败: %v", err)
		} else {
			fmt.Printf("解密后的数据: %s\n", string(decryptedData))
		}
	}

	// 5. 加解密最佳实践
	fmt.Println("\n=== 加解密最佳实践 ===")

	// 示例：安全存储敏感数据
	fmt.Println("安全存储敏感数据:")
	fmt.Println("1. 在服务器端进行加解密操作，不在客户端处理密钥")
	fmt.Println("2. 使用安全的随机数生成器生成IV")
	fmt.Println("3. 定期更新sessionKey")
	fmt.Println("4. 验证解密后数据的水印信息")

	// 示例：处理用户信息
	fmt.Println("\n处理用户信息:")
	fmt.Println("1. 获取用户信息时进行必要的授权检查")
	fmt.Println("2. 只解密必要的用户信息")
	fmt.Println("3. 遵守数据保护相关法规")
	fmt.Println("4. 及时清理不再需要的敏感数据")

	// 示例：错误处理
	fmt.Println("\n错误处理:")
	fmt.Println("1. 对解密失败的情况进行适当的错误处理")
	fmt.Println("2. 记录关键的错误信息以便调试")
	fmt.Println("3. 不要在错误信息中暴露敏感信息")
	fmt.Println("4. 为用户提供友好的错误提示")
}

// 获取性别描述
func getGender(gender int) string {
	switch gender {
	case 0:
		return "未知"
	case 1:
		return "男"
	case 2:
		return "女"
	default:
		return fmt.Sprintf("未知(%d)", gender)
	}
}

// 生成随机IV（示例实现，实际应用中应使用加密安全的随机数生成器）
func generateRandomIV() string {
	// 这里使用一个固定的IV仅用于演示
	// 在实际应用中，应该使用crypto/rand包生成随机IV
	return base64.StdEncoding.EncodeToString([]byte("1234567890123456"))
}

/* 使用说明

1. 数据加解密概述：
   - 小程序的敏感数据需要通过加密传输和存储
   - 微信提供了一套完整的加解密机制
   - 加解密使用AES-128-CBC算法
   - 数据包含水印信息用于验证其来源

2. 支持的加解密数据类型：
   - 用户信息（包括昵称、头像等）
   - 手机号信息
   - 自定义数据

3. 加解密流程：
   a. 小程序通过wx.login获取code
   b. 服务器使用code换取session_key
   c. 小程序调用相关接口获取加密数据
   d. 服务器使用session_key解密数据

4. 安全注意事项：
   - session_key 需要安全存储，不能泄露
   - 必须验证数据水印中的appid
   - 密钥和敏感数据应该只在服务器端处理
   - 建议使用HTTPS传输加密数据

5. 常见问题：
   - 41001：session_key 无效
   - 41003：iv 格式错误
   - 41004：encryptedData 格式错误
   - 41005：加密数据解密失败
   - 41006：数据水印验证失败

6. 最佳实践：
   - 在服务器端进行加解密操作
   - 使用安全的随机数生成器
   - 定期更新session_key
   - 验证解密后数据的水印信息
   - 遵守数据保护相关法规

7. 示例代码：
   小程序端获取加密数据：
   ```javascript
   wx.getUserInfo({
     success: function(res) {
       const encryptedData = res.encryptedData;
       const iv = res.iv;
       // 将encryptedData和iv发送到服务器进行解密
     }
   });
   ```

8. 错误处理：
   - 对解密失败的情况进行适当的错误处理
   - 记录关键的错误信息以便调试
   - 不要在错误信息中暴露敏感信息
   - 为用户提供友好的错误提示

9. 数据安全建议：
   - 只解密必要的敏感数据
   - 及时清理不再需要的敏感数据
   - 使用安全的存储方式
   - 实施访问控制和审计机制
*/
