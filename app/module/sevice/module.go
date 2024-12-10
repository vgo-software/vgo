package sevice

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"vgo-software/vgo/app/module/model"
	"vgo-software/vgo/internal/global"
	"vgo-software/vgo/pkg/response"
)

func CreateModule(ctx *gin.Context, name string) {
	modulePath := filepath.Join("app", name)

	// 创建module目录
	if err := createDir(modulePath, ctx, "module目录"); err != nil {
		return
	}

	// 目录列表和文件名
	directories := []string{"api", "backend", "model", "router"}
	fileName := name + ".go"

	// 遍历目录，创建子目录和文件
	for _, dir := range directories {
		modulePathDir := filepath.Join(modulePath, dir)

		// 创建子目录
		if err := createDir(modulePathDir, ctx, dir); err != nil {
			return
		}

		// 生成文件路径并创建文件
		if err := createFile(filepath.Join(modulePathDir, fileName), ctx, name, dir); err != nil {
			return
		}
	}
}

// 创建文件
func createFile(filePath string, ctx *gin.Context, name string, dir string) error {
	if _, err := os.Stat(filePath); err == nil {
		return nil // 文件已存在，不再创建
	} else if !os.IsNotExist(err) {
		logErrorAndRespond(ctx, fmt.Sprintf("检查文件失败: %s", filePath), err)
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		logErrorAndRespond(ctx, fmt.Sprintf("创建文件失败: %s", filePath), err)
		return err
	}
	defer file.Close()

	content := generateContent(name, dir)
	return writeFile(filePath, content, ctx, dir+"文件")
}

// 生成内容
func generateContent(name string, dir string) string {
	nameUpper := fmt.Sprintf("%s%s", strings.ToUpper(string(name[0])), name[1:]) // 首字母大写
	return generateDirContent(name, dir, nameUpper)
}

func generateDirContent(name, dir, nameUpper string) string {
	switch dir {
	case "router":
		return fmt.Sprintf(`
package router

import (
	"vgo-software/vgo/app/%s/api"
	"vgo-software/vgo/app/%s/backend"
)

func CollectRoutes() {
	api.RegisterRoutes()
	backend.RegisterRoutes()
}
`, name, name)

	case "model":
		return fmt.Sprintf(`
package model

import "vgo-software/vgo/app/common/model"

type %s struct {
	model.BaseModel
}
`, nameUpper)

	case "backend":
		return fmt.Sprintf(`
package backend

import (
	"vgo-software/vgo/app/common/backend"
	"vgo-software/vgo/app/%s/model"
	"vgo-software/vgo/internal/global"
)

func RegisterRoutes() {
	handler := backend.NewCRUDHandler(&model.%s{}, nil)
	global.BackendRouter.GET("/%ss", handler.Index)
	global.BackendRouter.POST("/%ss", handler.Create)
	global.BackendRouter.PUT("/%ss", handler.Update)
	global.BackendRouter.GET("/%ss/:id", handler.Show)
	global.BackendRouter.DELETE("/%ss", handler.Delete)
}
`, name, nameUpper, name, name, name, name, name)
	case "api":
		return fmt.Sprintf(`
package api

import (
	"vgo-software/vgo/app/common/backend"
	"vgo-software/vgo/app/%s/model"
	"vgo-software/vgo/internal/global"
)

func RegisterRoutes() {
	handler := backend.NewCRUDHandler(&model.%s{}, nil)
	global.ApiRouter.GET("/%ss", handler.Index)
}
`, name, nameUpper, name)
	default:
		return ""
	}
}

// 写入文件
func writeFile(filePath string, content string, ctx *gin.Context, fileName string) error {
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		logErrorAndRespond(ctx, fmt.Sprintf("写入%s失败", fileName), err)
		return err
	}
	return nil
}

// 创建目录
func createDir(path string, ctx *gin.Context, dirName string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		logErrorAndRespond(ctx, fmt.Sprintf("创建%s失败", dirName), err)
		return err
	}
	return nil
}

// 记录错误并返回响应
func logErrorAndRespond(ctx *gin.Context, message string, err error) {
	fmt.Println("具体错误信息:", err)
	response.Fail(ctx, message, err, nil)
}

// LoadModule 加载模块到数据库
func LoadModule() {
	modulePath := filepath.Join("app", "/")
	moduleList := make([]string, 0)
	entries, err := ioutil.ReadDir(modulePath)
	if err != nil {
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			moduleList = append(moduleList, entry.Name())
		}
	}
	global.DbCon.Exec("DELETE FROM modules")
	for _, module := range moduleList {
		global.DbCon.Create(&model.Module{
			Name: module,
		})
		global.Logger.Info(fmt.Sprintf("module: %s", module))
	}
}
