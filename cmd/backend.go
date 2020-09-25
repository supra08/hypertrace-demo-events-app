package cmd

import (
	"net"
	"strconv"

	"github.com/hypertrace/demo-events-app/pkg/log"
	"github.com/hypertrace/demo-events-app/pkg/tracing"
	backend "github.com/hypertrace/demo-events-app/services/backend"
	jexpvar "github.com/uber/jaeger-lib/metrics/expvar"
	"go.uber.org/zap"
)

func setupBackendServer(zapLogger *zap.Logger, host string, port int) error {
	metricsFactory = jexpvar.NewFactory(10)
	logger := log.NewFactory(zapLogger)
	server := backend.NewServer(
		net.JoinHostPort(host, strconv.Itoa(port)),
		tracing.Init("events", metricsFactory, logger),
		logger,
	)

	return logError(zapLogger, server.Run())
}
