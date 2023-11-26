package facades

import (
	"log"

	"github.com/goravel-kit/browser"
	"github.com/goravel-kit/browser/contracts"
)

func Browser() contracts.Browser {
	instance, err := browser.App.Make(browser.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Browser)
}
