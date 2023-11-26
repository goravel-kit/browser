package browser

import (
	"github.com/goravel/framework/contracts/foundation"
)

const Binding = "browser"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return NewBrowser(app.MakeConfig()), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	app.Publishes("github.com/goravel-kit/browser", map[string]string{
		"config/browser.go": app.ConfigPath("browser.go"),
	})
}
