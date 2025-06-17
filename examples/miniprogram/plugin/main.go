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

	// 获取插件服务
	plugin := mp.Plugin

	// 1. 申请使用插件
	fmt.Println("=== 申请使用插件 ===")

	pluginAppID := "wx1234567890abcdef" // 示例插件AppID，实际使用时替换为真实的插件AppID
	reason := "需要使用该插件提供的地图选点功能"

	fmt.Printf("申请使用插件:\n")
	fmt.Printf("  - 插件AppID: %s\n", pluginAppID)
	fmt.Printf("  - 申请理由: %s\n", reason)

	// 注释掉实际的API调用，避免在示例中真正申请插件
	/*
		err := plugin.ApplyPlugin(pluginAppID, reason)
		if err != nil {
			log.Printf("申请使用插件失败: %v", err)
		} else {
			fmt.Println("申请使用插件成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 2. 查询已添加的插件
	fmt.Println("\n=== 查询已添加的插件 ===")

	pluginList, err := plugin.GetPluginList()
	if err != nil {
		log.Printf("查询已添加的插件失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 个插件:\n", len(pluginList))

		for i, p := range pluginList {
			fmt.Printf("\n插件 %d:\n", i+1)
			fmt.Printf("  - AppID: %s\n", p.AppID)
			fmt.Printf("  - 名称: %s\n", p.Nickname)
			fmt.Printf("  - 描述: %s\n", p.Description)
			fmt.Printf("  - 版本: %s\n", p.Version)

			// 插件状态
			status := "未知"
			switch p.Status {
			case 1:
				status = "申请中"
			case 2:
				status = "申请通过"
			case 3:
				status = "已拒绝"
			case 4:
				status = "已超时"
			}

			fmt.Printf("  - 状态: %s\n", status)
		}
	}

	// 3. 解除插件绑定
	fmt.Println("\n=== 解除插件绑定 ===")

	fmt.Printf("解除插件绑定:\n")
	fmt.Printf("  - 插件AppID: %s\n", pluginAppID)

	// 注释掉实际的API调用，避免在示例中真正解除插件绑定
	/*
		err = plugin.UnbindPlugin(pluginAppID)
		if err != nil {
			log.Printf("解除插件绑定失败: %v", err)
		} else {
			fmt.Println("解除插件绑定成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 4. 插件开发者功能（仅供插件开发者使用）
	fmt.Println("\n=== 插件开发者功能 ===")

	// 4.1 查询插件使用方的申请列表
	fmt.Println("\n查询插件使用方的申请列表:")

	page := 1
	num := 10

	fmt.Printf("获取申请列表 (页码: %d, 每页数量: %d)\n", page, num)

	applyList, err := plugin.GetDevApplyList(page, num)
	if err != nil {
		log.Printf("查询插件使用方的申请列表失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 个申请:\n", len(applyList))

		for i, apply := range applyList {
			fmt.Printf("\n申请 %d:\n", i+1)
			fmt.Printf("  - 申请者AppID: %s\n", apply.AppID)
			fmt.Printf("  - 申请者昵称: %s\n", apply.Nickname)
			fmt.Printf("  - 申请理由: %s\n", apply.Reason)

			// 申请状态
			status := "未知"
			switch apply.Status {
			case 1:
				status = "申请中"
			case 2:
				status = "已通过"
			case 3:
				status = "已拒绝"
			case 4:
				status = "已超时"
			}

			fmt.Printf("  - 状态: %s\n", status)
		}
	}

	// 4.2 修改插件使用申请的状态
	fmt.Println("\n修改插件使用申请的状态:")

	applyAppID := "wx9876543210fedcba" // 示例申请者AppID，实际使用时替换为真实的申请者AppID
	applyStatus := 2                   // 2-同意申请

	fmt.Printf("修改申请状态:\n")
	fmt.Printf("  - 申请者AppID: %s\n", applyAppID)
	fmt.Printf("  - 新状态: %d (2-同意，3-拒绝)\n", applyStatus)

	// 注释掉实际的API调用，避免在示例中真正修改申请状态
	/*
		err = plugin.SetDevPluginApplyStatus(applyAppID, applyStatus)
		if err != nil {
			log.Printf("修改插件使用申请的状态失败: %v", err)
		} else {
			fmt.Println("修改插件使用申请的状态成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 4.3 查询插件开发者的权限
	fmt.Println("\n查询插件开发者的权限:")

	permissions, err := plugin.GetPluginDevPermission()
	if err != nil {
		log.Printf("查询插件开发者的权限失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 个权限:\n", len(permissions))
		for i, perm := range permissions {
			fmt.Printf("  %d. %s\n", i+1, perm)
		}
	}

	// 5. 插件使用最佳实践
	fmt.Println("\n=== 插件使用最佳实践 ===")

	// 示例：插件使用流程
	fmt.Println("插件使用流程:")
	fmt.Println("1. 在微信公众平台查找所需插件")
	fmt.Println("2. 通过API申请使用插件")
	fmt.Println("3. 等待插件开发者审核")
	fmt.Println("4. 审核通过后在项目中引入插件")
	fmt.Println("5. 按照插件文档进行开发")
	fmt.Println("6. 需要时可以解除插件绑定")

	// 示例：插件开发者流程
	fmt.Println("\n插件开发者流程:")
	fmt.Println("1. 开发并发布插件")
	fmt.Println("2. 定期查看插件使用申请")
	fmt.Println("3. 审核并处理使用申请")
	fmt.Println("4. 提供插件使用文档和支持")
	fmt.Println("5. 维护和更新插件")
}

/* 使用说明

1. 小程序插件概述：
   - 小程序插件是能够被其他小程序引用的代码包
   - 提供了复用组件和功能的能力
   - 需要申请使用并获得插件开发者的同意
   - 适用于共享通用功能和组件

2. 插件使用方：
   a. 申请使用插件
      - 需要提供插件AppID
      - 填写申请理由
      - 等待插件开发者审核

   b. 插件状态
      - 1：申请中
      - 2：申请通过
      - 3：已拒绝
      - 4：已超时

   c. 管理插件
      - 查看已添加的插件列表
      - 必要时解除插件绑定
      - 在项目中正确引用插件

3. 插件开发者：
   a. 申请列表管理
      - 查看使用方的申请列表
      - 审核申请（同意/拒绝）
      - 管理申请状态

   b. 权限管理
      - 查看开发者权限
      - 确保合规使用
      - 维护插件安全

4. 使用场景：
   - 地图组件
   - 支付功能
   - 数据可视化
   - 广告组件
   - 分享功能

5. 最佳实践：
   - 仔细阅读插件文档
   - 遵循插件使用规范
   - 及时处理申请和审核
   - 保持插件版本更新
   - 做好异常处理

6. 注意事项：
   - 插件版本兼容性
   - 申请审核时间
   - 插件使用限制
   - 数据安全问题
   - 性能影响评估

7. 错误处理：
   - 85074：非法的插件AppID
   - 85075：非法的申请理由
   - 85076：插件已添加
   - 85077：插件不存在
   - 85078：非法的插件状态

8. 开发建议：
   - 合理选择插件
   - 做好版本管理
   - 实现降级方案
   - 监控插件性能
   - 及时更新插件

9. 应用扩展：
   - 插件使用统计
   - 自动化审核流程
   - 插件性能监控
   - 使用情况分析
   - 多插件协同管理
*/
