package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("browser", map[string]any{
		// 模式（control, manage）
		"mode": "control",
		// 控制地址
		"control_url": "ws://127.0.0.1:9222/devtools/browser/af491956-e276-43a4-a246-7522139ff69f",
		// 管理地址
		"manage_url": "",
		// 跟踪模式
		"trace": true,
		// 无头（仅 manage 模式生效）
		"headless": true,
		// 自动打开开发者工具（仅 manage 模式生效）
		"devtools": false,
		// 忽略证书错误
		"ignore_cert_errors": false,
	})
}
