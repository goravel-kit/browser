package browser

import (
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Browser struct {
	browsers map[string]*rod.Browser
	mu       sync.Mutex
	locks    map[string]*sync.Mutex
}

func NewBrowser() *Browser {
	return &Browser{
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

	l := launcher.New().
		Headless(false).
		Devtools(true)

	newBrowser := rod.New().
		ControlURL(l.MustLaunch()).
		Trace(true).
		SlowMotion(2 * time.Second).
		MustConnect().MustIgnoreCertErrors(true)

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
