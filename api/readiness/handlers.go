package readiness

import "github.com/gin-gonic/gin"

func AddHandlers(
	router *gin.RouterGroup,
	settings *Settings,
) {
	AddGetReadinessHandler(router, settings)
	AddPutReadinessHandler(router, settings)
}
