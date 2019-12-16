package main

import (
	"demo-pod/api"
	"demo-pod/api/liveness"
	"demo-pod/api/notes"
	"demo-pod/api/readiness"
	"demo-pod/api/watch"
	"demo-pod/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"time"
)

var settings = api.Settings{
	LivenessSettings:  liveness.NewSettings(),
	NotesSettings:     notes.Settings{},
	ReadinessSettings: readiness.NewSettings(),
	WatchSettings:     watch.Settings{},
}

var rootCmd = &cobra.Command{
	Use:   "demo-pod [<ip>:<port>]",
	Short: "DON'T RUN THIS EVER ON PRODUCTION!!!",
	Long: `
DON'T RUN THIS EVER ON PRODUCTION!!!

demo-pod can be used in Kubernetes workshops to demonstrate different pod properties. 

demo-pod binds to 0.0.0.0:8080 by default if run without TLS and 0.0.0.0:8443 if run with TLS.
`,
	Run: func(cmd *cobra.Command, args []string) {

		processLogLevelFlag(cmd)

		delayStartup()

		ginEngine := api.Setup(&settings)

		address := ""

		if len(args) > 0 {
			address = args[0]
		}

		run(ginEngine, address)

	},
}

func run(engine *gin.Engine, address string) {
	var err error

	useTls := settings.TlsCertPath != "" && settings.TlsKeyPath != ""

	if !useTls {
		if address == "" {
			address = ":8080"
		}

		err = engine.Run(address)
	} else {
		if address == "" {
			address = ":8443"
		}

		err = engine.RunTLS(address, settings.TlsCertPath, settings.TlsKeyPath)
	}

	if err != nil {
		logger.Logger.Fatalln("Gin error: ", err)
	}
}

func delayStartup() {
	if settings.StartupDelay > 0 {
		logger.Logger.Info("Delaying startup by ", settings.StartupDelay.String())
		time.Sleep(settings.StartupDelay)
		logger.Logger.Info("Starting")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Logger.Fatalln(err)
	}
}

func main() {

	err := rootCmd.Execute()

	if err != nil {
		logger.Logger.Fatalln(err)
	}

}

func init() {
	rootCmd.PersistentFlags().StringVar(&settings.TlsKeyPath, "tls-key", "", "path to tls key file")
	rootCmd.PersistentFlags().StringVar(&settings.TlsCertPath, "tls-cert", "", "path to tls certificate file")
	rootCmd.PersistentFlags().String("log-level", logger.DEFAULT_LOG_LEVEL.String(), "panic | fatal | error | warn | info | debug | trace")
	rootCmd.PersistentFlags().StringVar(&settings.NotesSettings.StatePath, "notes-state-file-path", "./notes.json", "")
	rootCmd.PersistentFlags().DurationVar(&settings.StartupDelay, "startup-delay", 0, "specifies a delay on startup (e.g. '10s')")
	rootCmd.PersistentFlags().StringSliceVar(&settings.CorsOrigins, "cors-origin", []string{}, "defines an allowed origin")
	rootCmd.PersistentFlags().StringVar(&settings.BasePath, "base-path", "/", "defines a base path prefix for all http endpoints")
	rootCmd.PersistentFlags().StringToStringVar(&settings.WatchSettings.FileWhitelist, "watch", map[string]string{}, "Whitelist for files retrievable via /watch/{key}. Format of one option is key=/path/to/file")
}

func processLogLevelFlag(cmd *cobra.Command) {
	logLevel := cmd.Flag("log-level").Value.String()
	logger.SetLevel(logLevel)

	if logLevel == "debug" || logLevel == "trace" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
