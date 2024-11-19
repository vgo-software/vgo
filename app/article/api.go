package article

import (
	"vgo-software/vgo/app/common"
	"vgo-software/vgo/app/model"
	"vgo-software/vgo/internal/global"
)

func RegisterApiRoutes() {
	articleHandler := common.NewCRUDHandler(&model.Article{}, nil)
	global.ApiRouter.GET("/articles", articleHandler.Index)
	global.ApiRouter.GET("/articles/:id", articleHandler.Show)
}
