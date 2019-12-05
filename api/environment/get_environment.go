package environment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func AddGetEnvHandler(router *gin.RouterGroup) {
	router.GET("/environment", GetIpHandler())
}

func GetIpHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, getEnvironment(os.Environ()))

	}
}

func getEnvironment(data []string) map[string]string {
	items := make(map[string]string)
	for _, item := range data {
		key, val := getKeyVal(item)
		items[key] = val
	}
	return items
}

func getKeyVal(item string) (key, val string) {
	splits := strings.SplitN(item, "=", 2)
	key = splits[0]
	val = splits[1]
	return
}
