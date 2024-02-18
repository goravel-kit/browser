package browser

import (
	"sync"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/goravel/framework/contracts/config"
)

type Browser struct {
	mu        sync.Mutex
	config    config.Config
	browsers  map[string]*rod.Browser
	launchers map[string]*launcher.Launcher
	locks     map[string]*sync.Mutex
}

func NewBrowser(config config.Config) *Browser {
	return &Browser{
		config:    config,
		browsers:  make(map[string]*rod.Browser),
		launchers: make(map[string]*launcher.Launcher),
		locks:     make(map[string]*sync.Mutex),
	}
}

// New 获取浏览器
func (r *Browser) New(slug string) *rod.Browser {
	r.ensureLock(slug)

	r.locks[slug].Lock()
	defer r.locks[slug].Unlock()

	r.mu.Lock()
	defer r.mu.Unlock()

	if browser, exists := r.browsers[slug]; exists {
		return browser
	}

	var l *launcher.Launcher
	newBrowser := rod.New()
	if r.config.GetString("browser.mode", "control") == "manage" {
		l = launcher.MustNewManaged(r.config.GetString("browser.manage_url")).
			Set("disable-blink-features", "AutomationControlled").
			Set("window-size", "1920,1080").
			Headless(r.config.GetBool("browser.headless", true)).
			Devtools(r.config.GetBool("browser.devtools", false))
		if !r.config.GetBool("browser.headless", true) {
			l = l.XVFB("--auto-servernum", "--server-args=-screen 0 1920x1080x24")
		}
		newBrowser = newBrowser.Client(l.MustClient())
	} else {
		newBrowser = newBrowser.
			ControlURL(r.config.GetString("browser.control_url", ""))
	}

	newBrowser = newBrowser.
		NoDefaultDevice().
		Trace(r.config.GetBool("browser.trace", true)).
		MustConnect().
		MustIgnoreCertErrors(r.config.GetBool("browser.ignore_cert_errors", false))

	r.launchers[slug] = l
	r.browsers[slug] = newBrowser
	return newBrowser
}

// Destroy 销毁浏览器
func (r *Browser) Destroy(slug string) {
	r.ensureLock(slug)

	r.locks[slug].Lock()
	defer r.locks[slug].Unlock()

	r.mu.Lock()
	defer r.mu.Unlock()

	if browser, exists := r.browsers[slug]; exists {
		browser.MustClose()
		delete(r.browsers, slug)
	}

	if r.config.GetString("browser.mode", "control") == "manage" {
		if l, exists := r.launchers[slug]; exists {
			l.Cleanup()
			delete(r.launchers, slug)
		}
	}
}

// Lock 获取锁
func (r *Browser) Lock(slug string) {
	r.ensureLock(slug)
	r.locks[slug].Lock()
}

// Unlock 释放锁
func (r *Browser) Unlock(slug string) {
	r.ensureLock(slug)
	r.locks[slug].Unlock()
}

// ensureLock 确保每个 slug 都有一个对应的锁
func (r *Browser) ensureLock(slug string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.locks[slug]; !exists {
		r.locks[slug] = &sync.Mutex{}
	}
}
