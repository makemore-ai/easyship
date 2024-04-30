package main

import (
	"context"
	"github.com/easyship/controller"
	"github.com/easyship/infra"
	"github.com/easyship/infra/middleware"
	"github.com/easyship/util/log"
	"github.com/gin-gonic/gin"
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
	router.GET(
		"/", middleware.UserMiddleware, controller.IndexHandler)
	router.POST(
		"/searchPrompt", controller.PromptSearchHandle)
	router.GET(
		"/refreshPrompt", controller.RefreshPromptHandler)
	err = router.Run()
	if err != nil {
		print(err)
	}
}
