package readiness

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AddGetReadinessHandler(router *gin.RouterGroup, settings *Settings) {
	router.GET("/readiness", GetReadinessHandler(settings))
}

func GetReadinessHandler(settings *Settings) func(c *gin.Context) {
	return func(c *gin.Context) {

		readyAfter, _ := settings.readyAfter.Load().(*time.Time)

		c.JSON(http.StatusOK, &Readiness{ReadyAfter: *readyAfter})

	}
}
