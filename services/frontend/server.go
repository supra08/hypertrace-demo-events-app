package frontend

import (
	"net/http"
	"path"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/hypertrace/demo-events-app/pkg/log"
	"github.com/hypertrace/demo-events-app/pkg/tracing"
)

// Server implements jaeger-demo-frontend service
type Server struct {
	hostPort string
	tracer   opentracing.Tracer
	logger   log.Factory
	assetFS  http.FileSystem
	basepath string
}

// ConfigOptions used to make sure service clients
// can find correct server ports
type ConfigOptions struct {
	FrontendHostPort string
	Basepath         string
}

// NewServer creates a new frontend.Server
func NewServer(options ConfigOptions, tracer opentracing.Tracer, logger log.Factory) *Server {
	assetFS := FS(false)
	return &Server{
		hostPort: options.FrontendHostPort,
		tracer:   tracer,
		logger:   logger,
		assetFS:  assetFS,
		basepath: options.Basepath,
	}
}

// Run starts the frontend server
func (s *Server) Run() error {
	mux := s.createServeMux()
	s.logger.Bg().Info("Starting", zap.String("address", "http://"+path.Join(s.hostPort, s.basepath)))
	return http.ListenAndServe(s.hostPort, mux)
}

func (s *Server) createServeMux() http.Handler {
	mux := tracing.NewServeMux(s.tracer)
	p := path.Join("/", s.basepath)
	mux.Handle(p, http.StripPrefix(p, http.FileServer(s.assetFS)))
	return mux
}
