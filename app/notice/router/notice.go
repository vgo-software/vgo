package router

import (
	"vgo-software/vgo/app/notice/api"
	"vgo-software/vgo/app/notice/backend"
)

func CollectRoutes() {
	backend.RegisterNoticeRoutes()
	api.RegisterNoticeApiRoutes()
}
