package article

import (
	"ych/vgo/app/common"
	"ych/vgo/app/model"
	"ych/vgo/internal/global"
)

func RegisterApiRoutes() {
	articleHandler := common.NewCRUDHandler(&model.Article{}, nil)
	global.ApiRouter.GET("/articles", articleHandler.Index)
	global.ApiRouter.GET("/articles/:id", articleHandler.Show)
}
