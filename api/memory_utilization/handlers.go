package memory_utilization

import "github.com/gin-gonic/gin"

func AddHandlers(
	router *gin.RouterGroup,
) {
	AddGetMemoryHandler(router)
	AddPutMemoryUsageHandler(router)
}
