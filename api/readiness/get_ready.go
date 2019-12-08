package readiness

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"demo-pod/logger"
	"time"
)

func AddGetReadyHandler(router *gin.RouterGroup, settings *Settings) {
	router.GET("/ready", GetReadyHandler(settings))
}

func GetReadyHandler(settings *Settings) func(c *gin.Context) {
	return func(c *gin.Context) {

		readyAfter, _ := settings.readyAfter.Load().(*time.Time)

		now := time.Now().UTC()

		if readyAfter.After(now) {
			logger.Logger.Debugln(readyAfter, ">", now)
			c.Status(http.StatusInternalServerError)
		} else {
			logger.Logger.Debugln(readyAfter, "<=", now)
			c.Status(http.StatusNoContent)
		}

	}
}
