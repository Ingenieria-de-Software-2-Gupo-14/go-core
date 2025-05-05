package telemetry

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-go/v5/statsd"
)

type datadogClient struct {
	client *statsd.Client
}

// NewDatadogClient creates a new Datadog statsd client with the given address and options.
func NewDatadogClient(addr string, options ...statsd.Option) (Client, error) {
	c, err := statsd.New(addr, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create datadog statsd client: %w", err)
	}
	return &datadogClient{client: c}, nil
}

func (d *datadogClient) Count(name string, value int64, tags []string) {
	d.client.Count(name, value, tags, 1)
}

func (d *datadogClient) Incr(name string, tags []string) {
	d.client.Incr(name, tags, 1)
}

func (d *datadogClient) Decr(name string, tags []string) {
	d.client.Decr(name, tags, 1)
}

func (d *datadogClient) Timing(name string, value time.Duration, tags []string) {
	d.client.Timing(name, value, tags, 1)
}

var _ Client = (*datadogClient)(nil)
