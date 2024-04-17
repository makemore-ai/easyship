package controller

import (
	"github.com/easyship/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(ctx *gin.Context) {
	recommendPromptList := config.GetRecommendPromptList(ctx)
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"recommend_prompt_list": recommendPromptList,
	})
}
