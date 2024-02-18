package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("browser", map[string]any{
		// 控制地址
		"control_url": "",
		// 跟踪模式
		"trace": true,
		// 忽略证书错误
		"ignore_cert_errors": false,
	})
}
