package watch

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddGetWatchWhitelistHandler(router *gin.RouterGroup, settings *Settings) {
	router.GET("/watch-whitelist", GetWatchWhitelistHandler(settings))
}

func GetWatchWhitelistHandler(settings *Settings) func(c *gin.Context) {
	return func (c *gin.Context) {

		keys := make([]string, 0, len(settings.FileWhitelist))

		for k, _ := range settings.FileWhitelist {
			keys = append(keys, k)
		}

		c.JSON(http.StatusOK, keys)

	}
}
