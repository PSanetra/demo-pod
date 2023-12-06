package http_headers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddGetHttpHeadersHandler(router *gin.RouterGroup) {
	router.GET("/http/headers", GetHttpHeadersHandler())
}

func GetHttpHeadersHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, c.Request.Header)

	}
}
