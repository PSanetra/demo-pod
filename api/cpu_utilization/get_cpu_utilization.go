package cpu_utilization

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"net/http"
	"demo-pod/logger"
	"time"
)

func AddGetCpuUtilizationHandler(router *gin.RouterGroup) {
	router.GET("/cpu-utilization", GetCpuUtilizationHandler())
}

func GetCpuUtilizationHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		count, err := cpu.Counts(true)

		if err != nil {
			logger.Logger.Errorln("Could not read cpu count:", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		cpuResponse := make([]int, count)

		cpuTimes, err := cpu.Percent(1 * time.Second, true)

		if err != nil {
			logger.Logger.Errorln("Could not read cpu times:", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		for i, c := range cpuTimes {
			cpuResponse[i] = int(c)
		}

		c.JSON(http.StatusOK, cpuResponse)
	}
}
