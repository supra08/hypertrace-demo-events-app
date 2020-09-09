package main

import (
	"math/rand"
	"time"

	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger         *zap.Logger
	metricsFactory metrics.Factory
	host           string = "0.0.0.0"
	backendPort    int    = 8080
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

func main() {
	zapLogger := initServerConfig()
	backend := setupBackendServer(zapLogger, host, backendPort)

	go logError(zapLogger, backend.Run())
}
