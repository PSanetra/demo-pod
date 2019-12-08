package memory_utilization

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"demo-pod/logger"
	"runtime"
	"runtime/debug"
	"sync"
)

func AddPutMemoryBlockHandler(router *gin.RouterGroup) {
	router.PUT("/memory-block", PutMemoryBlockHandler())
}

var memoryBlock []byte
var memoryBlockMutex sync.RWMutex

func PutMemoryBlockHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		var bytes int

		if err := c.ShouldBindJSON(&bytes); err != nil {
			logger.Logger.Infoln("Could not read json: ", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		memoryBlockMutex.Lock()
		defer memoryBlockMutex.Unlock()

		memoryBlock = nil

		runtime.GC()
		debug.FreeOSMemory()

		if bytes > 0 {
			memoryBlock = make([]byte, bytes)

			// there is no memset in go :-/
			for i, _ := range memoryBlock {
				memoryBlock[i] = 1
			}
		}

		c.Status(http.StatusNoContent)

	}
}
