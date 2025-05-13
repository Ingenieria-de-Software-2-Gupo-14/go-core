package telemetry

import (
	"context"
	"time"
)

var DefaultClient Client = &dummyClient{}

type dummyClient struct{}

func (d *dummyClient) Count(ctx context.Context, name string, value int64, tags ...string) {
	// No-op
}

func (d *dummyClient) Incr(ctx context.Context, name string, tags ...string) {
	// No-op
}

func (d *dummyClient) Decr(ctx context.Context, name string, tags ...string) {
	// No-op
}

func (d *dummyClient) Timing(ctx context.Context, name string, value time.Duration, tags ...string) {
	// No-op
}
