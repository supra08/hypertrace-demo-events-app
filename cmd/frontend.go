package cmd

import (
	"net"
	"strconv"

	"github.com/hypertrace/demo-events-app/pkg/log"
	"github.com/hypertrace/demo-events-app/pkg/tracing"
	frontend "github.com/hypertrace/demo-events-app/services/frontend"
	jexpvar "github.com/uber/jaeger-lib/metrics/expvar"
	"go.uber.org/zap"
)

var (
	options frontend.ConfigOptions
)

func setupFrontendServer(zapLogger *zap.Logger, host string, port int, basepath string) error {
	options.FrontendHostPort = net.JoinHostPort("0.0.0.0", strconv.Itoa(frontendPort))
	options.Basepath = basepath
	metricsFactory = jexpvar.NewFactory(10)
	logger := log.NewFactory(zapLogger)
	server := frontend.NewServer(
		options,
		tracing.Init("frontend", metricsFactory, logger),
		logger,
	)

	return logError(zapLogger, server.Run())
}
