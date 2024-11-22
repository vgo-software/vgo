package router

import (
	"vgo-software/vgo/app/role/backend"
)

func CollectRoutes() {
	backend.RegisterRoleRoutes()
}
