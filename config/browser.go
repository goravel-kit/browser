package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("browser", map[string]any{
		// 管理地址
		"manage_url": "",
		// 跟踪模式
		"trace": true,
		// 无头
		"headless": true,
		// 自动打开开发者工具
		"devtools": false,
		// 忽略证书错误
		"ignore_cert_errors": false,
	})
}
