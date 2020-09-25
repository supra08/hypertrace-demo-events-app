package backend

import (
	"net/http"

	jexpvar "github.com/uber/jaeger-lib/metrics/expvar"

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
	database *database
}

func NewServer(hostPort string, tracer opentracing.Tracer, logger log.Factory) *Server {
	return &Server{
		hostPort: hostPort,
		tracer:   tracer,
		logger:   logger,
		database: newDatabase(
			tracing.Init("mysql", jexpvar.NewFactory(10), logger),
			logger.With(zap.String("component", "mysql")),
		),
	}
}

func (s *Server) Run() error {
	mux := s.createServeMux()
	return http.ListenAndServe(s.hostPort, mux)
}

func (s *Server) createServeMux() http.Handler {
	mux := tracing.NewServeMux(s.tracer)
	mux.Handle("/api/v1/comments", http.HandlerFunc(s.commentsHandler))
	mux.Handle("/api/v1/events", http.HandlerFunc(s.eventsHandler))
	return mux
}

func enableCors(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (s *Server) commentsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)

	ctx := r.Context()
	s.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	if err := r.ParseForm(); httperr.HandleError(w, err, http.StatusBadRequest) {
		return
	}

	switch r.Method {
	case "GET":
		s.FetchComments(w, r)
	case "POST":
		s.CreateComment(w, r)
	}
}

func (s *Server) eventsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)

	ctx := r.Context()
	s.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	if r.Method == "POST" {
		s.CreateEvent(w, r)
	} else {
		if eventID := r.URL.Query().Get("eventID"); eventID != "" {
			s.FetchEvent(w, r)
		} else {
			s.FetchAllEvents(w, r)
		}
	}
}
