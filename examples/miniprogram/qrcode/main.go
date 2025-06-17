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

	// 获取二维码服务
	qrcode := mp.QRCode

	// 创建保存二维码的目录
	qrDir := "qrcodes"
	if err := os.MkdirAll(qrDir, 0755); err != nil {
		log.Fatalf("创建目录失败: %v", err)
	}

	// 1. 基本小程序码生成
	fmt.Println("=== 生成基本小程序码 ===")

	scene := "product=123&source=qr"
	page := "pages/product/detail"
	filename := filepath.Join(qrDir, "basic_qrcode.png")

	fmt.Printf("生成基本小程序码:\n")
	fmt.Printf("  - 场景值: %s\n", scene)
	fmt.Printf("  - 页面路径: %s\n", page)
	fmt.Printf("  - 保存路径: %s\n", filename)

	err := qrcode.CreateQRCodeToFile(scene, page, filename)
	if err != nil {
		log.Printf("生成基本小程序码失败: %v", err)
	} else {
		fmt.Println("生成基本小程序码成功")
	}

	// 2. 自定义样式的小程序码生成
	fmt.Println("\n=== 生成自定义样式的小程序码 ===")

	// 创建自定义参数
	params := &miniprogram.QRCodeParams{
		Scene:     "promotion=summer&id=456",
		Page:      "pages/promotion/detail",
		Width:     600,   // 设置宽度为600px
		AutoColor: false, // 不使用自动颜色
		IsHyaline: true,  // 使用透明底色
	}

	// 设置自定义颜色（蓝色）
	params.LineColor.R = "0"
	params.LineColor.G = "0"
	params.LineColor.B = "255"

	filename = filepath.Join(qrDir, "custom_qrcode.png")

	fmt.Printf("生成自定义样式小程序码:\n")
	fmt.Printf("  - 场景值: %s\n", params.Scene)
	fmt.Printf("  - 页面路径: %s\n", params.Page)
	fmt.Printf("  - 宽度: %dpx\n", params.Width)
	fmt.Printf("  - 颜色: R=%s, G=%s, B=%s\n", params.LineColor.R, params.LineColor.G, params.LineColor.B)
	fmt.Printf("  - 透明底色: %v\n", params.IsHyaline)
	fmt.Printf("  - 保存路径: %s\n", filename)

	err = qrcode.GetQRCodeToFile(params, filename)
	if err != nil {
		log.Printf("生成自定义样式小程序码失败: %v", err)
	} else {
		fmt.Println("生成自定义样式小程序码成功")
	}

	// 3. 批量生成不同场景的小程序码
	fmt.Println("\n=== 批量生成不同场景的小程序码 ===")

	// 定义多个场景
	scenes := []struct {
		Scene string
		Page  string
		Name  string
	}{
		{
			Scene: "type=clothing&id=789",
			Page:  "pages/category/clothing",
			Name:  "clothing",
		},
		{
			Scene: "type=electronics&id=101",
			Page:  "pages/category/electronics",
			Name:  "electronics",
		},
		{
			Scene: "type=food&id=202",
			Page:  "pages/category/food",
			Name:  "food",
		},
	}

	fmt.Println("开始批量生成小程序码:")

	for _, s := range scenes {
		filename := filepath.Join(qrDir, fmt.Sprintf("%s_qrcode.png", s.Name))

		fmt.Printf("\n生成 %s 类别小程序码:\n", s.Name)
		fmt.Printf("  - 场景值: %s\n", s.Scene)
		fmt.Printf("  - 页面路径: %s\n", s.Page)
		fmt.Printf("  - 保存路径: %s\n", filename)

		// 使用自定义宽度和颜色
		color := struct{ R, G, B string }{"30", "144", "255"} // 道奇蓝
		err := qrcode.CreateQRCodeToFile(s.Scene, s.Page, filename, 500, false, color, false)

		if err != nil {
			log.Printf("生成 %s 类别小程序码失败: %v", s.Name, err)
		} else {
			fmt.Printf("生成 %s 类别小程序码成功\n", s.Name)
		}
	}

	// 4. 小程序码最佳实践
	fmt.Println("\n=== 小程序码最佳实践 ===")

	// 示例：小程序码生成策略
	fmt.Println("小程序码生成策略:")
	fmt.Println("1. 场景值设计")
	fmt.Println("   - 使用简短且有意义的参数名")
	fmt.Println("   - 避免使用特殊字符")
	fmt.Println("   - 合理组织参数顺序")
	fmt.Println("   - 考虑参数长度限制（最大32个可见字符）")

	fmt.Println("\n2. 页面路径规范")
	fmt.Println("   - 确保页面已发布")
	fmt.Println("   -使用正确的路径格式")
	fmt.Println("   - 不要在路径中包含参数")
	fmt.Println("   - 避免使用未经授权的页面")

	fmt.Println("\n3. 样式优化")
	fmt.Println("   - 选择合适的尺寸")
	fmt.Println("   - 考虑扫码场景")
	fmt.Println("   - 适当使用自定义颜色")
	fmt.Println("   - 根据背景选择是否使用透明底色")

	fmt.Println("\n4. 使用建议")
	fmt.Println("   - 提前生成并缓存码")
	fmt.Println("   - 避免频繁生成相同的码")
	fmt.Println("   - 实现错误重试机制")
	fmt.Println("   - 做好日志记录")

	// 示例：场景值设计
	fmt.Println("\n场景值设计示例:")
	fmt.Println("1. 商品详情页")
	fmt.Println("   scene: product=123&source=qr")
	fmt.Println("   page: pages/product/detail")

	fmt.Println("\n2. 活动页面")
	fmt.Println("   scene: activity=new_year&uid=456")
	fmt.Println("   page: pages/activity/main")

	fmt.Println("\n3. 分享邀请")
	fmt.Println("   scene: invite=789&from=user123")
	fmt.Println("   page: pages/register/index")

	// 示例：实现代码
	fmt.Println("\n实现代码示例:")
	fmt.Println("```go")
	fmt.Println("// 商品二维码生成器")
	fmt.Println("func generateProductQRCode(productID string, width int) (string, error) {")
	fmt.Println("    // 构建场景值")
	fmt.Println("    scene := fmt.Sprintf(\"product=%s&source=qr\", productID)")
	fmt.Println("    ")
	fmt.Println("    // 设置页面路径")
	fmt.Println("    page := \"pages/product/detail\"")
	fmt.Println("    ")
	fmt.Println("    // 生成文件名")
	fmt.Println("    filename := fmt.Sprintf(\"product_%s_qr.png\", productID)")
	fmt.Println("    ")
	fmt.Println("    // 设置自定义颜色")
	fmt.Println("    color := struct{R,G,B string}{\"30\", \"144\", \"255\"}")
	fmt.Println("    ")
	fmt.Println("    // 生成二维码")
	fmt.Println("    err := qrcode.CreateQRCodeToFile(scene, page, filename, width, false, color, false)")
	fmt.Println("    if err != nil {")
	fmt.Println("        return \"\", fmt.Errorf(\"生成商品二维码失败: %v\", err)")
	fmt.Println("    }")
	fmt.Println("    ")
	fmt.Println("    return filename, nil")
	fmt.Println("}")
	fmt.Println("```")
}

/* 使用说明

1. 小程序码概述：
   - 小程序码是小程序的二维码，用于线下推广和访问
   - 支持自定义参数和样式
   - 永久有效，数量暂无限制
   - 适用于各种营销和引流场景

2. 主要功能：
   a. 基本功能
      - 生成小程序码
      - 自定义场景值
      - 指定跳转页面
      - 保存到文件

   b. 自定义选项
      - 设置码的宽度
      - 自定义颜色
      - 配置透明底色
      - 自动颜色适配

3. 参数说明：
   a. 必需参数
      - scene: 场景值，最大32个字符
      - page: 跳转页面路径

   b. 可选参数
      - width: 码的宽度（280-1280px）
      - autoColor: 是否自动配置颜色
      - lineColor: 自定义颜色
      - isHyaline: 是否使用透明底色

4. 使用场景：
   - 商品详情页
   - 活动推广
   - 用户分享
   - 门店导流
   - 会员卡券

5. 注意事项：
   - 确保页面已发布
   - 场景值长度限制
   - 正确的页面路径
   - 合适的码大小
   - 清晰的扫描体验

6. 错误处理：
   - 40001：获取access_token失败
   - 41030：所传page页面不存在
   - 45009：调用频率超限
   - 其他错误码

7. 最佳实践：
   - 合理设计场景值
   - 优化扫码体验
   - 做好缓存机制
   - 实现错误重试
   - 监控生成情况

8. 应用扩展：
   - 批量生成工具
   - 码管理系统
   - 扫码统计分析
   - A/B测试支持
   - 防伪验证
*/
