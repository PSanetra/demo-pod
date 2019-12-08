package memory_utilization

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddGetMemoryBlockHandler(router *gin.RouterGroup) {
	router.GET("/memory-block", GetMemoryBlockHandler())
}

func GetMemoryBlockHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		memoryBlockMutex.RLock()
		defer memoryBlockMutex.RUnlock()

		size := len(memoryBlock)

		c.JSON(http.StatusOK, &size)

	}
}
