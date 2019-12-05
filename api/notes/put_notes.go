package notes

import (
	"demo-pod/logger"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func AddPutNotesHandler(router *gin.RouterGroup, settings *Settings) {
	router.PUT("/notes", PutNotesHandler(settings))
}

func PutNotesHandler(settings *Settings) func(c *gin.Context) {
	return func(c *gin.Context) {

		var notes Notes

		if err := c.ShouldBindJSON(&notes); err != nil {
			logger.Logger.Infoln("Could not read json: ", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		state, err := json.Marshal(notes)

		if err != nil {
			logger.Logger.Errorln("Could not marshal state: ", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		dir := path.Dir(settings.StatePath)

		err = os.MkdirAll(dir, 0700)

		if err != nil {
			logger.Logger.Errorln("Could not create directory: ", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		err = ioutil.WriteFile(settings.StatePath, state, 0600)

		if err != nil {
			logger.Logger.Errorln("Could not write state: ", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusNoContent)

	}
}
