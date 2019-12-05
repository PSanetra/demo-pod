package liveness

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"demo-pod/logger"
)

func AddPutLivenessHandler(router *gin.RouterGroup, settings *Settings) {
	router.PUT("/liveness", PutLivenessHandler(settings))
}

func PutLivenessHandler(settings *Settings) func(c *gin.Context) {
	return func(c *gin.Context) {

		var liveness Liveness

		if err := c.ShouldBindJSON(&liveness); err != nil {
			logger.Logger.Infoln("Could not read json: ", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		settings.alive.Store(liveness.Alive)

		c.Status(http.StatusNoContent)

	}
}
