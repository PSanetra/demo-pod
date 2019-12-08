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
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Setup(settings *Settings) *gin.Engine {

	engine := gin.New()
	engine.ForwardedByClientIP = false

	engine.Use(
		logger.GinLoggerMiddleware(gin.LoggerConfig{}, logger.Logger),
		gin.Recovery(),
		static.Serve("/", static.LocalFile("./static", true)),
	)

	for _, origin := range settings.CorsOrigins {

		engine.Use(cors.New(cors.Config{
			AllowOrigins:     []string{origin},
			AllowMethods:     []string{"PUT", "GET"},
			AllowHeaders:     []string{"content-type"},
			AllowCredentials: true,
		}))

	}

	addApiHandlers(engine, settings)

	engine.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})

	return engine

}

func addApiHandlers(
	engine *gin.Engine,
	settings *Settings,
) {
	apiRouterGroup := engine.Group("/api")

	environment.AddGetEnvHandler(apiRouterGroup)
	cpu_utilization.AddHandlers(apiRouterGroup)
	ip.AddGetIpHandler(apiRouterGroup)
	liveness.AddHandlers(apiRouterGroup, settings.LivenessSettings)
	memory_utilization.AddHandlers(apiRouterGroup)
	notes.AddHandlers(apiRouterGroup, &settings.NotesSettings)
	readiness.AddHandlers(apiRouterGroup, settings.ReadinessSettings)
	watch.AddHandlers(apiRouterGroup, &settings.WatchSettings)

}
