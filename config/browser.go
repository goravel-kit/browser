package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("browser", map[string]any{
		// 地址
		"url": "ws://",
		// 跟踪模式
		"trace": true,
		// 忽略证书错误
		"ignore_cert_errors": false,
	})
}
