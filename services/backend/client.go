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

// Get implements event.Interface#Get as an RPC
func (c *Client) Get(ctx context.Context, eventID string) (*Event, error) {
	url := fmt.Sprintf("http://" + c.hostPort + "/events")
	fmt.Println(url)
	var event Event
	if err := c.client.GetJSON(ctx, "/events", url, &event); err != nil {
		return nil, err
	}
	return &event, nil
}
