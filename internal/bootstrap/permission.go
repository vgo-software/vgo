package bootstrap

import (
	"sync"
	"vgo-software/vgo/internal/global"
	"vgo-software/vgo/internal/pkg/middleware/permission"
)

// permissionLock 读写锁
var permissionLock sync.RWMutex

// InitRbac 初始化RBAC权限
func InitRbac() {
	permissionLock.RLock()
	defer permissionLock.RUnlock()
	global.Rbac = permission.SetupCasbin()
	if global.Rbac == nil {
		panic("RBAC权限初始化失败")
	}
}
