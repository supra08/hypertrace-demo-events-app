package backend

import (
	"context"
)

// Event contains data about a event.
type Event struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Comment contains data about a comment.
type Comment struct {
	ID        string `json:"id"`
	Commenter string `json:"commenter"`
	Content   string `json:"content"`
	EventID   string `json:"event_id"`
}

// Interface exposed by the Event service.
type Interface interface {
	GetEvent(ctx context.Context, eventID string) (*Event, error)
	GetEvents(ctx context.Context) ([]*Event, error)
	CreateEvent(ctx context.Context, comment Comment) error
	GetComments(ctx context.Context, eventID string) ([]*Comment, error)
	CreateComment(ctx context.Context, comment Comment) error
}
