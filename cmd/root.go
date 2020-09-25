package cmd

import (
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger         *zap.Logger
	metricsFactory metrics.Factory
	host           string = "0.0.0.0"
	backendPort    int    = 8080
	frontendPort   int    = 3000
	basepath       string = "/"
)

func logError(logger *zap.Logger, err error) error {
	if err != nil {
		logger.Error("Error running command", zap.Error(err))
	}
	return err
}

func initServerConfig() *zap.Logger {
	rand.Seed(int64(time.Now().Nanosecond()))
	logger, _ = zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)
	zapLogger := logger.With(zap.String("service", "events"))
	return zapLogger
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "events-app",
	Short: "Events-app - A tracing demo application",
	Long:  `Events-app - A tracing demo application`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logger.Fatal("Some error occured, aborting...", zap.Error(err))
		os.Exit(-1)
	}
}
