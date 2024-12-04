package router

import (
	"vgo-software/vgo/app/article/api"
	"vgo-software/vgo/app/article/backend"
)

func CollectRoutes() {
	api.RegisterRoutes()
	backend.RegisterRoutes()
}
