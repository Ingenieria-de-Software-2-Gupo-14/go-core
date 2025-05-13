package telemetry_test

import (
	"context"
	"testing"
	"time"

	"github.com/Ingenieria-de-Software-2-Gupo-14/go-core/pkg/telemetry"

	"github.com/DataDog/datadog-go/v5/statsd"
)

func TestMetricsSettingClient(t *testing.T) {
	// Create a mock client
	mockClient := telemetry.NewDatadogClient(&statsd.NoOpClient{})

	t.Run("increment", func(t *testing.T) {
		// Create a context with the mock client
		ctx := telemetry.Context(context.Background(), mockClient)

		// Call the Incr function
		telemetry.Incr(ctx, "test_metric", "tag1", "tag2")
	})

	t.Run("decrement", func(t *testing.T) {
		// Create a context with the mock client
		ctx := telemetry.Context(context.Background(), mockClient)

		// Call the Decr function
		telemetry.Decr(ctx, "test_metric", "tag1", "tag2")
	})

	t.Run("count", func(t *testing.T) {
		// Create a context with the mock client
		ctx := telemetry.Context(context.Background(), mockClient)

		// Call the Count function
		telemetry.Count(ctx, "test_metric", 5, "tag1", "tag2")
	})

	t.Run("timing", func(t *testing.T) {
		// Create a context with the mock client
		ctx := telemetry.Context(context.Background(), mockClient)

		// Call the Timing function
		telemetry.Timing(ctx, "test_metric", time.Millisecond, "tag1", "tag2")
	})
}

func TestMetricsWithDefaultClient(t *testing.T) {
	t.Run("increment", func(t *testing.T) {
		// Call the Incr function
		telemetry.Incr(context.Background(), "test_metric", "tag1", "tag2")
	})

	t.Run("decrement", func(t *testing.T) {
		// Call the Decr function
		telemetry.Decr(context.Background(), "test_metric", "tag1", "tag2")
	})

	t.Run("count", func(t *testing.T) {
		// Call the Count function
		telemetry.Count(context.Background(), "test_metric", 5, "tag1", "tag2")
	})

	t.Run("timing", func(t *testing.T) {
		// Call the Timing function
		telemetry.Timing(context.Background(), "test_metric", time.Millisecond, "tag1", "tag2")
	})
}
