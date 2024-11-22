package app

import (
	AdminUserRouter "vgo-software/vgo/app/admin-user/router"
	ArticleRouter "vgo-software/vgo/app/article/router"
	CommonRouter "vgo-software/vgo/app/common/router"
	MenuRouter "vgo-software/vgo/app/menu/router"
	NoticeRouter "vgo-software/vgo/app/notice/router"
	RoleRouter "vgo-software/vgo/app/role/router"
	UploadRouter "vgo-software/vgo/app/upload/router"
	wsrouter "vgo-software/vgo/app/ws/router"

	"vgo-software/vgo/internal/global"
	"vgo-software/vgo/internal/pkg/middleware/auth"
	"vgo-software/vgo/internal/pkg/middleware/permission"
)

func InitRouter() {
	global.BackendRouter = global.Engine.Group("/backend")
	global.BackendRouter.Use(auth.AdminAuthMiddleware(), permission.Check(global.Rbac))

	global.ApiRouter = global.Engine.Group("/api/v1")

	ArticleRouter.CollectRoutes()
	NoticeRouter.CollectRoutes()

	AdminUserRouter.CollectRoutes()
	CommonRouter.CollectRoutes()
	wsrouter.CollectRoutes()
	UploadRouter.CollectRoutes()

	MenuRouter.CollectRoutes()
	RoleRouter.CollectRoutes()
}
