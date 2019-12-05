package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"demo-pod/utils"
	"time"
)

var timeFormat = "2006-01-02T15:04:05Z"

func GinLoggerMiddleware(conf gin.LoggerConfig, logger logrus.FieldLogger) gin.HandlerFunc {
	formatter := conf.Formatter
	if formatter == nil {
		formatter = func(param gin.LogFormatterParams) string {
			return ""
		}
	}

	out := conf.Output
	if out == nil {
		out = gin.DefaultWriter
	}

	notlogged := conf.SkipPaths

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			param := gin.LogFormatterParams{
				Request: c.Request,
				Keys:    c.Keys,
			}

			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)

			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path

			entry := logger.WithFields(logrus.Fields{
				"hostname":   hostname,
				"statusCode": param.StatusCode,
				"latency":    param.Latency, // time to process
				"clientIP":   param.ClientIP,
				"originalClientIP": utils.OriginalClientIp(c),
				"method":     c.Request.Method,
				"path":       path,
				"referer":    c.Request.Referer(),
				"dataLength": c.Writer.Size(),
				"userAgent":  c.Request.UserAgent(),
			})

			if len(c.Errors) > 0 {
				entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
			} else {
				msg := formatter(param)
				if param.StatusCode >= 500 {
					entry.Error(msg)
				} else if param.StatusCode >= 400 {
					entry.Warn(msg)
				} else {
					entry.Info(msg)
				}
			}

			fmt.Fprint(out, formatter(param))
		}
	}
}
