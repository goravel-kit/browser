package contracts

import "github.com/go-rod/rod"

type Browser interface {
	Get(slug string, withLock bool) *rod.Browser
}
