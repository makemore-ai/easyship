package main

import (
	"context"
	"github.com/easyship/controller"
	"github.com/easyship/infra"
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
	router.Static("/jquery", "./html/jquery")
	router.GET(
		"/", controller.IndexHandler)
	router.POST(
		"/searchPrompt", controller.PromptSearchHandle)
	err = router.Run()
	if err != nil {
		print(err)
	}
}
