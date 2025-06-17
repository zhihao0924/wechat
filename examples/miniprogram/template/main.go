package main

import (
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

	// 获取模板消息服务
	template := mp.Template

	// 重要提示
	fmt.Println("=== 重要提示 ===")
	fmt.Println("根据微信官方公告，小程序模板消息已于2020年1月10日下线，请使用订阅消息代替。")
	fmt.Println("本示例仅作为历史参考，实际开发中请使用订阅消息功能。")
	fmt.Println("订阅消息的使用方法请参考 examples/miniprogram/subscribe/main.go")
	fmt.Println()

	// 1. 创建模板数据
	fmt.Println("=== 创建模板数据 ===")

	data := template.CreateTemplateData()

	// 添加模板数据项
	template.AddTemplateData(data, "keyword1", "订单已支付", "#173177")
	template.AddTemplateData(data, "keyword2", "￥88.00", "#173177")
	template.AddTemplateData(data, "keyword3", "2023-05-20 15:30:00", "#173177")
	template.AddTemplateData(data, "keyword4", "感谢您的购买，祝您使用愉快！")

	fmt.Println("模板数据创建成功:")
	for key, item := range data {
		if itemMap, ok := item.(miniprogram.TemplateItem); ok {
			if itemMap.Color != "" {
				fmt.Printf("  - %s: %s (颜色: %s)\n", key, itemMap.Value, itemMap.Color)
			} else {
				fmt.Printf("  - %s: %s\n", key, itemMap.Value)
			}
		}
	}

	// 2. 发送模板消息（快速方式）
	fmt.Println("\n=== 发送模板消息（快速方式）===")

	toUser := "oWm5H5Cxxxxxxxx"           // 替换为接收者的OpenID
	templateID := "ABC123DEF456GHI789JKL" // 替换为模板ID
	formID := "FORM_ID_123456789"         // 替换为表单ID或支付ID
	page := "pages/order/detail?id=123"   // 点击模板卡片后的跳转页面

	fmt.Printf("发送模板消息给用户 %s\n", toUser)
	fmt.Printf("使用模板ID: %s\n", templateID)
	fmt.Printf("表单ID: %s\n", formID)
	fmt.Printf("跳转页面: %s\n", page)

	// 注释掉实际的API调用，避免在示例中真正发送消息
	/*
		err := template.SendTemplate(toUser, templateID, formID, data, page)
		if err != nil {
			log.Printf("发送模板消息失败: %v", err)
		} else {
			fmt.Println("发送模板消息成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 3. 使用完整的模板消息结构体发送消息
	fmt.Println("\n=== 使用完整的模板消息结构体发送消息 ===")

	// 创建模板消息
	message := &miniprogram.TemplateMessage{
		ToUser:          toUser,
		TemplateID:      templateID,
		FormID:          formID,
		Page:            page,
		Data:            data,
		EmphasisKeyword: "keyword1.DATA", // 需要放大的关键词
	}

	fmt.Println("使用完整的模板消息结构体:")
	fmt.Printf("  - 接收者: %s\n", message.ToUser)
	fmt.Printf("  - 模板ID: %s\n", message.TemplateID)
	fmt.Printf("  - 表单ID: %s\n", message.FormID)
	fmt.Printf("  - 跳转页面: %s\n", message.Page)
	fmt.Printf("  - 放大关键词: %s\n", message.EmphasisKeyword)
	fmt.Println("  - 数据项:")
	for key, item := range message.Data {
		if itemMap, ok := item.(miniprogram.TemplateItem); ok {
			if itemMap.Color != "" {
				fmt.Printf("    * %s: %s (颜色: %s)\n", key, itemMap.Value, itemMap.Color)
			} else {
				fmt.Printf("    * %s: %s\n", key, itemMap.Value)
			}
		}
	}

	// 注释掉实际的API调用，避免在示例中真正发送消息
	/*
		err = template.Send(message)
		if err != nil {
			log.Printf("发送模板消息失败: %v", err)
		} else {
			fmt.Println("发送模板消息成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 4. 模板消息最佳实践（历史参考）
	fmt.Println("\n=== 模板消息最佳实践（历史参考）===")

	// 示例：获取formID
	fmt.Println("获取formID的方式:")
	fmt.Println("1. 表单提交场景：通过form组件的submit事件获取formId")
	fmt.Println("   <form report-submit=\"true\" bindsubmit=\"formSubmit\">")
	fmt.Println("   // 在formSubmit函数中获取e.detail.formId")
	fmt.Println("2. 支付场景：使用支付成功后的prepay_id作为formId")

	// 示例：模板消息使用场景
	fmt.Println("\n模板消息使用场景（历史）:")
	fmt.Println("1. 订单状态通知（下单成功、支付成功、发货通知等）")
	fmt.Println("2. 服务进度通知（预约成功、服务开始、服务完成等）")
	fmt.Println("3. 活动通知（活动开始、活动结束、中奖通知等）")

	// 示例：模板消息注意事项
	fmt.Println("\n模板消息注意事项（历史）:")
	fmt.Println("1. formId的有效期为7天")
	fmt.Println("2. 每个formId只能使用一次")
	fmt.Println("3. 模板消息的内容应简洁明了")
	fmt.Println("4. 模板消息有调用频率限制")

	// 示例：迁移到订阅消息
	fmt.Println("\n迁移到订阅消息:")
	fmt.Println("1. 使用订阅消息替代模板消息")
	fmt.Println("2. 订阅消息需要用户主动订阅")
	fmt.Println("3. 订阅消息分为一次性订阅和长期订阅")
	fmt.Println("4. 订阅消息的使用方法请参考 examples/miniprogram/subscribe/main.go")
}

/* 使用说明

1. 重要提示：
   - 小程序模板消息已于2020年1月10日下线
   - 请使用订阅消息代替模板消息
   - 本示例仅作为历史参考，实际开发中请使用订阅消息功能

2. 模板消息概述（历史）：
   - 模板消息是小程序向用户发送通知消息的能力
   - 模板消息需要用户主动触发（如表单提交、支付等）
   - 模板消息有固定的格式和样式
   - 模板消息可以包含多个数据项

3. 模板消息使用流程（历史）：
   a. 获取formId
      - 表单提交场景：通过form组件的submit事件获取formId
      - 支付场景：使用支付成功后的prepay_id作为formId
   b. 准备模板数据
      - 创建模板数据对象
      - 添加各个关键词的数据项
      - 设置数据项的值和颜色
   c. 发送模板消息
      - 调用发送接口
      - 指定接收者、模板ID、formID和数据
      - 可选设置跳转页面和放大关键词

4. 模板消息数据项：
   - 关键词：模板中定义的变量名
   - 值：要显示的具体内容
   - 颜色：文字颜色，使用十六进制颜色码（如#173177）

5. 模板消息限制（历史）：
   - formId的有效期为7天
   - 每个formId只能使用一次
   - 模板消息有调用频率限制
   - 模板消息的内容应简洁明了

6. 迁移到订阅消息：
   - 使用订阅消息替代模板消息
   - 订阅消息需要用户主动订阅
   - 订阅消息分为一次性订阅和长期订阅
   - 订阅消息的使用方法请参考 examples/miniprogram/subscribe/main.go

7. 订阅消息与模板消息的区别：
   - 获取方式：订阅消息需要用户主动订阅，模板消息需要用户触发行为
   - 发送次数：一次性订阅消息只能发送一次，长期订阅消息可以发送多次
   - 有效期：订阅消息的订阅有效期为7天，模板消息的formId有效期为7天
   - 使用场景：订阅消息适用于更广泛的场景，模板消息仅适用于用户触发行为后的通知

8. 常见问题（历史）：
   - 41028：form_id不正确，或者过期
   - 41029：form_id已被使用
   - 41030：page不正确
   - 41031：接口调用超过限制
   - 45009：接口调用超过限额

9. 开发建议：
   - 使用订阅消息替代模板消息
   - 合理规划消息场景
   - 优化订阅流程体验
   - 遵循发送频率限制
   - 处理发送失败情况
*/
