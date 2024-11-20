package upload

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"io"
	"vgo-software/vgo/internal/global"
	"vgo-software/vgo/pkg/response"
)

func Run(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.Fail(ctx, "获取文件失败", err.Error(), nil)
		return
	}

	uploadType := ctx.PostForm("type")
	if uploadType == "" {
		response.Fail(ctx, "上传类型标识不能为空", nil)
		return
	}
	defTypes := map[string]struct {
		size  int64
		mimes []string
	}{
		"any": {
			size:  50 * 1024 * 1024, // 50MB
			mimes: []string{"image/jpeg", "image/png", "image/jpg", "video/mp4"},
		},
		"img": {
			size:  10 * 1024 * 1024, // 10MB
			mimes: []string{"image/jpeg", "image/png", "image/jpg"},
		},
		"video": {
			size:  10 * 1024 * 1024, // 10MB
			mimes: []string{"video/mp4", "video/avi"},
		},
	}
	if _, exists := defTypes[uploadType]; !exists {
		response.Fail(ctx, "上传类型标识不支持", nil)
		return
	}
	defType := defTypes[uploadType]
	if file.Size > defType.size {
		response.Fail(ctx, fmt.Sprintf("文件大小不能超过 %d MB", defType.size/(1024*1024)), nil)
		return
	}
	src, err := file.Open()
	if err != nil {
		response.Fail(ctx, "获取文件失败", err.Error())
		return
	}
	defer src.Close()
	fileContent, err := io.ReadAll(src)
	if err != nil {
		response.Fail(ctx, "获取文件失败", err.Error())
		return
	}
	// 获取文件mimetype
	mime := mimetype.Detect(fileContent)
	fileType := mime.String()
	validMime := false
	for _, mimeType := range defType.mimes {
		if mimeType == fileType {
			validMime = true
			break
		}
	}
	if !validMime {
		response.Fail(ctx, "不支持的文件类型", nil)
		return
	}

	sysUploadType := global.Config.App.UploadType
	switch sysUploadType {
	case "local":
		respUrl, err := putObjectByLocal("img", file)
		if err != nil {
			return
		}
		imgDomain := global.Config.App.FileDomain
		response.Success(ctx, "上传成功", gin.H{
			"fileUrl": imgDomain + respUrl,
		}, nil)
		break
	case "oss":
		respUrl, err := putObjectByOss(file)
		if err != nil {
			return
		}
		response.Success(ctx, "上传成功", gin.H{
			"fileUrl": respUrl,
		}, nil)
		break
	case "cos":
		respUrl, err := putObjectByCos(file)
		if err != nil {
			return
		}
		response.Success(ctx, "上传成功", gin.H{
			"fileUrl": respUrl,
		}, nil)
		break
	}
}
