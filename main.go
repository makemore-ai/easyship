package main

import (
	"github.com/easyship/controller"
	"github.com/easyship/infra"
	"github.com/easyship/infra/middleware"
	"github.com/easyship/util/log"
	"github.com/gin-gonic/gin"
	"path"

	"context"
)

const (
	GET  = "get"
	POST = ""
)

func main() {
	err := infra.Init(context.Background())
	if err != nil {
		log.ErrorWithContext(context.Background(), "Init infra error:%v", err)
		panic(err)
	}
	router := gin.Default()
	router.LoadHTMLGlob("./html/*.html")
	//静态文件资源
	router.Static("/static", "./static")
	router.Static("/js", "./html/js")
	router.Static("/css", "./html/css")
	router.Static("/jquery", "./html/jquery")
	router.Static("/page", "./html")

	// 为了yd兼容的一块逻辑，之后用nginx把这块逻辑转出去 弄台单机就可以
	router.Static("/yd-pdf", "./static")

	urlPattern := path.Join("/yd-info", "/*filepath")
	router.GET(urlPattern, controller.OssFileGetHandler)
	router.HEAD(urlPattern, controller.OssFileGetHandler)

	router.GET(
		"/", middleware.UserMiddleware, controller.IndexHandler)
	router.POST(
		"/searchPrompt", controller.PromptSearchHandle)
	router.GET(
		"/refreshPrompt", controller.RefreshPromptHandler)
	router.Use(gin.Recovery(), gin.Logger())
	err = router.Run()
	if err != nil {
		print(err)
	}
}
