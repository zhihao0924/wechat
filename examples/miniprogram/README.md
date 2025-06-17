# 微信小程序SDK示例

本目录包含了微信小程序SDK的各种功能示例，帮助开发者快速了解和使用SDK提供的功能。

## 目录结构

```
examples/miniprogram/
├── customer/       # 客服消息示例
├── subscribe/      # 订阅消息示例
├── qrcode/         # 小程序码示例
├── security/       # 内容安全检测示例
├── operation/      # 运营中心示例
└── README.md       # 本文件
```

## 使用说明

1. 克隆仓库到本地
   ```bash
   git clone https://github.com/zhihao0924/wechat.git
   cd wechat
   ```

2. 配置小程序信息
   在运行示例前，请将示例代码中的 `your_app_id` 和 `your_app_secret` 替换为你的小程序 AppID 和 AppSecret。

3. 运行示例
   ```bash
   cd examples/miniprogram/[示例目录]
   go run main.go
   ```

## 功能说明

### 1. 客服消息 (customer)

`customer/main.go` 展示了如何使用客服消息功能：

- 发送文本消息
- 发送图片消息
- 发送链接消息
- 发送小程序卡片消息
- 上传和获取临时素材

客服消息可以帮助小程序与用户进行互动，提供更好的用户体验。

### 2. 订阅消息 (subscribe)

`subscribe/main.go` 展示了如何使用订阅消息功能：

- 获取类目信息
- 获取模板标题列表
- 获取模板关键词
- 添加模板
- 获取模板列表
- 发送订阅消息
- 删除模板

订阅消息是小程序向用户发送通知的重要方式，可以在用户同意的情况下，向其推送相关信息。

### 3. 小程序码 (qrcode)

`qrcode/main.go` 展示了如何生成小程序码：

- 生成基本小程序码
- 生成自定义样式的小程序码
- 批量生成不同场景的小程序码
- 小程序码生成的最佳实践

小程序码可以帮助用户快速访问小程序的特定页面，是线下推广和用户分享的重要工具。

### 4. 内容安全检测 (security)

`security/main.go` 展示了如何使用内容安全检测功能：

- 文本内容安全检测
- 图片内容安全检测
- 媒体文件异步检测
- 风险处理策略

内容安全检测可以帮助开发者过滤不良信息，保障小程序内容的健康和合规。

### 5. 运营中心 (operation)

`operation/main.go` 展示了如何使用运营中心功能：

- 获取和修改域名配置
- 设置业务域名
- 获取性能数据
- 获取服务状态
- 运营最佳实践

运营中心功能可以帮助开发者更好地管理和优化小程序，提升用户体验。

## 注意事项

1. **安全性**
   - 请勿在代码中硬编码 AppID 和 AppSecret，建议使用环境变量或配置文件
   - 在生产环境中，请妥善保管密钥信息

2. **接口限制**
   - 微信官方对接口调用频率有限制，请合理控制调用频率
   - 部分接口需要特定的小程序类目或认证，请查阅官方文档确认

3. **内容合规**
   - 使用内容安全检测功能时，仍需人工复核高风险内容
   - 确保小程序内容符合微信平台规范和相关法律法规

4. **用户体验**
   - 订阅消息发送应尊重用户选择，不要频繁发送或发送与用户无关的内容
   - 客服消息应及时响应用户查询，提供良好的服务体验

## 最佳实践

1. **错误处理**
   - 示例代码中包含了基本的错误处理，实际应用中应更完善地处理各种异常情况
   - 建议实现日志记录和监控机制，及时发现和解决问题

2. **性能优化**
   - 合理缓存访问令牌(access_token)，避免频繁请求
   - 对于高频调用的接口，考虑实现本地缓存机制

3. **代码组织**
   - 将业务逻辑与SDK调用分离，便于维护和测试
   - 使用依赖注入方式管理SDK实例，便于单元测试

## 贡献指南

欢迎贡献更多示例代码或改进现有示例：

1. Fork 本仓库
2. 创建你的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 开启一个 Pull Request

## 相关资源

- [微信官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/)
- [微信小程序开发指南](https://developers.weixin.qq.com/miniprogram/dev/framework/quickstart/)
- [微信开发者社区](https://developers.weixin.qq.com/community/develop/mixflow)

## 许可证

本示例代码遵循 MIT 许可证。详见 [LICENSE](../../LICENSE) 文件。
