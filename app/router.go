package app

import (
	"vgo-software/vgo/app/admin-user"
	"vgo-software/vgo/app/article"
	"vgo-software/vgo/app/common"
	"vgo-software/vgo/app/menu"
	"vgo-software/vgo/app/notice"
	"vgo-software/vgo/app/role"
	"vgo-software/vgo/app/upload"
	"vgo-software/vgo/app/ws"
	"vgo-software/vgo/internal/global"
	"vgo-software/vgo/internal/pkg/middleware/auth"
	"vgo-software/vgo/internal/pkg/middleware/permission"
)

func InitRouter() {
	global.BackendRouter = global.Engine.Group("/backend")
	global.BackendRouter.Use(auth.AdminAuthMiddleware(), permission.Check(global.Rbac))
	global.ApiRouter = global.Engine.Group("/api/v1")

	admin_user.CollectRoutes()
	article.CollectRoutes()
	common.CollectRoutes()
	menu.CollectRoutes()
	notice.CollectRoutes()
	role.CollectRoutes()
	upload.CollectRoutes()
	ws.CollectRoutes()
}
