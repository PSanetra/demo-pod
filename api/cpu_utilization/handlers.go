package cpu_utilization

import "github.com/gin-gonic/gin"

func AddHandlers(
	router *gin.RouterGroup,
) {
	AddGetCpuUtilizationHandler(router)
	AddPutCpuStressHandler(router)
}
