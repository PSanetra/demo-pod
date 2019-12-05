package notes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"demo-pod/logger"
)

func AddGetNotesHandler(router *gin.RouterGroup, settings *Settings) {
	router.GET("/notes", GetNotesHandler(settings))
}

func GetNotesHandler(settings *Settings) func(c *gin.Context) {
	return func (c *gin.Context) {

		data, err := ioutil.ReadFile(settings.StatePath)

		if err == os.ErrNotExist {
			logger.Logger.Errorln("Could not find file: ", err)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if err == os.ErrPermission {
			logger.Logger.Errorln("Could not access file: ", err)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		if err != nil {
			logger.Logger.Errorln("Could not read file: ", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var notes Notes

		err = json.Unmarshal(data, &notes)

		if err != nil {
			logger.Logger.Errorln("Could not unmarshal data: ", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, &notes)

	}
}
