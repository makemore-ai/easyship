package controller

import (
	"github.com/easyship/config"
	"github.com/easyship/infra/constant"
	"github.com/easyship/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(ctx *gin.Context) {
	recommendPromptList := config.GetRecommendPromptList(ctx)
	requestUserInfoModel, _ := ctx.Get(constant.USER_INFO)
	requestUserInfo, _ := requestUserInfoModel.(*model.RequestUserInfo)
	ctx.HTML(http.StatusOK, "prompt.html", gin.H{
		"recommend_prompt_list": recommendPromptList,
		"isMobile":              requestUserInfo.IsMobile,
	})
}
