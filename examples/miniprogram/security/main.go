package main

import (
	"fmt"
	"io/ioutil"
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

	// 获取安全服务
	security := mp.Security

	// 1. 文本内容安全检测
	fmt.Println("=== 文本内容安全检测 ===")

	// 示例文本
	textSamples := []struct {
		Description string
		Content     string
	}{
		{
			Description: "正常文本",
			Content:     "这是一段正常的文本内容，描述了小程序的功能和使用方法。",
		},
		{
			Description: "可能包含敏感内容的文本",
			Content:     "这是一段测试文本，包含一些敏感词如赌博、色情等，用于测试内容检测功能。",
		},
	}

	for _, sample := range textSamples {
		fmt.Printf("\n检测文本: %s\n", sample.Description)
		fmt.Printf("内容: %s\n", sample.Content)

		// 注释掉实际的API调用，避免在示例中真正调用检测API
		/*
			result, err := security.MsgSecCheck(sample.Content)
			if err != nil {
				log.Printf("文本检测失败: %v", err)
				continue
			}

			fmt.Printf("检测结果:\n")
			fmt.Printf("  - 风险等级: %s\n", security.GetRiskLevel(result))
			fmt.Printf("  - 风险类型: %s\n", security.GetRiskType(result))
			fmt.Printf("  - 是否有风险: %v\n", security.IsRisky(result))
			fmt.Printf("  - 跟踪ID: %s\n", result.TraceID)
		*/
		fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")
	}

	// 2. 图片内容安全检测
	fmt.Println("\n=== 图片内容安全检测 ===")

	// 创建示例图片目录
	imgDir := "sample_images"
	if err := os.MkdirAll(imgDir, 0755); err != nil {
		log.Printf("创建图片目录失败: %v", err)
	}

	// 示例图片路径
	imagePath := filepath.Join(imgDir, "sample.jpg")

	fmt.Printf("检测图片: %s\n", imagePath)
	fmt.Println("(假设这是一张需要检测的图片)")

	// 注释掉实际的API调用，避免在示例中真正调用检测API
	/*
		// 检查文件是否存在
		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			log.Printf("图片文件不存在: %s", imagePath)
		} else {
			result, err := security.ImgSecCheck(imagePath)
			if err != nil {
				log.Printf("图片检测失败: %v", err)
			} else {
				fmt.Printf("检测结果:\n")
				fmt.Printf("  - 风险等级: %s\n", security.GetRiskLevel(result))
				fmt.Printf("  - 风险类型: %s\n", security.GetRiskType(result))
				fmt.Printf("  - 是否有风险: %v\n", security.IsRisky(result))
				fmt.Printf("  - 跟踪ID: %s\n", result.TraceID)
			}
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 3. 图片字节数据检测
	fmt.Println("\n=== 图片字节数据检测 ===")

	fmt.Println("从内存中检测图片数据")
	fmt.Println("(假设这是从网络或其他来源获取的图片数据)")

	// 注释掉实际的API调用，避免在示例中真正调用检测API
	/*
		// 读取图片文件到内存
		imageData, err := ioutil.ReadFile(imagePath)
		if err != nil {
			log.Printf("读取图片文件失败: %v", err)
		} else {
			result, err := security.ImgSecCheckBytes(imageData)
			if err != nil {
				log.Printf("图片字节数据检测失败: %v", err)
			} else {
				fmt.Printf("检测结果:\n")
				fmt.Printf("  - 风险等级: %s\n", security.GetRiskLevel(result))
				fmt.Printf("  - 风险类型: %s\n", security.GetRiskType(result))
				fmt.Printf("  - 是否有风险: %v\n", security.IsRisky(result))
				fmt.Printf("  - 跟踪ID: %s\n", result.TraceID)
			}
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 4. 媒体文件异步检测
	fmt.Println("\n=== 媒体文件异步检测 ===")

	// 示例媒体URL
	mediaURL := "https://example.com/sample-audio.mp3"

	fmt.Printf("异步检测音频: %s\n", mediaURL)

	// 注释掉实际的API调用，避免在示例中真正调用检测API
	/*
		result, err := security.MediaCheckAsync(mediaURL, miniprogram.MediaTypeAudio)
		if err != nil {
			log.Printf("媒体文件异步检测失败: %v", err)
		} else {
			fmt.Printf("检测请求已提交:\n")
			fmt.Printf("  - 跟踪ID: %s\n", result.TraceID)
			fmt.Println("  - 检测结果将通过回调通知")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 示例图片URL
	imageURL := "https://example.com/sample-image.jpg"

	fmt.Printf("\n异步检测图片: %s\n", imageURL)

	// 注释掉实际的API调用，避免在示例中真正调用检测API
	/*
		result, err = security.MediaCheckAsync(imageURL, miniprogram.MediaTypeImage)
		if err != nil {
			log.Printf("媒体文件异步检测失败: %v", err)
		} else {
			fmt.Printf("检测请求已提交:\n")
			fmt.Printf("  - 跟踪ID: %s\n", result.TraceID)
			fmt.Println("  - 检测结果将通过回调通知")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 5. 内容安全最佳实践
	fmt.Println("\n=== 内容安全最佳实践 ===")

	// 示例：内容安全检测流程
	fmt.Println("内容安全检测流程:")
	fmt.Println("1. 用户生成内容（UGC）前置检测")
	fmt.Println("2. 内容发布前检测")
	fmt.Println("3. 定期对已发布内容进行检测")
	fmt.Println("4. 根据检测结果采取相应措施")
	fmt.Println("5. 建立人工审核机制")

	// 示例：风险处理策略
	fmt.Println("\n风险处理策略:")
	fmt.Println("1. 安全内容（pass）- 允许发布")
	fmt.Println("2. 有风险内容（risky）- 自动拦截")
	fmt.Println("3. 需要人工审核内容（review）- 进入审核队列")

	// 示例：实现代码
	fmt.Println("\n实现代码示例:")
	fmt.Println("```go")
	fmt.Println("// 用户评论安全检测")
	fmt.Println("func checkUserComment(userID, comment string) (bool, string) {")
	fmt.Println("    // 检测评论内容")
	fmt.Println("    result, err := security.MsgSecCheck(comment)")
	fmt.Println("    if err != nil {")
	fmt.Println("        log.Printf(\"评论检测失败: %v\", err)")
	fmt.Println("        return false, \"内容检测失败，请稍后重试\"")
	fmt.Println("    }")
	fmt.Println("    ")
	fmt.Println("    // 获取风险等级")
	fmt.Println("    riskLevel := security.GetRiskLevel(result)")
	fmt.Println("    ")
	fmt.Println("    // 根据风险等级处理")
	fmt.Println("    switch riskLevel {")
	fmt.Println("    case \"safe\":")
	fmt.Println("        // 安全内容，允许发布")
	fmt.Println("        return true, \"\"")
	fmt.Println("    case \"risky\":")
	fmt.Println("        // 有风险内容，拒绝发布")
	fmt.Println("        riskType := security.GetRiskType(result)")
	fmt.Println("        log.Printf(\"用户 %s 发布的评论包含 %s 内容\", userID, riskType)")
	fmt.Println("        return false, \"评论包含违规内容，请修改后重试\"")
	fmt.Println("    case \"review\":")
	fmt.Println("        // 需要人工审核，暂不发布")
	fmt.Println("        addToReviewQueue(userID, comment, result.TraceID)")
	fmt.Println("        return false, \"评论需要审核，审核通过后将自动发布\"")
	fmt.Println("    default:")
	fmt.Println("        // 未知情况，拒绝发布")
	fmt.Println("        return false, \"内容检测异常，请稍后重试\"")
	fmt.Println("    }")
	fmt.Println("}")
	fmt.Println("```")

	// 示例：图片检测工具函数
	fmt.Println("\n图片检测工具函数:")
	fmt.Println("```go")
	fmt.Println("// 检测用户上传的图片")
	fmt.Println("func checkUserImage(userID string, imageData []byte) (bool, string) {")
	fmt.Println("    // 检测图片内容")
	fmt.Println("    result, err := security.ImgSecCheckBytes(imageData)")
	fmt.Println("    if err != nil {")
	fmt.Println("        log.Printf(\"图片检测失败: %v\", err)")
	fmt.Println("        return false, \"图片检测失败，请稍后重试\"")
	fmt.Println("    }")
	fmt.Println("    ")
	fmt.Println("    // 判断是否有风险")
	fmt.Println("    if security.IsRisky(result) {")
	fmt.Println("        riskType := security.GetRiskType(result)")
	fmt.Println("        log.Printf(\"用户 %s 上传的图片包含 %s 内容\", userID, riskType)")
	fmt.Println("        return false, \"图片包含违规内容，请更换后重试\"")
	fmt.Println("    }")
	fmt.Println("    ")
	fmt.Println("    return true, \"\"")
	fmt.Println("}")
	fmt.Println("```")
}

/* 使用说明

1. 内容安全概述：
   - 内容安全检测是小程序保障平台内容合规的重要功能
   - 支持文本、图片、音频等多种内容类型的检测
   - 可以识别色情、暴力、政治敏感等多种违规内容
   - 适用于用户生成内容（UGC）的审核和过滤

2. 检测类型：
   a. 文本检测
      - 适用于用户评论、聊天消息等文本内容
      - 同步接口，实时返回检测结果
      - 支持多种违规类型识别

   b. 图片检测
      - 支持本地图片文件和内存中的图片数据
      - 同步接口，实时返回检测结果
      - 可识别图片中的违规内容

   c. 媒体异步检测
      - 支持音频和图片的异步检测
      - 通过回调接收检测结果
      - 适用于大文件或批量检测场景

3. 检测结果：
   a. 风险等级
      - safe: 内容安全
      - risky: 内容有风险
      - review: 需要人工审核

   b. 风险类型
      - 100: 正常
      - 10001: 广告
      - 20001: 时政
      - 20002: 色情
      - 20003: 辱骂
      - 20006: 违法犯罪
      - 20008: 欺诈
      - 20012: 低俗
      - 20013: 版权
      - 21000: 其他

4. 使用场景：
   - 用户评论审核
   - 社区内容检测
   - 图片上传过滤
   - 音频内容审核
