package api

import (
	"bytes"
	"demo-pod/logger"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
)

const staticFilesPath = "./static"

func serveStatic(basePath string) gin.HandlerFunc {
	fs := static.LocalFile(staticFilesPath, false)
	fileserver := http.FileServer(fs)
	if basePath != "" {
		fileserver = http.StripPrefix(basePath, fileserver)
	}
	return func(c *gin.Context) {
		if staticFileExists(staticFilesPath, basePath, c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

func indexHandler(indexHtml string, basePath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, basePath) {
			c.Header("Content-Type", "text/html")
			c.String(http.StatusOK, indexHtml)
		}
	}
}

func renderIndex(settings *Settings) string {
	tmpl, err := template.ParseFiles("./static/index.html")

	if err != nil {
		logger.Logger.Fatalln("Could not parse index.html as template: ", err)
		return ""
	}

	var renderedIndexBuffer bytes.Buffer
	indexTemplateSettings := IndexTemplateSettings{
		BasePath: settings.BasePath,
	}

	if !strings.HasSuffix(indexTemplateSettings.BasePath, "/") {
		indexTemplateSettings.BasePath += "/"
	}

	err = tmpl.Execute(
		&renderedIndexBuffer,
		&indexTemplateSettings,
	)

	if err != nil {
		logger.Logger.Fatalln("Could not render index.html template: ", err)
		return ""
	}

	return renderedIndexBuffer.String()
}

func staticFileExists(staticFilesPath string, prefix string, filepath string) bool {
	// We want to render the index.html as template
	if strings.HasSuffix(filepath, static.INDEX) {
		return false
	}

	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		name := path.Join(staticFilesPath, p)
		stats, err := os.Stat(name)
		if err != nil {
			return false
		}
		if stats.IsDir() {
			return false
		}
		return true
	}
	return false
}
