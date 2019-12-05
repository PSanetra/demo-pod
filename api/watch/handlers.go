package watch

import (
	"github.com/gin-gonic/gin"
)

func AddHandlers(
	router *gin.RouterGroup,
	settings *Settings,
) {
	AddGetWatchWhitelistHandler(router, settings)
	AddGetWatchHandler(router, settings)
}
