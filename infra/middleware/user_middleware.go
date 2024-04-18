package middleware

import (
	"github.com/easyship/infra"
	"github.com/easyship/model"
	"github.com/easyship/util/log"
	"github.com/gin-gonic/gin"
	"strings"
)

// 所有mobile的userAgent
var MOBILE_KEYWORDS = []string{"Mobile", "Android", "Silk/", "Kindle",
	"BlackBerry", "Opera Mini", "Opera Mobi"}

// 用于将用户header信息包装到ctx中
func UserMiddleware(ctx *gin.Context) {
	headers := ctx.Request.Header
	userAgent := headers.Get("User-Agent")

	isMobile := false
	for i := 0; i < len(MOBILE_KEYWORDS); i++ {
		if strings.Contains(userAgent, MOBILE_KEYWORDS[i]) {
			isMobile = true
			break
		}
	}
	requestUserInfo := &model.RequestUserInfo{
		IsMobile: isMobile,
	}
	log.InfoWithContext(ctx, "IsMobile:%v", isMobile)
	ctx.Set(infra.USER_INFO, requestUserInfo)
}
