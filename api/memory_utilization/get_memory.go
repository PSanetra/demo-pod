package memory_utilization

import (
	"github.com/cloudfoundry/bytefmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/mem"
	"net/http"
	"demo-pod/logger"
)

func AddGetMemoryHandler(router *gin.RouterGroup) {
	router.GET("/memory", GetMemoryHandler())
}

func GetMemoryHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		memory, err := mem.VirtualMemory()

		if err != nil {
			logger.Logger.Errorln("Could not read virtual memory:", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		memoryResponse := Memory{
			InUse: bytefmt.ByteSize(memory.Used),
			Available: bytefmt.ByteSize(memory.Available),
		}

		c.JSON(http.StatusOK, &memoryResponse)
	}
}
