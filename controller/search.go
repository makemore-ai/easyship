package controller

import (
	"github.com/easyship/service"
	"github.com/easyship/util"
	"github.com/easyship/util/log"
	"github.com/gin-gonic/gin"

	"net/http"
)

func PromptSearchHandle(ctx *gin.Context) {
	searchText := ctx.PostForm("searchText")
	promptDto, err := service.SearchPrompt(ctx, &searchText)
	if err != nil {
		log.ErrorWithContext(ctx, "SearchHandle query:%v error:%v", searchText, err)
		ctx.JSON(http.StatusOK, util.ErrResp(err))
		return
	}
	ctx.JSON(http.StatusOK, util.SuccessResp(promptDto))
}
