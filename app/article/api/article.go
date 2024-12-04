package api

import (
	"vgo-software/vgo/app/article/model"
	"vgo-software/vgo/app/common/backend"
	"vgo-software/vgo/internal/global"
)

func RegisterRoutes() {
	articleHandler := backend.NewCRUDHandler(&model.Article{}, nil)
	global.ApiRouter.GET("/articles", articleHandler.Index)
	global.ApiRouter.GET("/articles/:id", articleHandler.Show)
}
