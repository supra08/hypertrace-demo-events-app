package backend

import (
	"context"
	"fmt"
	"net/http"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"

	"github.com/hypertrace/demo-events-app/pkg/tracing"
)

// Client is a remote client that implements event.Interface
type Client struct {
	tracer   opentracing.Tracer
	client   *tracing.HTTPClient
	hostPort string
}

// NewClient creates a new event.Client
func NewClient(tracer opentracing.Tracer, hostPort string) *Client {
	return &Client{
		tracer: tracer,
		client: &tracing.HTTPClient{
			Client: &http.Client{Transport: &nethttp.Transport{}},
			Tracer: tracer,
		},
		hostPort: hostPort,
	}
}

// GetEvent implements event.Interface#Get as an RPC
func (c *Client) GetEvent(ctx context.Context, eventID string) (*Event, error) {
	url := fmt.Sprintf("http://"+c.hostPort+"/events?eventID=%s", eventID)
	fmt.Println(url)
	var event Event
	if err := c.client.GetJSON(ctx, "/events", url, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

// GetEvents implements event.Interface#Get as an RPC
func (c *Client) GetEvents(ctx context.Context) ([]*Event, error) {
	url := fmt.Sprintf("http://" + c.hostPort + "/events")
	fmt.Println(url)
	var events []*Event
	if err := c.client.GetJSON(ctx, "/events", url, events); err != nil {
		return nil, err
	}
	return events, nil
}

// GetComments implements comment.Interface#Get as an RPC
func (c *Client) GetComments(ctx context.Context, eventID string) ([]*Comment, error) {
	url := fmt.Sprintf("http://"+c.hostPort+"/comments?eventID=%s", eventID)
	fmt.Println(url)
	var comments []*Comment
	if err := c.client.GetJSON(ctx, "/comments", url, comments); err != nil {
		return nil, err
	}
	return comments, nil
}

// CreateComment implements comment.Interface#Post as an RPC
func (c *Client) CreateComment(ctx context.Context, comment Comment) error {
	url := fmt.Sprintf("http://" + c.hostPort + "/comments")
	fmt.Println(url)
	if err := c.client.PostJSON(ctx, "/comments", url, comment); err != nil {
		return err
	}
	return nil
}

// CreateEvent implements event.Interface#Post as an RPC
func (c *Client) CreateEvent(ctx context.Context, event Event) error {
	url := fmt.Sprintf("http://" + c.hostPort + "/events")
	fmt.Println(url)
	if err := c.client.PostJSON(ctx, "/events", url, event); err != nil {
		return err
	}
	return nil
}
