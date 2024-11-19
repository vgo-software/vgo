package common

import (
	"vgo-software/vgo/internal/global"
)

func CollectRoutes() {
	global.Engine.GET("/test", Test)
}
