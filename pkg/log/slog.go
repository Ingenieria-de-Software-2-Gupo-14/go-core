package log

import (
	"context"
	"log/slog"
	"os"
)

// SlogLogger implements the Logger interface using slog.
type SlogLogger struct {
	logger *slog.Logger
}

// NewSlogLogger creates a new Logger using the provided slog.Logger.
func NewSlogLogger(logger *slog.Logger) Logger {
	if logger == nil {
		// Default to a JSON handler writing to stdout
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}
	return &SlogLogger{logger: logger}
}

// With creates a new logger with additional context.
func (s *SlogLogger) With(fields ...any) Logger {
	return &SlogLogger{
		logger: s.logger.With(fields...),
	}
}

// Fatal logs a message at Fatal level and terminates the program.
func (s *SlogLogger) Fatal(ctx context.Context, msg string, fields ...any) {
	s.logger.Log(ctx, slog.LevelError, msg, fields...)
	os.Exit(1) // Standard slog doesn't exit, so we do it ourselves
}

// Debug logs a message at Debug level.
func (s *SlogLogger) Debug(ctx context.Context, msg string, fields ...any) {
	s.logger.Log(ctx, slog.LevelDebug, msg, fields...)
}

// Info logs a message at Info level.
func (s *SlogLogger) Info(ctx context.Context, msg string, fields ...any) {
	s.logger.InfoContext(ctx, msg, fields...)
}

// Warn logs a message at Warn level.
func (s *SlogLogger) Warn(ctx context.Context, msg string, fields ...any) {
	s.logger.WarnContext(ctx, msg, fields...)
}

// Error logs a message at Error level.
func (s *SlogLogger) Error(ctx context.Context, msg string, fields ...any) {
	s.logger.ErrorContext(ctx, msg, fields...)
}
