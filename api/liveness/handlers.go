package liveness

import "github.com/gin-gonic/gin"

func AddHandlers(
	router *gin.RouterGroup,
	settings *Settings,
) {
	AddGetAliveHandler(router, settings)
	AddGetLivenessHandler(router, settings)
	AddPutLivenessHandler(router, settings)
}
