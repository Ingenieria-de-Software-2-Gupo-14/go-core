package telemetry

import (
	"context"
	"time"
)

// Count tracks how many times something happened per second.
func Count(ctx context.Context, name string, value int64, tags []string) {
	FromContext(ctx).Count(ctx, name, value, tags)
}

// Decr is just Count of -1.
func Decr(ctx context.Context, name string, tags []string) {
	FromContext(ctx).Decr(ctx, name, tags)
}

// Incr is just Count of 1.
func Incr(ctx context.Context, name string, tags []string) {
	FromContext(ctx).Incr(ctx, name, tags)
}

// Timing sends timing information, it is an alias for TimeInMilliseconds.
func Timing(ctx context.Context, name string, value time.Duration, tags []string) {
	FromContext(ctx).Timing(ctx, name, value, tags)
}
