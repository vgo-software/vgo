package backend

import (
	"github.com/gin-gonic/gin"
	"time"
	"vgo-software/vgo/app/common/backend"
	"vgo-software/vgo/app/module/model"
	"vgo-software/vgo/app/module/sevice"
	"vgo-software/vgo/internal/global"
	"vgo-software/vgo/internal/pkg/middleware/rate-limiter"
	"vgo-software/vgo/pkg/helper"
	"vgo-software/vgo/pkg/response"
)

// Create 新增
func Create(ctx *gin.Context) {
	var m model.Module
	if err := helper.BindJSON(ctx, &m); err != nil {
		response.Fail(ctx, "参数错误", err.Error(), nil)
		return
	}
	// 验证规则
	rules := map[string]map[string]string{
		"Name": {
			"required": "Name不能为空",
		},
	}
	if res, err := helper.Validate(m, rules); !res {
		response.Fail(ctx, err, nil)
		return
	}
	sevice.CreateModule(ctx, m.Name)
	sevice.LoadModule()
	response.Success(ctx, "成功", m, nil)
}

func RegisterRoutes() {
	ModuleValidateRules := &backend.ValidationRules{
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
	handler := backend.NewCRUDHandler(&model.Module{}, ModuleValidateRules)

	global.BackendRouter.GET("/modules", rate_limiter.Limiter(1, time.Second*1), handler.Index)
	global.BackendRouter.POST("/modules", Create)
	global.BackendRouter.PUT("/modules", handler.Update)
	global.BackendRouter.GET("/modules/:id", handler.Show)
}
