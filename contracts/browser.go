package contracts

import "github.com/go-rod/rod"

type Browser interface {
	// New 获取浏览器
	New(slug string) *rod.Browser
	// Destroy 销毁浏览器
	Destroy(slug string)
	// Lock 获取锁
	Lock(slug string)
	// Unlock 释放锁
	Unlock(slug string)
}
