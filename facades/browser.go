package facades

import (
	"log"

	"goravel/packages/browser"
	"goravel/packages/browser/contracts"
)

func Browser() contracts.Browser {
	instance, err := browser.App.Make(browser.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Browser)
}
