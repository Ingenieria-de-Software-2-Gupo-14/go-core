package telemetry

import (
	"github.com/gin-gonic/gin"
)

func AddTracer(tracer Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := Context(c.Request.Context(), tracer)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
