package cpu_utilization

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddGetCpuStressHandler(router *gin.RouterGroup) {
	router.GET("/cpu-stress", GetCpuStressHandler())
}

func GetCpuStressHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		stressCount := getStressCount()

		c.JSON(http.StatusOK, &stressCount)
	}
}
