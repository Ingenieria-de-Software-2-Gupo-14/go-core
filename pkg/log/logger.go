package log

import (
	"context"
)

type Logger interface {
	With(fields ...any) Logger
	Fatal(ctx context.Context, msg string, fields ...any)
	Debug(ctx context.Context, msg string, fields ...any)
	Info(ctx context.Context, msg string, fields ...any)
	Warn(ctx context.Context, msg string, fields ...any)
	Error(ctx context.Context, msg string, fields ...any)
}
