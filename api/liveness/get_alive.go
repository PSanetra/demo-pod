package liveness

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddGetAliveHandler(router *gin.RouterGroup, settings *Settings) {
	router.GET("/alive", GetAliveHandler(settings))
}

func GetAliveHandler(settings *Settings) func(c *gin.Context) {
	return func(c *gin.Context) {

		alive, _ := settings.alive.Load().(bool)

		if alive {
			c.Status(http.StatusNoContent)
		} else {
			c.Status(http.StatusInternalServerError)
		}

	}
}
