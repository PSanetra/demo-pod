package ip

import (
	"demo-pod/logger"
	"demo-pod/utils"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"strings"
)

func AddGetIpHandler(router *gin.RouterGroup) {
	router.GET("/ip", GetIpHandler())
}

func GetIpHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		clientIp, _, _ := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr))

		ip := Ip{
			PodIpList:  getPodIpList(),
			ClientIp:   clientIp,
			OriginalIp: utils.OriginalClientIp(c),
		}

		c.JSON(http.StatusOK, &ip)

	}
}

func getPodIpList() []string {
	var ipList []string
	ifaces, err := net.Interfaces()

	if err != nil {
		logger.Logger.Warningln("Could not read interfaces:", err)
		return []string{}
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()

		if err != nil {
			logger.Logger.Warningln("Could not read addresses:", err)
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			default:
				continue
			}

			ipList = append(ipList, ip.String())
		}
	}

	return ipList
}
