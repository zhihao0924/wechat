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

	// 获取直播服务
	live := mp.Live

	// 1. 获取直播间列表
	fmt.Println("=== 获取直播间列表 ===")

	start := 0
	limit := 10

	fmt.Printf("获取直播间列表 (起始位置: %d, 条数: %d)\n", start, limit)

	rooms, total, err := live.GetRoomList(start, limit)
	if err != nil {
		log.Printf("获取直播间列表失败: %v", err)
	} else {
		fmt.Printf("总直播间数量: %d\n", total)
		fmt.Printf("获取到 %d 个直播间:\n", len(rooms))

		for i, room := range rooms {
			fmt.Printf("\n直播间 %d:\n", i+1)
			fmt.Printf("  - 直播间ID: %d\n", room.RoomID)
			fmt.Printf("  - 直播间名称: %s\n", room.Name)
			fmt.Printf("  - 主播昵称: %s\n", room.AnchorName)

			// 格式化时间
			startTime := time.Unix(room.StartTime, 0).Format("2006-01-02 15:04:05")
			endTime := time.Unix(room.EndTime, 0).Format("2006-01-02 15:04:05")

			fmt.Printf("  - 开始时间: %s\n", startTime)
			fmt.Printf("  - 结束时间: %s\n", endTime)

			// 直播间状态
			status := "未知"
			switch room.Status {
			case 101:
				status = "直播中"
			case 102:
				status = "未开始"
			case 103:
				status = "已结束"
			case 104:
				status = "禁播"
			case 105:
				status = "暂停中"
			case 106:
				status = "异常"
			case 107:
				status = "已过期"
			}

			fmt.Printf("  - 直播状态: %s\n", status)
			fmt.Printf("  - 背景图片: %s\n", room.CoverImg)
			fmt.Printf("  - 分享图片: %s\n", room.ShareImg)
		}
	}

	// 2. 获取直播间回放
	fmt.Println("\n=== 获取直播间回放 ===")

	// 假设我们使用第一个直播间的ID
	var roomID int
	if len(rooms) > 0 {
		roomID = rooms[0].RoomID
	} else {
		roomID = 1 // 示例ID，实际使用时应替换为真实的直播间ID
	}

	start = 0
	limit = 10

	fmt.Printf("获取直播间回放 (直播间ID: %d, 起始位置: %d, 条数: %d)\n", roomID, start, limit)

	videos, total, err := live.GetReplay(roomID, start, limit)
	if err != nil {
		log.Printf("获取直播间回放失败: %v", err)
	} else {
		fmt.Printf("总回放视频数量: %d\n", total)
		fmt.Printf("获取到 %d 个回放视频:\n", len(videos))

		for i, video := range videos {
			fmt.Printf("  %d. %s\n", i+1, video)
		}
	}

	// 3. 获取直播间商品列表
	fmt.Println("\n=== 获取直播间商品列表 ===")

	// 商品状态：0-未审核，1-审核中，2-审核通过，3-审核驳回
	status := 2 // 审核通过的商品
	start = 0
	limit = 10

	fmt.Printf("获取直播间商品列表 (状态: %d, 起始位置: %d, 条数: %d)\n", status, start, limit)

	goods, total, err := live.GetGoods(status, start, limit)
	if err != nil {
		log.Printf("获取直播间商品列表失败: %v", err)
	} else {
		fmt.Printf("总商品数量: %d\n", total)
		fmt.Printf("获取到 %d 个商品:\n", len(goods))

		for i, item := range goods {
			fmt.Printf("\n商品 %d:\n", i+1)
			fmt.Printf("  - 商品ID: %d\n", item.GoodsID)
			fmt.Printf("  - 商品名称: %s\n", item.Name)
			fmt.Printf("  - 商品价格: %.2f\n", item.Price)
			fmt.Printf("  - 商品图片: %s\n", item.CoverImg)
			fmt.Printf("  - 商品链接: %s\n", item.URL)
			fmt.Printf("  - 第三方ID: %s\n", item.ThirdPartyID)

			// 商品审核状态
			auditStatus := "未知"
			switch item.Status {
			case 0:
				auditStatus = "未审核"
			case 1:
				auditStatus = "审核中"
			case 2:
				auditStatus = "审核通过"
			case 3:
				auditStatus = "审核驳回"
			}

			fmt.Printf("  - 审核状态: %s\n", auditStatus)
		}
	}

	// 4. 添加商品并提交审核
	fmt.Println("\n=== 添加商品并提交审核 ===")

	// 创建商品信息
	newGoods := &miniprogram.LiveGoods{
		CoverImg:     "https://example.com/product.jpg", // 商品图片
		Name:         "示例商品",                        // 商品名称
		Price:        99.99,                             // 商品价格
		URL:          "pages/product/detail?id=123",     // 商品详情页链接
		ThirdPartyID: "product_123",                     // 第三方商品ID
	}

	fmt.Println("添加商品并提交审核:")
	fmt.Printf("  - 商品名称: %s\n", newGoods.Name)
	fmt.Printf("  - 商品价格: %.2f\n", newGoods.Price)
	fmt.Printf("  - 商品图片: %s\n", newGoods.CoverImg)
	fmt.Printf("  - 商品链接: %s\n", newGoods.URL)
	fmt.Printf("  - 第三方ID: %s\n", newGoods.ThirdPartyID)

	// 注释掉实际的API调用，避免在示例中真正添加商品
	/*
		goodsID, err := live.AddGoods(newGoods)
		if err != nil {
			log.Printf("添加商品失败: %v", err)
		} else {
			fmt.Printf("添加商品成功，商品ID: %d\n", goodsID)
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 5. 更新商品
	fmt.Println("\n=== 更新商品 ===")

	// 假设我们要更新的商品ID
	goodsID := 12345 // 示例ID，实际使用时应替换为真实的商品ID

	// 更新商品信息
	updateGoods := &miniprogram.LiveGoods{
		GoodsID:      goodsID,
		CoverImg:     "https://example.com/product_updated.jpg", // 更新后的商品图片
		Name:         "示例商品（更新版）",                        // 更新后的商品名称
		Price:        88.88,                                     // 更新后的商品价格
		URL:          "pages/product/detail?id=123",             // 商品详情页链接
		ThirdPartyID: "product_123",                             // 第三方商品ID
	}

	fmt.Printf("更新商品 (ID: %d):\n", goodsID)
	fmt.Printf("  - 商品名称: %s\n", updateGoods.Name)
	fmt.Printf("  - 商品价格: %.2f\n", updateGoods.Price)
	fmt.Printf("  - 商品图片: %s\n", updateGoods.CoverImg)

	// 注释掉实际的API调用，避免在示例中真正更新商品
	/*
		err = live.UpdateGoods(updateGoods)
		if err != nil {
			log.Printf("更新商品失败: %v", err)
		} else {
			fmt.Println("更新商品成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 6. 撤回商品审核
	fmt.Println("\n=== 撤回商品审核 ===")

	fmt.Printf("撤回商品审核 (ID: %d)\n", goodsID)

	// 注释掉实际的API调用，避免在示例中真正撤回商品审核
	/*
		err = live.ResetAuditGoods(goodsID)
		if err != nil {
			log.Printf("撤回商品审核失败: %v", err)
		} else {
			fmt.Println("撤回商品审核成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 7. 重新提交商品审核
	fmt.Println("\n=== 重新提交商品审核 ===")

	fmt.Printf("重新提交商品审核 (ID: %d)\n", goodsID)

	// 注释掉实际的API调用，避免在示例中真正重新提交商品审核
	/*
		err = live.ResubmitAuditGoods(goodsID)
		if err != nil {
			log.Printf("重新提交商品审核失败: %v", err)
		} else {
			fmt.Println("重新提交商品审核成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 8. 删除商品
	fmt.Println("\n=== 删除商品 ===")

	fmt.Printf("删除商品 (ID: %d)\n", goodsID)

	// 注释掉实际的API调用，避免在示例中真正删除商品
	/*
		err = live.DeleteGoods(goodsID)
		if err != nil {
			log.Printf("删除商品失败: %v", err)
		} else {
			fmt.Println("删除商品成功")
		}
	*/
	fmt.Println("注意：这里仅作为示例，实际使用时请取消注释")

	// 9. 直播功能最佳实践
	fmt.Println("\n=== 直播功能最佳实践 ===")

	// 示例：直播间管理
	fmt.Println("直播间管理:")
	fmt.Println("1. 在微信公众平台后台创建直播间")
	fmt.Println("2. 设置直播间基本信息（名称、封面、时间等）")
	fmt.Println("3. 通过API获取直播间列表")
	fmt.Println("4. 在小程序中展示直播间列表")
	fmt.Println("5. 引导用户进入直播间观看")

	// 示例：商品管理
	fmt.Println("\n商品管理:")
	fmt.Println("1. 提前准备商品信息（图片、名称、价格等）")
	fmt.Println("2. 通过API添加商品并提交审核")
	fmt.Println("3. 定期检查商品审核状态")
	fmt.Println("4. 将审核通过的商品关联到直播间")
	fmt.Println("5. 在直播过程中上架/下架商品")

	// 示例：直播数据分析
	fmt.Println("\n直播数据分析:")
	fmt.Println("1. 记录直播间观看人数和互动数据")
	fmt.Println("2. 分析商品点击率和转化率")
	fmt.Println("3. 评估直播效果和ROI")
	fmt.Println("4. 根据数据优化下次直播策略")
}

/* 使用说明

1. 小程序直播概述：
   - 小程序直播是微信为小程序提供的直播能力
   - 支持商品展示、用户互动、数据分析等功能
   - 需要在微信公众平台开通直播权限
   - 适用于电商、教育、活动等场景

2. 直播间管理：
   a. 创建直播间
      - 在微信公众平台后台创建
      - 设置直播间名称、封面、时间等信息
      - 可以通过API获取直播间列表

   b. 直播间状态
      - 101：直播中
      - 102：未开始
      - 103：已结束
      - 104：禁播
      - 105：暂停中
      - 106：异常
      - 107：已过期

   c. 直播
