package upload

import (
	"vgo-software/vgo/internal/global"
)

func CollectRoutes() {
	global.BackendRouter.POST("/upload/local/img", ImgUpload)
	global.BackendRouter.POST("/upload/local/video", VideoUpload)
	global.BackendRouter.POST("/upload/oss", RunOssUpload)
	global.BackendRouter.POST("/upload/cos", RunCos)
	global.BackendRouter.POST("/upload/auto", Run)
}
