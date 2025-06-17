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

	// 获取数据分析服务
	analysis := mp.Analysis

	// 设置日期范围
	now := time.Now()
	endDate := now.AddDate(0, 0, -1)                // 昨天
	beginDate := endDate.AddDate(0, 0, -6)          // 7天前
	beginDateMonth := endDate.AddDate(0, -1, 0)     // 1个月前
	beginDateWeek := endDate.AddDate(0, 0, -7*4+1)  // 4周前

	// 1. 获取小程序概况趋势
	fmt.Println("=== 小程序概况趋势 ===")
	fmt.Printf("日期范围: %s 至 %s\n", beginDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	summaryList, err := analysis.GetDailySummary(beginDate, endDate)
	if err != nil {
		log.Printf("获取概况趋势失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 条概况趋势数据:\n", len(summaryList))
		for _, item := range summaryList {
			// 将 yyyymmdd 格式转换为 yyyy-mm-dd 格式
			date := fmt.Sprintf("%s-%s-%s", item.RefDate[:4], item.RefDate[4:6], item.RefDate[6:])
			fmt.Printf("  - 日期: %s\n", date)
			fmt.Printf("    累计用户数: %d\n", item.VisitTotal)
			fmt.Printf("    转发次数: %d\n", item.SharePV)
			fmt.Printf("    转发人数: %d\n", item.ShareUV)
			fmt.Println()
		}
	}

	// 2. 获取小程序日访问趋势
	fmt.Println("=== 小程序日访问趋势 ===")
	fmt.Printf("日期范围: %s 至 %s\n", beginDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	dailyTrendList, err := analysis.GetDailyVisitTrend(beginDate, endDate)
	if err != nil {
		log.Printf("获取日访问趋势失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 条日访问趋势数据:\n", len(dailyTrendList))
		for _, item := range dailyTrendList {
			date := fmt.Sprintf("%s-%s-%s", item.RefDate[:4], item.RefDate[4:6], item.RefDate[6:])
			fmt.Printf("  - 日期: %s\n", date)
			fmt.Printf("    打开次数: %d\n", item.SessionCnt)
			fmt.Printf("    访问次数: %d\n", item.VisitPV)
			fmt.Printf("    访问人数: %d\n", item.VisitUV)
			fmt.Printf("    新用户数: %d\n", item.VisitUVNew)
			fmt.Printf("    人均停留时长: %.2f 秒\n", item.StayTimeUV)
			fmt.Printf("    次均停留时长: %.2f 秒\n", item.StayTimeSession)
			fmt.Printf("    平均访问深度: %.2f\n", item.VisitDepth)
			fmt.Println()
		}
	}

	// 3. 获取小程序周访问趋势
	fmt.Println("=== 小程序周访问趋势 ===")
	fmt.Printf("日期范围: %s 至 %s\n", beginDateWeek.Format("2006-01-02"), endDate.Format("2006-01-02"))

	weeklyTrendList, err := analysis.GetWeeklyVisitTrend(beginDateWeek, endDate)
	if err != nil {
		log.Printf("获取周访问趋势失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 条周访问趋势数据:\n", len(weeklyTrendList))
		for _, item := range weeklyTrendList {
			date := fmt.Sprintf("%s-%s-%s", item.RefDate[:4], item.RefDate[4:6], item.RefDate[6:])
			fmt.Printf("  - 周开始日期: %s\n", date)
			fmt.Printf("    打开次数: %d\n", item.SessionCnt)
			fmt.Printf("    访问次数: %d\n", item.VisitPV)
			fmt.Printf("    访问人数: %d\n", item.VisitUV)
			fmt.Printf("    新用户数: %d\n", item.VisitUVNew)
			fmt.Printf("    人均停留时长: %.2f 秒\n", item.StayTimeUV)
			fmt.Printf("    次均停留时长: %.2f 秒\n", item.StayTimeSession)
			fmt.Printf("    平均访问深度: %.2f\n", item.VisitDepth)
			fmt.Println()
		}
	}

	// 4. 获取小程序月访问趋势
	fmt.Println("=== 小程序月访问趋势 ===")
	fmt.Printf("日期范围: %s 至 %s\n", beginDateMonth.Format("2006-01-02"), endDate.Format("2006-01-02"))

	monthlyTrendList, err := analysis.GetMonthlyVisitTrend(beginDateMonth, endDate)
	if err != nil {
		log.Printf("获取月访问趋势失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 条月访问趋势数据:\n", len(monthlyTrendList))
		for _, item := range monthlyTrendList {
			date := fmt.Sprintf("%s-%s-%s", item.RefDate[:4], item.RefDate[4:6], item.RefDate[6:])
			fmt.Printf("  - 月开始日期: %s\n", date)
			fmt.Printf("    打开次数: %d\n", item.SessionCnt)
			fmt.Printf("    访问次数: %d\n", item.VisitPV)
			fmt.Printf("    访问人数: %d\n", item.VisitUV)
			fmt.Printf("    新用户数: %d\n", item.VisitUVNew)
			fmt.Printf("    人均停留时长: %.2f 秒\n", item.StayTimeUV)
			fmt.Printf("    次均停留时长: %.2f 秒\n", item.StayTimeSession)
			fmt.Printf("    平均访问深度: %.2f\n", item.VisitDepth)
			fmt.Println()
		}
	}

	// 5. 获取小程序访问分布
	fmt.Println("=== 小程序访问分布 ===")
	fmt.Printf("日期: %s\n", endDate.Format("2006-01-02"))

	distribution, err := analysis.GetVisitDistribution(endDate)
	if err != nil {
		log.Printf("获取访问分布失败: %v", err)
	} else {
		fmt.Println("访问来源分布:")
		for _, item := range distribution.AccessSourceList {
			fmt.Printf("  - 来源: %s\n", item.AccessSource)
			fmt.Printf("    打开次数: %d\n", item.SessionCnt)
			fmt.Printf("    占比: %.2f%%\n", item.SessionPct*100)
		}

		fmt.Println("\n访问时长分布:")
		for _, item := range distribution.AccessStayTimeList {
			fmt.Printf("  - 时长区间: %s\n", item.StayTime)
			fmt.Printf("    打开次数: %d\n", item.SessionCnt)
			fmt.Printf("    占比: %.2f%%\n", item.SessionPct*100)
		}

		fmt.Println("\n访问深度分布:")
		for _, item := range distribution.AccessDepthList {
			fmt.Printf("  - 访问深度: %s\n", item.VisitDepth)
			fmt.Printf("    打开次数: %d\n", item.SessionCnt)
			fmt.Printf("    占比: %.2f%%\n", item.SessionPct*100)
		}
	}

	// 6. 获取小程序访问页面
	fmt.Println("\n=== 小程序访问页面 ===")
	fmt.Printf("日期范围: %s 至 %s\n", beginDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	pageList, err := analysis.GetVisitPage(beginDate, endDate)
	if err != nil {
		log.Printf("获取访问页面失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 条页面访问数据:\n", len(pageList))
		for i, item := range pageList {
			if i >= 5 { // 只显示前5条数据
				fmt.Println("  ... 更多页面数据省略 ...")
				break
			}

			date := fmt.Sprintf("%s-%s-%s", item.RefDate[:4], item.RefDate[4:6], item.RefDate[6:])
			fmt.Printf("  - 日期: %s\n", date)
			fmt.Printf("    页面路径: %s\n", item.PagePath)
			fmt.Printf("    访问次数: %d\n", item.PageVisitPV)
			fmt.Printf("    访问人数: %d\n", item.PageVisitUV)
			fmt.Printf("    平均停留时长: %.2f 秒\n", item.PageStayTime)
			fmt.Printf("    进入页次数: %d\n", item.EntryPagePV)
			fmt.Printf("    退出页次数: %d\n", item.ExitPagePV)
			fmt.Printf("    转发次数: %d\n", item.PageSharePV)
			fmt.Printf("    转发人数: %d\n", item.PageShareUV)
			fmt.Println()
		}
	}

	// 7. 获取小程序用户画像
	fmt.Println("=== 小程序用户画像 ===")
	fmt.Printf("日期范围: %s 至 %s\n", beginDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	portrait, err := analysis.GetUserPortrait(beginDate, endDate)
	if err != nil {
		log.Printf("获取用户画像失败: %v", err)
	} else {
		fmt.Println("性别分布:")
		for _, item := range portrait.GenderList {
			fmt.Printf("  - %s: %d 人 (%.2f%%)\n", item.Name, item.Value, item.Ratio*100)
		}

		fmt.Println("\n年龄分布:")
		for _, item := range portrait.AgeList {
			fmt.Printf("  - %s: %d 人 (%.2f%%)\n", item.Name, item.Value, item.Ratio*100)
		}

		fmt.Println("\n地域分布 (前5名):")
		for i, item := range portrait.ProvinceList {
			if i >= 5 { // 只显示前5条数据
				fmt.Println("  ... 更多地域数据省略 ...")
				break
			}
			fmt.Printf("  - %s: %d 人 (%.2f%%)\n", item.Name, item.Value, item.Ratio*100)
		}

		fmt.Println("\n终端分布:")
		for _, item := range portrait.PlatformList {
			fmt.Printf("  - %s: %d 人 (%.2f%%)\n", item.Name, item.Value, item.Ratio*100)
		}
	}

	// 8. 数据分析最佳实践
	fmt.Println("\n=== 数据分析最佳实践 ===")

	// 示例：计算关键指标
	fmt.Println("计算关键指标示例:")

	// 假设我们有一周的数据
	if len(dailyTrendList) > 0 {
		var totalVisitUV, totalVisitUVNew, totalSessionCnt int
		var totalStayTime float64

		for _, item := range dailyTrendList {
			totalVisitUV += item.VisitUV
			totalVisitUVNew += item.VisitUVNew
			totalSessionCnt += item.SessionCnt
			totalStayTime += float64(item.VisitUV) * item.StayTimeUV
		}

		// 计算平均每日访问人数
		avgDailyVisitUV := float64(totalVisitUV) / float64(len(dailyTrendList))

		// 计算新用户占比
		newUserRatio := float64(totalVisitUVNew) / float64(totalVisitUV) * 100

		// 计算人均访问次数
		avgVisitPerUser := float64(totalSessionCnt) / float64(totalVisitUV)

		// 计算总体平均停留时长
		avgStayTime := totalStayTime / float64(totalVisitUV)

		fmt.Printf("  - 平均每日访问人数: %.2f\n", avgDailyVisitUV)
		fmt.Printf("  - 新用户占比: %.2f%%\n", newUserRatio)
		fmt.Printf("  - 人均访问次数: %.2f\n", avgVisitPerUser)
		fmt.Printf("  - 总体平均停留时长: %.2f 秒\n", avgStayTime)
	}

	//