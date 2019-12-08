package liveness

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddGetLivenessHandler(router *gin.RouterGroup, settings *Settings) {
	router.GET("/liveness", GetLivenessHandler(settings))
}

func GetLivenessHandler(settings *Settings) func(c *gin.Context) {
	return func(c *gin.Context) {

		alive, _ := settings.alive.Load().(bool)

		c.JSON(http.StatusOK, &Liveness{Alive: alive})

	}
}
