package backend

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/hypertrace/demo-events-app/pkg/httperr"
	"go.uber.org/zap"
)

// CreatePostResponse contains the json response for POST requests
type CreatePostResponse struct {
	Status string `json:"status"`
}

// CreateComment controller handles commment creation
func (s *Server) CreateComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var comment Comment
	if err = json.Unmarshal(b, &comment); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	commentUUID, err := uuid.NewUUID()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	commentID := commentUUID.String()

	comment.ID = commentID

	err1 := s.database.CreateComment(ctx, &comment)
	if httperr.HandleError(w, err, http.StatusInternalServerError) {
		s.logger.For(ctx).Error("request failed", zap.Error(err1))
		return
	}

	responseBody := CreatePostResponse{"success"}
	jsonResponse, err := json.Marshal(responseBody)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// FetchComments controller handles retrieval of comments
func (s *Server) FetchComments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	s.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))
	if err := r.ParseForm(); httperr.HandleError(w, err, http.StatusBadRequest) {
		s.logger.For(ctx).Error("bad request", zap.Error(err))
		return
	}

	eventID := r.URL.Query().Get("event_id")
	if eventID == "" {
		http.Error(w, "Missing required 'eventID' parameter", http.StatusBadRequest)
		return
	}

	response, err := s.database.GetComments(ctx, eventID)
	if httperr.HandleError(w, err, http.StatusInternalServerError) {
		s.logger.For(ctx).Error("request failed", zap.Error(err))
		return
	}

	data, err := json.Marshal(response)
	if httperr.HandleError(w, err, http.StatusInternalServerError) {
		s.logger.For(ctx).Error("cannot marshal response", zap.Error(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// CreateEvent controller handles creation of events
func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var event Event
	if err = json.Unmarshal(b, &event); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	eventUUID, err := uuid.NewUUID()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	eventID := eventUUID.String()

	event.ID = eventID

	err1 := s.database.CreateEvent(ctx, &event)
	if httperr.HandleError(w, err, http.StatusInternalServerError) {
		s.logger.For(ctx).Error("request failed", zap.Error(err1))
		return
	}

	response := CreatePostResponse{"success"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// FetchEvent controller handles the retrieval of a specific event
func (s *Server) FetchEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	s.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))
	if err := r.ParseForm(); httperr.HandleError(w, err, http.StatusBadRequest) {
		s.logger.For(ctx).Error("bad request", zap.Error(err))
		return
	}

	eventID := r.URL.Query().Get("eventID")
	if eventID == "" {
		http.Error(w, "Missing required 'eventID' parameter", http.StatusBadRequest)
		return
	}

	response, err := s.database.GetEvent(ctx, eventID)
	if httperr.HandleError(w, err, http.StatusInternalServerError) {
		s.logger.For(ctx).Error("request failed", zap.Error(err))
		return
	}

	data, err := json.Marshal(response)
	if httperr.HandleError(w, err, http.StatusInternalServerError) {
		s.logger.For(ctx).Error("cannot marshal response", zap.Error(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// FetchAllEvents handles the retrieval of all events
func (s *Server) FetchAllEvents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	s.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))
	if err := r.ParseForm(); httperr.HandleError(w, err, http.StatusBadRequest) {
		s.logger.For(ctx).Error("bad request", zap.Error(err))
		return
	}

	response, err := s.database.GetEvents(ctx)
	if httperr.HandleError(w, err, http.StatusInternalServerError) {
		s.logger.For(ctx).Error("request failed", zap.Error(err))
		return
	}

	data, err := json.Marshal(response)
	if httperr.HandleError(w, err, http.StatusInternalServerError) {
		s.logger.For(ctx).Error("cannot marshal response", zap.Error(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
