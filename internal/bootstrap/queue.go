package bootstrap

import (
	"sync"
	"vgo-software/vgo/internal/global"
	"vgo-software/vgo/internal/pkg/queue"
)

// queueLock 读写锁
var queueLock sync.RWMutex

// InitQueue 初始化队列
func InitQueue() {
	queueConf := global.Config.QueueConf
	if queueConf.Enable != 1 {
		return
	}
	queueLock.RLock()
	defer queueLock.RUnlock()
	queue.InitQueue()
}
