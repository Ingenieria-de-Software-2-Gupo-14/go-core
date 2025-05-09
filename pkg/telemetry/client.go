package telemetry

import "time"

type Client interface {
	Count(name string, value int64, tags []string)
	Incr(name string, tags []string)
	Decr(name string, tags []string)
	Timing(name string, value time.Duration, tags []string)
}
