package telemetry

import (
	"context"
	"fmt"
	"time"

	"github.com/DataDog/datadog-go/v5/statsd"
)

type datadogClient struct {
	client statsd.ClientInterface
}

// NewDatadogClient creates a new Datadog statsd client with the given address and options.
func NewDatadog(addr string, options ...statsd.Option) (Client, error) {
	c, err := statsd.New(addr, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create datadog statsd client: %w", err)
	}
	return &datadogClient{client: c}, nil
}

func NewDatadogClient(client statsd.ClientInterface) Client {
	return &datadogClient{client: client}
}

func (d *datadogClient) Count(_ctx context.Context, name string, value int64, tags []string) {
	d.client.Count(name, value, tags, 1)
}

func (d *datadogClient) Incr(_ctx context.Context, name string, tags []string) {
	d.client.Incr(name, tags, 1)
}

func (d *datadogClient) Decr(_ctx context.Context, name string, tags []string) {
	d.client.Decr(name, tags, 1)
}

func (d *datadogClient) Timing(_ctx context.Context, name string, value time.Duration, tags []string) {
	d.client.Timing(name, value, tags, 1)
}

var _ Client = (*datadogClient)(nil)
