package liveness

import "github.com/gin-gonic/gin"

func AddHandlers(
	router *gin.RouterGroup,
	settings *Settings,
) {
	AddGetLivenessHandler(router, settings)
	AddPutLivenessHandler(router, settings)
}
