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

	// 获取订阅消息服务
	subscribe := mp.Subscribe

	// 示例用户OpenID
	userOpenID := "oXYZ123_example_openid" // 替换为实际用户的OpenID

	// 1. 获取当前帐号所设置的类目信息
	fmt.Println("=== 获取类目信息 ===")

	categories, err := subscribe.GetCategory()
	if err != nil {
		log.Printf("获取类目信息失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 个类目:\n", len(categories))
		for i, category := range categories {
			fmt.Printf("  %d. ID: %d, 名称: %s\n", i+1, category.ID, category.Name)
		}
	}

	// 2. 获取模板标题列表
	fmt.Println("\n=== 获取模板标题列表 ===")

	// 假设我们使用第一个类目的ID
	var categoryID string
	if len(categories) > 0 {
		categoryID = fmt.Sprintf("%d", categories[0].ID)
	} else {
		categoryID = "2" // 示例类目ID，实际使用时应替换为真实的类目ID
	}

	start := 0
	limit := 10

	fmt.Printf("获取类目 %s 的模板标题列表 (起始位置: %d, 条数: %d)\n", categoryID, start, limit)

	templateTitles, total, err := subscribe.GetPubTemplateTitles(categoryID, start, limit)
	if err != nil {
		log.Printf("获取模板标题列表失败: %v", err)
	} else {
		fmt.Printf("总模板标题数量: %d\n", total)
		fmt.Printf("获取到 %d 个模板标题:\n", len(templateTitles))

		for i, title := range templateTitles {
			fmt.Printf("  %d. ID: %d, 标题: %s\n", i+1, title.TID, title.Title)
			fmt.Printf("     类型: %s, 类目ID: %s\n", getTemplateTypeDesc(title.Type), title.CategoryID)
		}
	}

	// 3. 获取模板标题下的关键词库
	fmt.Println("\n=== 获取模板关键词库 ===")

	// 假设我们使用第一个模板标题的ID
	var templateTitleID int
	if len(templateTitles) > 0 {
		templateTitleID = templateTitles[0].TID
	} else {
		templateTitleID = 12345 // 示例模板标题ID，实际使用时应替换为真实的模板标题ID
	}

	fmt.Printf("获取模板标题 %d 的关键词库\n", templateTitleID)

	keywords, err := subscribe.GetPubTemplateKeywords(templateTitleID)
	if err != nil {
		log.Printf("获取模板关键词库失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 个关键词:\n", len(keywords))

		for i, keyword := range keywords {
			fmt.Printf("  %d. ID: %d, 名称: %s\n", i+1, keyword.KID, keyword.Name)
			fmt.Printf("     示例: %s, 规则: %s\n", keyword.Example, keyword.Rule)
		}
	}

	// 4. 添加模板
	fmt.Println("\n=== 添加模板 ===")

	// 选择关键词ID列表
	var kidList []int
	if len(keywords) >= 2 {
		kidList = []int{keywords[0].KID, keywords[1].KID}
	} else {
		kidList = []int{1, 2} // 示例关键词ID，实际使用时应替换为真实的关键词ID
	}

	sceneDesc := "订单状态通知" // 模板用途描述

	fmt.Printf("添加模板:\n")
	fmt.Printf("  - 模板标题ID: %d\n", templateTitleID)
	fmt.Printf("  - 关键词ID列表: %v\n", kidList)
	fmt.Printf("  - 用途描述: %s\n", sceneDesc)

	// 注释掉实际的API调用，避免在示例中真正添加模板
	/*
		priTmplID, err := subscribe.AddTemplate(templateTitleID, kidList, sceneDesc)
		if err != nil {
			log.Printf("添加模板失败: %v", err)
		} else {
			fmt.Printf("添加模板成功，模板ID: %s\n", priTmplID)
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 假设添加成功，使用示例模板ID
	priTmplID := "XYZ123_example_template_id" // 替换为实际的模板ID

	// 5. 获取个人模板列表
	fmt.Println("\n=== 获取个人模板列表 ===")

	templates, err := subscribe.GetTemplateList()
	if err != nil {
		log.Printf("获取个人模板列表失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 个个人模板:\n", len(templates))

		for i, template := range templates {
			fmt.Printf("\n模板 %d:\n", i+1)
			fmt.Printf("  - 模板ID: %s\n", template.PriTmplID)
			fmt.Printf("  - 标题: %s\n", template.Title)
			fmt.Printf("  - 内容: %s\n", template.Content)
			fmt.Printf("  - 示例: %s\n", template.Example)
			fmt.Printf("  - 类型: %s\n", getTemplateTypeDesc(template.Type))
		}
	}

	// 6. 发送订阅消息
	fmt.Println("\n=== 发送订阅消息 ===")

	// 创建订阅消息数据
	data := subscribe.CreateSubscribeData()

	// 添加数据项
	subscribe.AddSubscribeData(data, "character_string1", "OD12345678")      // 订单号
	subscribe.AddSubscribeData(data, "phrase2", "已发货")                    // 订单状态
	subscribe.AddSubscribeData(data, "thing3", "感谢您的购买，祝您使用愉快！") // 备注
	subscribe.AddSubscribeData(data, "time4", "2023-05-15 14:30:00")         // 发货时间

	// 跳转页面
	page := "pages/order/detail?id=OD12345678"

	fmt.Printf("发送订阅消息给用户 %s:\n", userOpenID)
	fmt.Printf("  - 模板ID: %s\n", priTmplID)
	fmt.Printf("  - 跳转页面: %s\n", page)
	fmt.Printf("  - 数据项:\n")
	for key, item := range data {
		fmt.Printf("    %s: %s\n", key, item.Value)
	}

	// 注释掉实际的API调用，避免在示例中真正发送消息
	/*
		err = subscribe.SendSubscribe(userOpenID, priTmplID, data, page)
		if err != nil {
			log.Printf("发送订阅消息失败: %v", err)
		} else {
			fmt.Println("发送订阅消息成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 7. 删除模板
	fmt.Println("\n=== 删除模板 ===")

	fmt.Printf("删除模板 (ID: %s)\n", priTmplID)

	// 注释掉实际的API调用，避免在示例中真正删除模板
	/*
		err = subscribe.DeleteTemplate(priTmplID)
		if err != nil {
			log.Printf("删除模板失败: %v", err)
		} else {
			fmt.Println("删除模板成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 8. 订阅消息最佳实践
	fmt.Println("\n=== 订阅消息最佳实践 ===")

	// 示例：订阅消息流程
	fmt.Println("订阅消息流程:")
	fmt.Println("1. 在小程序端获取用户订阅授权")
	fmt.Println("2. 在合适的业务场景触发消息发送")
	fmt.Println("3. 根据模板格式构建消息数据")
	fmt.Println("4. 通过API发送订阅消息")
	fmt.Println("5. 处理发送结果和错误")

	// 示例：常见业务场景
	fmt.Println("\n常见业务场景:")
	fmt.Println("1. 订单状态变更通知")
	fmt.Println("2. 物流配送提醒")
	fmt.Println("3. 活动开始提醒")
	fmt.Println("4. 预约成功通知")
	fmt.Println("5. 服务进度更新")

	// 示例：前端订阅代码
	fmt.Println("\n前端订阅代码示例:")
	fmt.Println("```javascript")
	fmt.Println("// 获取用户订阅授权")
	fmt.Println("wx.requestSubscribeMessage({")
	fmt.Println("  tmplIds: ['模板ID1', '模板ID2'],")
	fmt.Println("  success: (res) => {")
	fmt.Println("    if (res['模板ID1'] === 'accept') {")
	fmt.Println("      // 用户同意订阅")
	fmt.Println("      // 可以在此处或通过接口将订阅状态保存到后端")
	fmt.Println("    }")
	fmt.Println("  },")
	fmt.Println("  fail: (err) => {")
	fmt.Println("    console.error('订阅消息请求失败', err)")
	fmt.Println("  }")
	fmt.Println("})")
	fmt.Println("```")

	// 示例：后端发送代码
	fmt.Println("\n后端发送代码示例:")
	fmt.Println("```go")
	fmt.Println("func sendOrderStatusNotification(orderID, userOpenID string) {")
	fmt.Println("  // 获取订单信息")
	fmt.Println("  order := getOrderInfo(orderID)")
	fmt.Println("  ")
	fmt.Println("  // 创建订阅消息数据")
	fmt.Println("  data := subscribe.CreateSubscribeData()")
	fmt.Println("  subscribe.AddSubscribeData(data, \"character_string1\", order.OrderID)")
	fmt.Println("  subscribe.AddSubscribeData(data, \"phrase2\", order.StatusText)")
	fmt.Println("  subscribe.AddSubscribeData(data, \"thing3\", order.Remark)")
	fmt.Println("  subscribe.AddSubscribeData(data, \"time4\", order.UpdateTime)")
	fmt.Println("  ")
	fmt.Println("  // 发送订阅消息")
	fmt.Println("  page := \"pages/order/detail?id=\" + order.OrderID")
	fmt.Println("  err := subscribe.SendSubscribe(userOpenID, \"订单状态模板ID\", data, page)")
	fmt.Println("  if err != nil {")
	fmt.Println("    log.Printf(\"发送订阅消息失败: %v\", err)")
	fmt.Println("  }")
	fmt.Println("}")
	fmt.Println("```")
}

// 获取模板类型描述
func getTemplateTypeDesc(typeCode int) string {
	switch typeCode {
	case 2:
		return "一次性订阅"
	case 3:
		return "长期订阅"
	default:
		return "未知类型"
	}
}

/* 使用说明

1. 订阅消息概述：
   - 订阅消息是小程序向用户发送通知消息的能力
   - 需要用户主动订阅才能发送
   - 分为一次性订阅和长期订阅两种类型
   - 适用于服务通知、状态变更、活动提醒等场景

2. 订阅消息类型：
   a. 一次性订阅
      - 用户订阅一次，开发者可发送一条消息
      - 发送后订阅关系自动解除
      - 适用于单次交易、活动提醒等场景

   b. 长期订阅
      - 用户订阅一次，开发者可多次发送消息
      - 订阅关系长期有效
      - 适用于持续性服务、定期提醒等场景

3. 模板管理：
   a. 类目和模板
      - 每个小程序可以选择多个类目
      - 每个类目下有多个模板标题
      - 每个模板标题有多个关键词
      - 开发者可以选择关键词
