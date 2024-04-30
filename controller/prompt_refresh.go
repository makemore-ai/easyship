package controller

import (
	"github.com/easyship/service"
	"github.com/easyship/util"
	"github.com/easyship/util/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RefreshPromptHandler 润色Prompt
func RefreshPromptHandler(ctx *gin.Context) {
	reqPrompt := ctx.Query("prompt")
	log.InfoWithContext(ctx, "reqPrompt:%v", reqPrompt)
	moduleIdStr := ctx.Query("moduleId")
	var moduleId int32 = 0
	if len(moduleIdStr) == 0 {
		moduleId = 0
	}
	streamFunc, err := service.RefreshPromptWithSSE(ctx, reqPrompt, moduleId)
	if err != nil {
		log.ErrorWithContext(ctx, "QwenTextRequestWithSSE error:%v", err)
		// 返回json值
		ctx.JSON(http.StatusOK, util.ErrResp(err))
		return
	}
	intSSEHeader(ctx)
	ctx.Stream(streamFunc)
}

func intSSEHeader(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
}
