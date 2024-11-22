package router

import (
	"vgo-software/vgo/app/common/api"
	"vgo-software/vgo/internal/global"
)

func CollectRoutes() {
	global.Engine.GET("/test", api.Test)
}
