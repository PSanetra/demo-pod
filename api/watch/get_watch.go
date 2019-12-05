package watch

import (
	"crypto/sha256"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"demo-pod/logger"
)

type WatchResponse struct {
	Content string `json:"content"`
}

func AddGetWatchHandler(router *gin.RouterGroup, settings *Settings) {
	router.GET("/watch/:key", GetWatchHandler(settings))
}

func GetWatchHandler(settings *Settings) func(c *gin.Context) {
	return func(c *gin.Context) {

		key := c.Param("key")
		path, ok := settings.FileWhitelist[key]

		if !ok {
			logger.Logger.Info("Path for key not found: ", key)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		mode := c.DefaultQuery("mode", "default")

		switch mode {
		case "default":
			returnFile(c, path)
			break
		case "sse":
			watchFile(c, key, path)
			break
		}

	}
}

func watchFile(c *gin.Context, key string, path string) {
	logger.Logger.Debugln("watching file ", path)
	var contentHash [sha256.Size]byte

	// ensure nginx doesn't buffer this response but sends the chunks right away
	c.Header("X-Accel-Buffering", "no")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add(path)
	if err != nil {
		logger.Logger.Errorln("Can not watch ", path, ": ", err)
		return
	}

	c.Stream(func(w io.Writer) bool {
		contentHash = sendFileEvent(c, "initial", path, contentHash)
		return false
	})

	c.Stream(func(w io.Writer) bool {
		select {
		case <-c.Request.Context().Done():
			logger.Logger.Debugln("Request context is done: ", c.ClientIP())
			return false
		case event, ok := <-watcher.Events:
			if !ok {
				return true
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				logger.Logger.Debugln("File modified:", event.String())
				contentHash = sendFileEvent(c, "file-modified", path, contentHash)
			}
			break
		case err, ok := <-watcher.Errors:
			if !ok {
				return true
			}
			logger.Logger.Errorln("watchFile error:", err)
			break
		}

		return true
	})
}

func sendFileEvent(c *gin.Context, eventType string, path string, previousHash [sha256.Size]byte) [sha256.Size]byte {
	var hash [sha256.Size]byte
	data, err := ioutil.ReadFile(path)

	if err != nil {
		logger.Logger.Infoln("Could not read file ", path, ": ", err)
		return hash
	}

	hash = sha256.Sum256(data)

	if hash == previousHash {
		logger.Logger.Debugf("Skipping duplicate event for hash %x\n", hash)
		return hash
	}

	c.SSEvent(eventType, &WatchResponse{
		Content: string(data),
	})

	return hash
}

func returnFile(c *gin.Context, path string) {
	data, err := ioutil.ReadFile(path)

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

	c.Data(http.StatusOK, "text/plain", data)
}
