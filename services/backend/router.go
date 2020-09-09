package backend

import (
	"net/http"

	"github.com/hypertrace/demo-events-app/pkg/httperr"
	"github.com/hypertrace/demo-events-app/pkg/log"
	"github.com/hypertrace/demo-events-app/pkg/tracing"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type Server struct {
	hostPort string
	tracer   opentracing.Tracer
	logger   log.Factory
}

func NewServer(hostPort string, tracer opentracing.Tracer, logger log.Factory) *Server {
	return &Server{
		hostPort: hostPort,
		tracer:   tracer,
		logger:   logger,
	}
}

func (s *Server) Run() error {
	mux := s.createServeMux()
	return http.ListenAndServe(s.hostPort, mux)
}

func (s *Server) createServeMux() http.Handler {
	mux := tracing.NewServeMux(s.tracer)
	mux.Handle("/api/v1/events", http.HandlerFunc(s.eventsHandler))
	return mux
}

func (s *Server) eventsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	s.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	if err := r.ParseForm(); httperr.HandleError(w, err, http.StatusBadRequest) {
		return
	}

	switch r.Method {
	case "GET":
		FetchEvent(w, r)
	case "POST":
		CreateEvent(w, r)
	case "PUT":
		UpdateEvent(w, r)
	case "DELETE":
		DeleteEvent(w, r)
	}
}
