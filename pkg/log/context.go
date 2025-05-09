package log

import (
	"context"
	"sync"
)

type loggerCtxKey struct{}

var (
	// defaultLogger is the default logger used when no logger is provided.
	defaultLogger Logger = NewSlogLogger(nil)
	m                    = &sync.Mutex{}
)

// Default returns the default [Logger].
func Default() Logger {
	m.Lock()
	defer m.Unlock()
	return defaultLogger
}

func SetDefault(l *Logger) {
	m.Lock()
	defer m.Unlock()
	if l == nil {
		defaultLogger = NewSlogLogger(nil)
	} else {
		defaultLogger = *l
	}
}

// ContextWithLogger returns a new context with the provided logger.
func ContextWithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey{}, logger)
}

// Context returns the logger from the context. If no logger is found, it
func FromContext(ctx context.Context) Logger {
	if logger, ok := ctx.Value(loggerCtxKey{}).(Logger); ok {
		return logger
	}

	return Default()
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
func With(ctx context.Context, fields ...any) context.Context {
	logger := FromContext(ctx).With(fields...)
	return context.WithValue(ctx, loggerCtxKey{}, logger)
}

// Debug logs a message at Debug level.
func Debug(ctx context.Context, msg string, fields ...any) {
	FromContext(ctx).Debug(ctx, msg, fields...)
}

// Error logs a message at Error level.
func Error(ctx context.Context, msg string, fields ...any) {
	FromContext(ctx).Error(ctx, msg, fields...)
}

// Fatal logs a message at Fatal level.
func Fatal(ctx context.Context, msg string, fields ...any) {
	FromContext(ctx).Fatal(ctx, msg, fields...)
}

// Info logs a message at Info level.
func Info(ctx context.Context, msg string, fields ...any) {
	FromContext(ctx).Info(ctx, msg, fields...)
}

// Warn logs a message at Warn level.
func Warn(ctx context.Context, msg string, fields ...any) {
	FromContext(ctx).Warn(ctx, msg, fields...)
}
