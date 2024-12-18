package backend

import (
	"github.com/gin-gonic/gin"
	"time"
	"vgo-software/vgo/app/article/model"
	"vgo-software/vgo/app/common/backend"
	"vgo-software/vgo/internal/global"
	"vgo-software/vgo/internal/pkg/middleware/rate-limiter"
	"vgo-software/vgo/pkg/helper"
	"vgo-software/vgo/pkg/response"
)

// Change 改变状态
func Change(ctx *gin.Context) {
	var notice model.Article
	if err := helper.BindJSON(ctx, &notice); err != nil {
		response.Fail(ctx, "参数错误", err.Error(), nil)
		return
	}
	if err := global.DbCon.Model(&model.Article{}).Where("id = ?", notice.ID).Updates(notice).Error; err != nil {
		response.Fail(ctx, "更新失败", err.Error())
		return
	}
	response.Success(ctx, "成功", nil, nil)
}

func RegisterRoutes() {
	ArticleValidateRules := &backend.ValidationRules{
		Create: map[string]map[string]string{
			"Title": {
				"required": "标题不能为空",
				"min":      "标题长度不能少于2个字符",
			},
		},
		Update: map[string]map[string]string{
			"Title": {
				"required": "标题不能为空",
				"min":      "标题长度不能少于2666个字符",
			},
		},
	}
	articleHandler := backend.NewCRUDHandler(&model.Article{}, ArticleValidateRules)

	global.BackendRouter.GET("/articles", rate_limiter.Limiter(1, time.Second*1), articleHandler.Index)
	global.BackendRouter.POST("/articles", articleHandler.Create)
	global.BackendRouter.PUT("/articles", articleHandler.Update)
	global.BackendRouter.GET("/articles/:id", articleHandler.Show)
	global.BackendRouter.DELETE("/articles", articleHandler.Delete)
	global.BackendRouter.POST("/articles/change", Change)
}
