package controller

import (
	"github.com/easyship/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var fileServer = http.FileServerFS(service.DefaultOssFileSystem)

func OssFileGetHandler(ctx *gin.Context) {
	// 回传pdf文件
	fileServer.ServeHTTP(ctx.Writer, ctx.Request)
}
