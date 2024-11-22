package router

import (
	"vgo-software/vgo/app/admin-user/backend"
)

func CollectRoutes() {
	backend.RegisterAdminUserRoutes()
}
