package browser

import (
	"sync"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/goravel/framework/contracts/config"
)

type Browser struct {
	config   config.Config
	browsers map[string]*rod.Browser
	mu       sync.Mutex
	locks    map[string]*sync.Mutex
}

func NewBrowser(config config.Config) *Browser {
	return &Browser{
		config:   config,
		browsers: make(map[string]*rod.Browser),
		locks:    make(map[string]*sync.Mutex),
	}
}

func (r *Browser) Get(slug string, withLock bool) *rod.Browser {
	if withLock {
		r.ensureLock(slug)
		r.locks[slug].Lock()
		defer r.locks[slug].Unlock()
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if browser, exists := r.browsers[slug]; exists {
		return browser
	}

	l := launcher.New().Set("disable-blink-features", "AutomationControlled").
		Headless(r.config.GetBool("browser.headless", true)).
		Devtools(r.config.GetBool("browser.devtools", false))
	defer l.Cleanup()

	newBrowser := rod.New().
		ControlURL(l.MustLaunch()).
		DefaultDevice(devices.LaptopWithHiDPIScreen).
		Trace(r.config.GetBool("browser.trace", true)).
		MustConnect().
		MustIgnoreCertErrors(r.config.GetBool("browser.ignore_cert_errors", false))

	r.browsers[slug] = newBrowser
	return newBrowser
}

// ensureLock 确保每个 slug 都有一个对应的锁
func (r *Browser) ensureLock(slug string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.locks[slug]; !exists {
		r.locks[slug] = &sync.Mutex{}
	}
}
