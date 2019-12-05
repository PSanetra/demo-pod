package readiness

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"demo-pod/logger"
)

func AddPutReadinessHandler(router *gin.RouterGroup, settings *Settings) {
	router.PUT("/readiness", PutReadinessHandler(settings))
}

func PutReadinessHandler(settings *Settings) func(c *gin.Context) {
	return func(c *gin.Context) {

		var readiness Readiness

		if err := c.ShouldBindJSON(&readiness); err != nil {
			logger.Logger.Infoln("Could not read json: ", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		readyAfter := readiness.ReadyAfter.UTC()

		settings.readyAfter.Store(&readyAfter)

		c.Status(http.StatusNoContent)

	}
}
