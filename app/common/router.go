package common

import (
	"ych/vgo/internal/global"
)

func CollectRoutes() {
	global.Engine.GET("/test", Test)
}
