package telemetry

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

// datadogAPIClient is a client that sends metrics to Datadog using the Datadog API
type datadogAPIClient struct {
	api          *datadogV2.MetricsApi
	flushPeriod  time.Duration
	metrics      []datadogV2.MetricSeries
	mutex        sync.Mutex
	ctx          context.Context
	resourceName string
	resourceType string
}

// DatadogAPIOption is a function that configures a datadogAPIClient
type DatadogAPIOption func(*datadogAPIClient)

// WithAPIResourceName sets the resource name for the datadogAPIClient
func WithAPIResourceName(name string) DatadogAPIOption {
	return func(c *datadogAPIClient) {
		c.resourceName = name
	}
}

// WithAPIResourceType sets the resource type for the datadogAPIClient
func WithAPIResourceType(resourceType string) DatadogAPIOption {
	return func(c *datadogAPIClient) {
		c.resourceType = resourceType
	}
}

// WithAPIFlushPeriod sets the flush period for the datadogAPIClient
func WithAPIFlushPeriod(period time.Duration) DatadogAPIOption {
	return func(c *datadogAPIClient) {
		c.flushPeriod = period
	}
}

// NewDatadogAPI creates a new client that sends metrics using the Datadog API
func NewDatadogAPI(opts ...DatadogAPIOption) (Client, error) {
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)

	client := &datadogAPIClient{
		api:          datadogV2.NewMetricsApi(apiClient),
		metrics:      make([]datadogV2.MetricSeries, 0),
		flushPeriod:  1 * time.Second,
		resourceName: "default",
		resourceType: "host",
		ctx:          datadog.NewDefaultContext(context.Background()),
	}

	for _, opt := range opts {
		opt(client)
	}

	go client.periodicFlush()

	return client, nil
}

// Count tracks how many times something happened per second.
func (c *datadogAPIClient) Count(ctx context.Context, name string, value int64, tags ...string) {
	metric := datadogV2.MetricSeries{
		Metric: name,
		Type:   datadogV2.METRICINTAKETYPE_COUNT.Ptr(),
		Points: []datadogV2.MetricPoint{
			{
				Timestamp: datadog.PtrInt64(time.Now().Unix()),
				Value:     datadog.PtrFloat64(float64(value)),
			},
		},
		Tags: tags,
		Resources: []datadogV2.MetricResource{
			{
				Name: datadog.PtrString(c.resourceName),
				Type: datadog.PtrString(c.resourceType),
			},
		},
	}
	c.submitMetric(metric)
}

// Incr is just Count of 1.
func (c *datadogAPIClient) Incr(ctx context.Context, name string, tags ...string) {
	c.Count(ctx, name, 1, tags...)
}

// Decr is just Count of -1.
func (c *datadogAPIClient) Decr(ctx context.Context, name string, tags ...string) {
	c.Count(ctx, name, -1, tags...)
}

// Timing sends timing information, it is an alias for TimeInMilliseconds.
func (c *datadogAPIClient) Timing(ctx context.Context, name string, value time.Duration, tags ...string) {
	metric := datadogV2.MetricSeries{
		Metric: name,
		Type:   datadogV2.METRICINTAKETYPE_UNSPECIFIED.Ptr(),
		Points: []datadogV2.MetricPoint{
			{
				Timestamp: datadog.PtrInt64(time.Now().Unix()),
				Value:     datadog.PtrFloat64(float64(value.Milliseconds())),
			},
		},
		Tags: tags,
		Resources: []datadogV2.MetricResource{
			{
				Name: datadog.PtrString(c.resourceName),
				Type: datadog.PtrString(c.resourceType),
			},
		},
	}

	c.submitMetric(metric)
}

// submitMetric submits a metric to the Datadog API
func (c *datadogAPIClient) submitMetric(metric datadogV2.MetricSeries) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.metrics = append(c.metrics, metric)
}

// Flush sends all pending metrics to Datadog
func (c *datadogAPIClient) Flush() error {
	c.mutex.Lock()
	if len(c.metrics) == 0 {
		c.mutex.Unlock()
		return nil
	}

	metrics := c.metrics
	c.metrics = make([]datadogV2.MetricSeries, 0)
	c.mutex.Unlock()

	payload := datadogV2.MetricPayload{
		Series: metrics,
	}

	_, r, err := c.api.SubmitMetrics(c.ctx, payload, *datadogV2.NewSubmitMetricsOptionalParameters())
	if err != nil {
		return fmt.Errorf("failed to submit metrics to Datadog: %w, HTTP response: %v", err, r)
	}

	return nil
}

// periodicFlush periodically flushes metrics to Datadog
func (c *datadogAPIClient) periodicFlush() {
	ticker := time.NewTicker(c.flushPeriod)
	defer ticker.Stop()

	for range ticker.C {
		//TODO: handle err
		c.Flush()
	}
}

// Ensure datadogAPIClient implements Client interface
var _ Client = (*datadogAPIClient)(nil)
