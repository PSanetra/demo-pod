package readiness

import "github.com/gin-gonic/gin"

func AddHandlers(
	router *gin.RouterGroup,
	settings *Settings,
) {
	AddGetReadyHandler(router, settings)
	AddGetReadinessHandler(router, settings)
	AddPutReadinessHandler(router, settings)
}
