package backend

import (
	"context"
	"errors"

	tags "github.com/opentracing/opentracing-go/ext"

	"github.com/hypertrace/demo-events-app/pkg/delay"
	"github.com/hypertrace/demo-events-app/pkg/log"
	"github.com/hypertrace/demo-events-app/pkg/tracing"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type database struct {
	tracer   opentracing.Tracer
	logger   log.Factory
	events   map[string]*Event
	comments map[string]*Comment
	lock     *tracing.Mutex
}

func newDatabase(tracer opentracing.Tracer, logger log.Factory) *database {
	events, _ := GetEventsFromConfig()
	comments, _ := GetCommentsFromConfig()
	return &database{
		tracer: tracer,
		logger: logger,
		lock: &tracing.Mutex{
			SessionBaggageKey: "request",
		},
		events:   events,
		comments: comments,
	}
}

func (d *database) GetEvent(ctx context.Context, eventID string) (*Event, error) {
	d.logger.For(ctx).Info("Loading event", zap.String("event_id", eventID))

	// simulate opentracing instrumentation of an SQL query
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := d.tracer.StartSpan("SQL SELECT", opentracing.ChildOf(span.Context()))
		tags.SpanKindRPCClient.Set(span)
		tags.PeerService.Set(span, "mysql")
		// #nosec
		span.SetTag("sql.query", "SELECT * FROM event WHERE event_id="+eventID)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	if !MySQLMutexDisabled {
		// simulate misconfigured connection pool that only gives one connection at a time
		d.lock.Lock(ctx)
		defer d.lock.Unlock()
	}

	// simulate RPC delay
	delay.Sleep(MySQLGetDelay, MySQLGetDelayStdDev)

	if event, ok := d.events[eventID]; ok {
		return event, nil
	}
	return nil, errors.New("invalid event ID")
}

func (d *database) GetEvents(ctx context.Context) ([]*Event, error) {
	// simulate opentracing instrumentation of an SQL query
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := d.tracer.StartSpan("SQL SELECT", opentracing.ChildOf(span.Context()))
		tags.SpanKindRPCClient.Set(span)
		tags.PeerService.Set(span, "mysql")
		// #nosec
		span.SetTag("sql.query", "SELECT * FROM event")
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	var events []*Event
	for _, event := range d.events {
		events = append(events, event)
	}

	return events, nil
}

func (d *database) CreateEvent(ctx context.Context, event *Event) error {
	// simulate opentracing instrumentation of an SQL query
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := d.tracer.StartSpan("SQL SELECT", opentracing.ChildOf(span.Context()))
		tags.SpanKindRPCClient.Set(span)
		tags.PeerService.Set(span, "mysql")
		// #nosec
		span.SetTag("sql.query", "SELECT * FROM event WHERE event_id=")
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	if !MySQLMutexDisabled {
		// simulate misconfigured connection pool that only gives one connection at a time
		d.lock.Lock(ctx)
		defer d.lock.Unlock()
	}

	// simulate RPC delay
	delay.Sleep(MySQLGetDelay, MySQLGetDelayStdDev)

	d.events[event.ID] = event
	return nil
}

func (d *database) GetComments(ctx context.Context, eventID string) ([]*Comment, error) {
	// simulate opentracing instrumentation of an SQL query
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := d.tracer.StartSpan("SQL SELECT", opentracing.ChildOf(span.Context()))
		tags.SpanKindRPCClient.Set(span)
		tags.PeerService.Set(span, "mysql")
		// #nosec
		span.SetTag("sql.query", "SELECT * FROM comment WHERE event_id="+eventID)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	var comments []*Comment
	for _, comment := range d.comments {
		if comment.EventID == eventID {
			comments = append(comments, comment)
		}
	}

	return comments, nil
}

func (d *database) CreateComment(ctx context.Context, comment *Comment) error {
	// simulate opentracing instrumentation of an SQL query
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := d.tracer.StartSpan("SQL SELECT", opentracing.ChildOf(span.Context()))
		tags.SpanKindRPCClient.Set(span)
		tags.PeerService.Set(span, "mysql")
		// #nosec
		span.SetTag("sql.query", "SELECT * FROM comment WHERE event_id=")
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	if !MySQLMutexDisabled {
		// simulate misconfigured connection pool that only gives one connection at a time
		d.lock.Lock(ctx)
		defer d.lock.Unlock()
	}

	// simulate RPC delay
	delay.Sleep(MySQLGetDelay, MySQLGetDelayStdDev)

	d.comments[comment.ID] = comment
	return nil
}
