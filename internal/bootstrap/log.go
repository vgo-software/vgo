package bootstrap

import (
	"sync"
	"vgo-software/vgo/internal/global"
	"vgo-software/vgo/internal/log"
)

// logLock 读写锁
var logLock sync.RWMutex

// InitLogger 初始化日志
func InitLogger() {
	logLock.RLock()
	defer logLock.RUnlock()
	global.Logger = log.InitLogDriver()
	if global.Logger == nil {
		panic("日志驱动初始化失败")
	}
}
