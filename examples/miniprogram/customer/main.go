package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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

	// 获取客服服务
	customer := mp.Customer

	// 示例用户OpenID
	userOpenID := "oXYZ123_example_openid" // 替换为实际用户的OpenID

	// 1. 发送文本消息
	fmt.Println("=== 发送文本消息 ===")

	textContent := "您好！这是一条客服文本消息。如有问题，请随时咨询我们的客服团队。"

	fmt.Printf("发送文本消息给用户 %s:\n", userOpenID)
	fmt.Printf("  - 内容: %s\n", textContent)

	// 注释掉实际的API调用，避免在示例中真正发送消息
	/*
		err := customer.SendText(userOpenID, textContent)
		if err != nil {
			log.Printf("发送文本消息失败: %v", err)
		} else {
			fmt.Println("发送文本消息成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 2. 发送图片消息
	fmt.Println("\n=== 发送图片消息 ===")

	// 图片媒体ID（需要先上传临时素材获取）
	mediaID := "MEDIA_ID_EXAMPLE" // 替换为实际的媒体ID

	fmt.Printf("发送图片消息给用户 %s:\n", userOpenID)
	fmt.Printf("  - 媒体ID: %s\n", mediaID)

	// 注释掉实际的API调用，避免在示例中真正发送消息
	/*
		err = customer.SendImage(userOpenID, mediaID)
		if err != nil {
			log.Printf("发送图片消息失败: %v", err)
		} else {
			fmt.Println("发送图片消息成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 3. 发送图文链接消息
	fmt.Println("\n=== 发送图文链接消息 ===")

	linkTitle := "新品上市通知"
	linkDescription := "点击查看我们的最新产品信息和优惠活动"
	linkURL := "https://example.com/products/new"
	linkThumbURL := "https://example.com/images/product-thumb.jpg"

	fmt.Printf("发送图文链接消息给用户 %s:\n", userOpenID)
	fmt.Printf("  - 标题: %s\n", linkTitle)
	fmt.Printf("  - 描述: %s\n", linkDescription)
	fmt.Printf("  - 链接: %s\n", linkURL)
	fmt.Printf("  - 缩略图: %s\n", linkThumbURL)

	// 注释掉实际的API调用，避免在示例中真正发送消息
	/*
		err = customer.SendLink(userOpenID, linkTitle, linkDescription, linkURL, linkThumbURL)
		if err != nil {
			log.Printf("发送图文链接消息失败: %v", err)
		} else {
			fmt.Println("发送图文链接消息成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 4. 发送小程序卡片消息
	fmt.Println("\n=== 发送小程序卡片消息 ===")

	cardTitle := "会员中心"
	cardPagePath := "pages/member/index"
	cardThumbMediaID := "THUMB_MEDIA_ID_EXAMPLE" // 替换为实际的缩略图媒体ID

	fmt.Printf("发送小程序卡片消息给用户 %s:\n", userOpenID)
	fmt.Printf("  - 标题: %s\n", cardTitle)
	fmt.Printf("  - 页面路径: %s\n", cardPagePath)
	fmt.Printf("  - 缩略图媒体ID: %s\n", cardThumbMediaID)

	// 注释掉实际的API调用，避免在示例中真正发送消息
	/*
		err = customer.SendMiniProgramPage(userOpenID, cardTitle, cardPagePath, cardThumbMediaID)
		if err != nil {
			log.Printf("发送小程序卡片消息失败: %v", err)
		} else {
			fmt.Println("发送小程序卡片消息成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 5. 设置客服输入状态
	fmt.Println("\n=== 设置客服输入状态 ===")

	fmt.Printf("设置对用户 %s 的输入状态为「正在输入」\n", userOpenID)

	// 注释掉实际的API调用，避免在示例中真正设置状态
	/*
		err = customer.SetTyping(userOpenID, true)
		if err != nil {
			log.Printf("设置输入状态失败: %v", err)
		} else {
			fmt.Println("设置输入状态成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 模拟客服正在处理
	fmt.Println("客服正在处理中...")

	// 取消输入状态
	fmt.Printf("取消对用户 %s 的输入状态\n", userOpenID)

	// 注释掉实际的API调用，避免在示例中真正设置状态
	/*
		err = customer.SetTyping(userOpenID, false)
		if err != nil {
			log.Printf("取消输入状态失败: %v", err)
		} else {
			fmt.Println("取消输入状态成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 6. 获取临时素材
	fmt.Println("\n=== 获取临时素材 ===")

	// 临时素材的媒体ID
	tempMediaID := "TEMP_MEDIA_ID_EXAMPLE" // 替换为实际的临时素材媒体ID

	fmt.Printf("获取临时素材 (媒体ID: %s)\n", tempMediaID)

	// 注释掉实际的API调用，避免在示例中真正获取素材
	/*
		mediaData, err := customer.GetTempMedia(tempMediaID)
		if err != nil {
			log.Printf("获取临时素材失败: %v", err)
		} else {
			// 保存临时素材到文件
			tempDir := "temp_media"
			if err := os.MkdirAll(tempDir, 0755); err != nil {
				log.Printf("创建临时目录失败: %v", err)
			} else {
				filePath := filepath.Join(tempDir, tempMediaID+".jpg") // 假设是图片
				if err := os.WriteFile(filePath, mediaData, 0644); err != nil {
					log.Printf("保存临时素材失败: %v", err)
				} else {
					fmt.Printf("临时素材已保存到: %s\n", filePath)
				}
			}
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 7. 客服消息最佳实践
	fmt.Println("\n=== 客服消息最佳实践 ===")

	// 示例：客服消息流程
	fmt.Println("客服消息流程:")
	fmt.Println("1. 用户在小程序中发起客服会话")
	fmt.Println("2. 接收用户消息并进行处理")
	fmt.Println("3. 根据用户需求发送相应类型的客服消息")
	fmt.Println("4. 使用输入状态提示用户客服正在处理")
	fmt.Println("5. 结束会话或转人工客服")

	// 示例：自动回复系统
	fmt.Println("\n自动回复系统示例:")
	fmt.Println("func handleUserMessage(userID, content string) {")
	fmt.Println("    // 1. 分析用户消息内容")
	fmt.Println("    if strings.Contains(content, \"价格\") {")
	fmt.Println("        // 发送价格相关信息")
	fmt.Println("        customer.SendText(userID, \"我们的产品价格为...\") ")
	fmt.Println("    } else if strings.Contains(content, \"配送\") {")
	fmt.Println("        // 发送配送相关信息")
	fmt.Println("        customer.SendText(userID, \"我们的配送范围包括...\") ")
	fmt.Println("    } else if strings.Contains(content, \"退款\") {")
	fmt.Println("        // 发送退款政策并提供客服小程序卡片")
	fmt.Println("        customer.SendText(userID, \"关于退款政策...\") ")
	fmt.Println("        customer.SendMiniProgramPage(userID, \"退款申请\", \"pages/refund/apply\", \"media_id\") ")
	fmt.Println("    } else {")
	fmt.Println("        // 默认回复")
	fmt.Println("        customer.SendText(userID, \"您好，请问有什么可以帮助您的？\") ")
	fmt.Println("    }")
	fmt.Println("}")

	// 示例：多媒体消息组合
	fmt.Println("\n多媒体消息组合示例:")
	fmt.Println("func sendProductInfo(userID, productID string) {")
	fmt.Println("    // 1. 发送产品文本介绍")
	fmt.Println("    customer.SendText(userID, \"这是我们的热销产品...\") ")
	fmt.Println("    ")
	fmt.Println("    // 2. 发送产品图片")
	fmt.Println("    customer.SendImage(userID, \"product_image_media_id\") ")
	fmt.Println("    ")
	fmt.Println("    // 3. 发送产品详情链接")
	fmt.Println("    customer.SendLink(userID, \"产品详情\", \"查看完整规格和参数\", \"pages/product/detail?id=\"+productID, \"thumb_url\") ")
	fmt.Println("    ")
	fmt.Println("    // 4. 发送购买入口小程序卡片")
	fmt.Println("    customer.SendMiniProgramPage(userID, \"立即购买\", \"pages/product/buy?id=\"+productID, \"buy_thumb_media_id\") ")
	fmt.Println("}")
}

/* 使用说明

1. 客服消息概述：
   - 客服消息是小程序与用户进行消息交互的能力
   - 支持多种消息类型：文本、图片、图文链接、小程序卡片等
   - 可用于用户咨询、订单通知、活动推广等场景
   - 需要用户主动发起会话才能发送消息

2. 消息类型：
   a. 文本消息
      - 纯文本内容
      - 适用于简单的问答和通知
      - 支持表情符号和换行

   b. 图片消息
      - 需要先上传临时素材获取媒体ID
      - 图片格式支持JPG、PNG
      - 临时素材有效期为3天

   c. 图文链接消息
      - 包含标题、描述、链接和缩略图
      - 可跳转到小程序内页面或外部网页
      - 适合展示详细信息

   d. 小程序卡片消息
      - 包含标题、页面路径和缩略图
      - 点击后直接打开小程序指定页面
      - 适合引导用户完成特定操作

3. 客服输入状态：
   - 可以设置"正在输入"状态
   - 提升用户等待体验
   - 处理完成后可取消输入状态

4. 临时素材管理：
   - 获取客服消息中的临时素材
   - 临时素材有效期为3天
   - 可用于保存用户发送的图片等内容

5. 使用场景：
   - 用户咨询回复
   - 订单状态通知
   - 活动推广
   - 售后服务
   - 个性化推荐

6. 最佳实践：
   - 及时响应用户消息
   - 合理组合不同类型的消息
   - 设计自动回复系统
   - 实现人工客服交接
   - 保存会话历史记录

7. 注意事项：
   - 消息发送频率限制
   - 会话时效性（48小时）
   - 临时素材有效期
   - 消息内容合规性
   - 用户体验优化

8. 错误处理：
   - 45047：客服接口下行条数超
