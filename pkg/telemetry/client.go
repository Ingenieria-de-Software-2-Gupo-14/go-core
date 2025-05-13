package telemetry

import (
	"context"
	"time"
)

type Client interface {
	Count(ctx context.Context, name string, value int64, tags ...string)
	Incr(ctx context.Context, name string, tags ...string)
	Decr(ctx context.Context, name string, tags ...string)
	Timing(ctx context.Context, name string, value time.Duration, tags ...string)
}
