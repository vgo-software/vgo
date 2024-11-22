package router

import (
	"vgo-software/vgo/app/upload/api/ali-oss"
	"vgo-software/vgo/app/upload/api/auto"
	"vgo-software/vgo/app/upload/api/local"
	"vgo-software/vgo/app/upload/api/tencent-cos"
	"vgo-software/vgo/internal/global"
)

func CollectRoutes() {
	global.BackendRouter.POST("/upload/local/img", local.ImgUpload)
	global.BackendRouter.POST("/upload/local/video", local.VideoUpload)
	global.BackendRouter.POST("/upload/oss", ali_oss.Run)
	global.BackendRouter.POST("/upload/cos", tencent_cos.Run)
	global.BackendRouter.POST("/upload/auto", auto.Run)
}
