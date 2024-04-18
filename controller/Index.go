package controller

import (
	"github.com/easyship/config"
	"github.com/easyship/infra"
	"github.com/easyship/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(ctx *gin.Context) {
	recommendPromptList := config.GetRecommendPromptList(ctx)
	requestUserInfoModel, _ := ctx.Get(infra.USER_INFO)
	requestUserInfo, _ := requestUserInfoModel.(*model.RequestUserInfo)
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"recommend_prompt_list": recommendPromptList,
		"isMobile":              requestUserInfo.IsMobile,
	})
}
