package bootstrap

import (
	"sync"
	"vgo-software/vgo/internal/database"
	"vgo-software/vgo/internal/global"
)

// dbLock 读写锁
var dbLock sync.RWMutex

// InitDB 获取数据连接
func InitDB() {
	dbLock.RLock()
	defer dbLock.RUnlock()
	dbConf := global.Config.DbConf
	if dbConf.Driver == "mysql" {
		global.DbCon = database.ConnectMysql()
	}
	if global.DbCon == nil {
		panic("数据库连接失败")
	}
}
