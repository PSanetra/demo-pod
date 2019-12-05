package ip

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"demo-pod/utils"
	"strings"
)

func AddGetIpHandler(router *gin.RouterGroup) {
	router.GET("/ip", GetIpHandler())
}

func GetIpHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		clientIp, _, _ := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr))

		ip := Ip{
			ClientIp: clientIp,
			OriginalIp: utils.OriginalClientIp(c),
		}

		c.JSON(http.StatusOK, &ip)

	}
}
