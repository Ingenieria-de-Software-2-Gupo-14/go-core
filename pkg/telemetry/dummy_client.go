package telemetry

import "time"

var DefaultClient Client = &dummyClient{}

type dummyClient struct{}

func (d *dummyClient) Count(name string, value int64, tags []string) {
	// No-op
}

func (d *dummyClient) Incr(name string, tags []string) {
	// No-op
}

func (d *dummyClient) Decr(name string, tags []string) {
	// No-op
}

func (d *dummyClient) Timing(name string, value time.Duration, tags []string) {
	// No-op
}
