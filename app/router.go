package app

import (
	"ych/vgo/app/admin-user"
	"ych/vgo/app/article"
	"ych/vgo/app/common"
	"ych/vgo/app/menu"
	"ych/vgo/app/notice"
	"ych/vgo/app/role"
	"ych/vgo/app/upload"
	"ych/vgo/app/ws"
	"ych/vgo/internal/global"
	"ych/vgo/internal/pkg/middleware/auth"
	"ych/vgo/internal/pkg/middleware/permission"
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
