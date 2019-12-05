package cpu_utilization

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"demo-pod/logger"
	"runtime"
	"sync"
)

func AddPutCpuStressHandler(router *gin.RouterGroup) {
	router.PUT("/cpu-stress", PutCpuStressHandler())
}

var stressCancelFuncs []context.CancelFunc
var stressCancelFuncsMutex sync.Mutex

func PutCpuStressHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		var threads int

		if err := c.ShouldBindJSON(&threads); err != nil {
			logger.Logger.Infoln("Could not read json: ", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		stressCancelFuncsMutex.Lock()
		defer stressCancelFuncsMutex.Unlock()

		cancelExistingStressRoutines()

		createStressRoutines(threads)

		c.Status(http.StatusNoContent)

	}
}

func createStressRoutines(threads int) {
	stressCancelFuncs = make([]context.CancelFunc, threads)

	for i, _ := range stressCancelFuncs {
		var ctx context.Context
		ctx, stressCancelFuncs[i] = context.WithCancel(context.Background())
		go stress(ctx)
	}
}

func cancelExistingStressRoutines() {
	if len(stressCancelFuncs) != 0 {
		for _, cancel := range stressCancelFuncs {
			cancel()
		}
	}
}

func stress(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		runtime.Gosched()
	}

}
