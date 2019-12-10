package api

import (
	"demo-pod/api/cpu_utilization"
	"demo-pod/api/environment"
	"demo-pod/api/ip"
	"demo-pod/api/liveness"
	"demo-pod/api/memory_utilization"
	"demo-pod/api/notes"
	"demo-pod/api/readiness"
	"demo-pod/api/watch"
	"demo-pod/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(settings *Settings) *gin.Engine {

	engine := gin.New()
	engine.ForwardedByClientIP = false

	engine.Use(
		logger.GinLoggerMiddleware(gin.LoggerConfig{}, logger.Logger),
		gin.Recovery(),
		serveStatic(settings.BasePath),
	)

	router := engine.Group(settings.BasePath)

	for _, origin := range settings.CorsOrigins {

		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{origin},
			AllowMethods:     []string{"PUT", "GET"},
			AllowHeaders:     []string{"content-type"},
			AllowCredentials: true,
		}))

	}

	addApiHandlers(router, settings)

	indexHtml := renderIndex(settings)

	router.GET("/index.html", indexHandler(indexHtml, settings.BasePath))
	engine.NoRoute(indexHandler(indexHtml, settings.BasePath))

	return engine

}

func addApiHandlers(
	router *gin.RouterGroup,
	settings *Settings,
) {
	apiRouterGroup := router.Group("/api")

	environment.AddGetEnvHandler(apiRouterGroup)
	cpu_utilization.AddHandlers(apiRouterGroup)
	ip.AddGetIpHandler(apiRouterGroup)
	liveness.AddHandlers(apiRouterGroup, settings.LivenessSettings)
	memory_utilization.AddHandlers(apiRouterGroup)
	notes.AddHandlers(apiRouterGroup, &settings.NotesSettings)
	readiness.AddHandlers(apiRouterGroup, settings.ReadinessSettings)
	watch.AddHandlers(apiRouterGroup, &settings.WatchSettings)

}
