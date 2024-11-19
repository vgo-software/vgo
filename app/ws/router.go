package ws

import (
	"vgo-software/vgo/internal/global"
)

func CollectRoutes() {
	global.Engine.GET("/ws/link", Link)
	global.Engine.POST("/ws/send", Send)
	global.Engine.POST("/ws/send_to_all", SendToAll)
}
