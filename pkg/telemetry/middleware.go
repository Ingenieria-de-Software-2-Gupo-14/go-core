package telemetry

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func MetricsMiddleware(metricsClient Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := Context(c.Request.Context(), metricsClient)
		c.Request = c.Request.WithContext(ctx)

		startTime := time.Now()
		c.Next()

		tags := []string{
			"method:" + c.Request.Method,
			"path:" + c.FullPath(),
			"status:" + strconv.Itoa(c.Writer.Status()),
		}

		metricsClient.Incr(ctx, "traffic.request", tags...)
		metricsClient.Timing(ctx, "traffic.response.time", time.Since(startTime), tags...)
	}
}
