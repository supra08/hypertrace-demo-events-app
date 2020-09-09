package backend

import (
	"context"
)

// Event contains data about a event.
type Event struct {
	ID   string
	Name string
}

// Interface exposed by the Event service.
type Interface interface {
	Get(ctx context.Context, eventID string) (*Event, error)
}
