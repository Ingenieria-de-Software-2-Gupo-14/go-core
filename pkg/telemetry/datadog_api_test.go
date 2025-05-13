package telemetry_test

import (
	"context"
	"testing"
	"time"

	"github.com/Ingenieria-de-Software-2-Gupo-14/go-core/pkg/telemetry"
)

func TestDatadogAPIClient(t *testing.T) {
	// This test won't actually send metrics to Datadog,
	// it just ensures that the client can be created and methods work

	// Create a client with custom options
	client, err := telemetry.NewDatadogAPI(
		telemetry.WithAPIResourceName("test-host"),
		telemetry.WithAPIResourceType("service"),
		telemetry.WithAPIFlushPeriod(5*time.Second),
	)

	if err != nil {
		t.Fatalf("Failed to create Datadog API client: %v", err)
	}

	// Test that the client methods don't panic
	t.Run("count", func(t *testing.T) {
		client.Count(context.Background(), "test.metric.count", 5, []string{"tag:value"})
	})

	t.Run("increment", func(t *testing.T) {
		client.Incr(context.Background(), "test.metric.incr", []string{"tag:value"})
	})

	t.Run("decrement", func(t *testing.T) {
		client.Decr(context.Background(), "test.metric.decr", []string{"tag:value"})
	})

	t.Run("timing", func(t *testing.T) {
		client.Timing(context.Background(), "test.metric.timing", 100*time.Millisecond, []string{"tag:value"})
	})
}
