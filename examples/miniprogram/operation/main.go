package main

import (
	"fmt"
	"log"
	"time"

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

	// 获取运营服务
	operation := mp.Operation

	// 1. 获取域名配置
	fmt.Println("=== 获取域名配置 ===")

	domainInfo, err := operation.GetDomainInfo()
	if err != nil {
		log.Printf("获取域名配置失败: %v", err)
	} else {
		fmt.Println("域名配置信息:")

		fmt.Println("\nRequest合法域名:")
		if len(domainInfo.RequestDomain) > 0 {
			for i, domain := range domainInfo.RequestDomain {
				fmt.Printf("  %d. %s\n", i+1, domain)
			}
		} else {
			fmt.Println("  (无)")
		}

		fmt.Println("\nWebSocket合法域名:")
		if len(domainInfo.WSRequestDomain) > 0 {
			for i, domain := range domainInfo.WSRequestDomain {
				fmt.Printf("  %d. %s\n", i+1, domain)
			}
		} else {
			fmt.Println("  (无)")
		}

		fmt.Println("\nUpload合法域名:")
		if len(domainInfo.UploadDomain) > 0 {
			for i, domain := range domainInfo.UploadDomain {
				fmt.Printf("  %d. %s\n", i+1, domain)
			}
		} else {
			fmt.Println("  (无)")
		}

		fmt.Println("\nDownload合法域名:")
		if len(domainInfo.DownloadDomain) > 0 {
			for i, domain := range domainInfo.DownloadDomain {
				fmt.Printf("  %d. %s\n", i+1, domain)
			}
		} else {
			fmt.Println("  (无)")
		}

		fmt.Println("\n业务域名:")
		if len(domainInfo.BizDomain) > 0 {
			for i, domain := range domainInfo.BizDomain {
				fmt.Printf("  %d. %s\n", i+1, domain)
			}
		} else {
			fmt.Println("  (无)")
		}
	}

	// 2. 修改域名配置
	fmt.Println("\n=== 修改域名配置 ===")

	// 创建新的域名配置
	newDomainInfo := &miniprogram.DomainInfo{
		RequestDomain: []string{
			"https://api.example.com",
			"https://api.yourdomain.com",
		},
		WSRequestDomain: []string{
			"wss://socket.example.com",
		},
		UploadDomain: []string{
			"https://upload.example.com",
		},
		DownloadDomain: []string{
			"https://download.example.com",
		},
	}

	fmt.Println("准备修改域名配置:")

	fmt.Println("\nRequest合法域名:")
	for i, domain := range newDomainInfo.RequestDomain {
		fmt.Printf("  %d. %s\n", i+1, domain)
	}

	fmt.Println("\nWebSocket合法域名:")
	for i, domain := range newDomainInfo.WSRequestDomain {
		fmt.Printf("  %d. %s\n", i+1, domain)
	}

	fmt.Println("\nUpload合法域名:")
	for i, domain := range newDomainInfo.UploadDomain {
		fmt.Printf("  %d. %s\n", i+1, domain)
	}

	fmt.Println("\nDownload合法域名:")
	for i, domain := range newDomainInfo.DownloadDomain {
		fmt.Printf("  %d. %s\n", i+1, domain)
	}

	// 注释掉实际的API调用，避免在示例中真正修改域名配置
	/*
		err = operation.ModifyDomain(newDomainInfo, "set")
		if err != nil {
			log.Printf("修改域名配置失败: %v", err)
		} else {
			fmt.Println("\n域名配置修改成功")
		}
	*/
	fmt.Println("\n注意：这里仅作为示例，实际使用时请取消注释")

	// 3. 设置业务域名
	fmt.Println("\n=== 设置业务域名 ===")

	// 业务域名列表
	bizDomains := []string{
		"https://h5.example.com",
		"https://web.yourdomain.com",
	}

	fmt.Println("准备设置业务域名:")
	for i, domain := range bizDomains {
		fmt.Printf("  %d. %s\n", i+1, domain)
	}

	// 注释掉实际的API调用，避免在示例中真正设置业务域名
	/*
		err = operation.SetWebviewDomain(bizDomains, "set")
		if err != nil {
			log.Printf("设置业务域名失败: %v", err)
		} else {
			fmt.Println("\n业务域名设置成功")
		}
	*/
	fmt.Println("\n注意：这里仅作为示例，实际使用时请取消注释")

	// 4. 获取性能数据
	fmt.Println("\n=== 获取性能数据 ===")

	// 设置时间范围
	now := time.Now()
	endTime := now.Format("20060102")
	startTime := now.AddDate(0, 0, -7).Format("20060102") // 7天前

	fmt.Printf("获取性能数据 (时间范围: %s 至 %s)\n", startTime, endTime)

	performanceData, err := operation.GetPerformanceData(startTime, endTime)
	if err != nil {
		log.Printf("获取性能数据失败: %v", err)
	} else {
		fmt.Printf("\n获取到 %d 条性能数据:\n", len(performanceData))

		for i, data := range performanceData {
			fmt.Printf("\n数据 %d:\n", i+1)
			fmt.Printf("  - 时间范围: %s\n", data.TimeRange)
			fmt.Printf("  - 页面访问次数: %d\n", data.PageCount)
			fmt.Printf("  - 平均耗时: %.2f 毫秒\n", data.AvgTime)
			fmt.Printf("  - 错误率: %.2f%%\n", data.ErrorRate*100)
		}
	}

	// 5. 获取服务状态
	fmt.Println("\n=== 获取服务状态 ===")

	serverStatus, err := operation.GetServerStatus()
	if err != nil {
		log.Printf("获取服务状态失败: %v", err)
	} else {
		fmt.Println("服务状态信息:")
		fmt.Printf("  - 状态码: %d\n", serverStatus.Status)
		fmt.Printf("  - 状态描述: %s\n", serverStatus.Description)
		fmt.Printf("  - 错误次数: %d\n", serverStatus.ErrorCount)

		if serverStatus.LastError != "" {
			fmt.Printf("  - 最后一次错误: %s\n", serverStatus.LastError)
		} else {
			fmt.Println("  - 最后一次错误: (无)")
		}
	}

	// 6. 运营最佳实践
	fmt.Println("\n=== 运营最佳实践 ===")

	// 示例：域名配置管理
	fmt.Println("域名配置管理:")
	fmt.Println("1. 定期检查域名配置")
	fmt.Println("2. 使用HTTPS协议保证安全")
	fmt.Println("3. 确保域名有效且可访问")
	fmt.Println("4. 避免频繁修改域名配置")
	fmt.Println("5. 业务域名需要在小程序管理后台配置业务域名校验文件")

	// 示例：性能监控
	fmt.Println("\n性能监控:")
	fmt.Println("1. 定期检查小程序性能数据")
	fmt.Println("2. 关注页面加载时间和错误率")
	fmt.Println("3. 针对性能问题进行优化")
	fmt.Println("4. 设置性能告警阈值")
	fmt.Println("5. 分析用户体验与性能的关系")

	// 示例：服务状态监控
	fmt.Println("\n服务状态监控:")
	fmt.Println("1. 实现自动化监控系统")
	fmt.Println("2. 设置状态告警机制")
	fmt.Println("3. 制定故障恢复流程")
	fmt.Println("4. 定期进行服务压力测试")
	fmt.Println("5. 建立服务可用性报告")

	// 示例：实现代码
	fmt.Println("\n实现代码示例:")
	fmt.Println("```go")
	fmt.Println("// 域名配置检查工具")
	fmt.Println("func checkDomainConfig() {")
	fmt.Println("    // 获取当前域名配置")
	fmt.Println("    domainInfo, err := operation.GetDomainInfo()")
	fmt.Println("    if err != nil {")
	fmt.Println("        log.Printf(\"获取域名配置失败: %v\", err)")
	fmt.Println("        return")
	fmt.Println("    }")
	fmt.Println("    ")
	fmt.Println("    // 检查必要的域名是否配置")
	fmt.Println("    requiredDomains := []string{\"https://api.yourdomain.com\"}")
	fmt.Println("    missingDomains := []string{}")
	fmt.Println("    ")
	fmt.Println("    for _, required := range requiredDomains {")
	fmt.Println("        found := false")
	fmt.Println("        for _, domain := range domainInfo.RequestDomain {")
	fmt.Println("            if domain == required {")
	fmt.Println("                found = true")
	fmt.Println("                break")
	fmt.Println("            }")
	fmt.Println("        }")
	fmt.Println("        ")
	fmt.Println("        if !found {")
	fmt.Println("            missingDomains = append(missingDomains, required)")
	fmt.Println("        }")
	fmt.Println("    }")
	fmt.Println("    ")
	fmt.Println("    // 处理缺失的域名")
	fmt.Println("    if len(missingDomains) > 0 {")
	fmt.Println("        log.Printf(\"发现 %d 个必要域名未配置\", len(missingDomains))")
	fmt.Println("        ")
	fmt.Println("        // 添加缺失的域名")
	fmt.Println("        newDomains := append(domainInfo.RequestDomain, missingDomains...)")
	fmt.Println("        domainInfo.RequestDomain = newDomains")
	fmt.Println("        ")
	fmt.Println("        // 更新域名配置")
	fmt.Println("        err = operation.ModifyDomain(domainInfo, \"set\")")
	fmt.Println("        if err != nil {")
	fmt.Println("            log.Printf(\"更新域名配置失败: %v\", err)")
	fmt.Println("        } else {")
	fmt.Println("            log.Println(\"域名配置已更新\")")
	fmt.Println("        }")
	fmt.Println("    } else {")
	fmt.Println("        log.Println(\"域名配置检查通过\")")
	fmt.Println("    }")
	fmt.Println("}")
	fmt.Println("```")

	// 示例：性能监控工具
	fmt.Println("\n性能监控工具:")
	fmt.Println("```go")
	fmt.Println("// 性能监控报告")
	fmt.Println("func generatePerformanceReport() {")
	fmt.Println("    // 设置时间范围")
	fmt.Println("    now := time.Now()")
	fmt.Println("    endTime := now.Format(\"20060102\")")
	fmt.Println("    startTime := now.AddDate(0, 0, -30).Format(\"20060102\") // 30天")
	fmt.Println("    ")
	fmt.Println("    // 获取性能数据")
	fmt.Println("    data, err := operation.GetPerformanceData(startTime, endTime)")
	fmt.Println("    if err != nil {")
	fmt.Println("        log.Printf(\"获取性能数据失败: %v\", err)")
	fmt.Println("        return")
	fmt.Println("    }")
	fmt.Println("    ")
	fmt.Println("    // 分析性能数据")
	fmt.Println("    var totalPageCount int")
	fmt.Println("    var totalAvgTime float64")
	fmt.Println("    var maxErrorRate float64")
	fmt.Println("    var worstDay string")
	fmt.Println("    ")
	fmt.Println("    for _, item := range data {")
	fmt.Println("        totalPageCount += item.PageCount")
	fmt.Println("        totalAvgTime += item.AvgTime * float64(item.PageCount)")
	fmt.Println("        ")
	fmt.Println("        if item.ErrorRate > maxErrorRate {")
	fmt.Println("            maxErrorRate = item.ErrorRate")
	fmt.Println("            worstDay = item.TimeRange")
	fmt.Println("        }")
	fmt.Println("    }")
	fmt.Println("    ")
	fmt.Println("    // 计算平均值")
	fmt.Println("    avgTime := totalAvgTime / float64(totalPageCount)")
	fmt.Println("    ")
	fmt.Println("    // 生成报告")
	fmt.Println("    fmt.Printf(\"性能监控报告 (%s - %s)\\n\", startTime, endTime)")
	fmt.Println("    fmt.Printf(\"总访问次数: %d\\n\", totalPageCount)")
	fmt.Println("    fmt.Printf(\"平均响应时间: %.2f 毫秒\\n\", avgTime)")
	fmt.Println("    fmt.Printf(\"最高错误率: %.2f%% (发生在 %s)\\n\", maxErrorRate*100, worstDay)")
	fmt.Println("    ")
	fmt.Println("    // 性能评估")
	fmt.Println("    if avgTime > 1000 { // 超过1秒")
	fmt.Println("        fmt.Println(\"警告：平均响应时间过高，需要优化\")")
	fmt.Println("    }")
	fmt.Println("    ")
	fmt.Println("    if maxErrorRate > 0.01 { // 错误率超过1%")
	fmt.Println("        fmt.Printf(\"警告：在 %s 发现较高错误率，请检查原因\\n\", worstDay)")
	fmt.Println("    }")
	fmt.Println("}")
	fmt.Println("```")
	fmt.Println("\n/* 使用说明")
	fmt.Println("")
	fmt.Println("1. 运营中心概述：")
	fmt.Println("   - 运营中心提供小程序运营相关的功能")
	fmt.Println("   - 包括域名管理、性能监控、服务状态等")
	fmt.Println("   - 帮助开发者更好地管理和优化小程序")
	fmt.Println("")
	fmt.Println("2. 主要功能：")
	fmt.Println("   a. 域名管理")
	fmt.Println("      - 获取和修改域名配置")
	fmt.Println("      - 设置业务域名")
	fmt.Println("      - 管理各类合法域名")
	fmt.Println("   ")
	fmt.Println("   b. 性能监控")
	fmt.Println("      - 获取性能数据")
	fmt.Println("      - 分析访问情况")
	fmt.Println("      - 监控错误率")
	fmt.Println("   ")
	fmt.Println("   c. 服务状态")
	fmt.Println("      - 监控服务状态")
	fmt.Println("      - 跟踪错误信息")
	fmt.Println("      - 保障服务质量")
	fmt.Println("")
	fmt.Println("3. 使用场景：")
	fmt.Println("   - 小程序发布前的域名配置")
	fmt.Println("   - 运营过程中的性能优化")
	fmt.Println("   - 服务质量监控和告警")
	fmt.Println("   - 问题排查和分析")
	fmt.Println("")
	fmt.Println("4. 注意事项：")
	fmt.Println("   - 域名必须经过ICP备案")
	fmt.Println("   - 及时处理性能问题")
	fmt.Println("   - 保持服务稳定性")
	fmt.Println("   - 做好监控和告警")
	fmt.Println("*/")
	fmt.Println("}")
