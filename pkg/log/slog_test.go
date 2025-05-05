package log

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"
)

func TestSlogLogger(t *testing.T) {
	// Configurar un buffer para capturar la salida
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slogLogger := slog.New(handler)
	logger := NewSlogLogger(slogLogger)

	// Probar el método Info
	ctx := context.Background()
	logger.Info(ctx, "test message", "key", "value")

	// Verificar la salida
	var logEntry map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
		t.Fatalf("Error al analizar la salida del log: %v", err)
	}

	if logEntry["msg"] != "test message" {
		t.Errorf("Mensaje esperado: 'test message', obtenido: %v", logEntry["msg"])
	}
	if logEntry["key"] != "value" {
		t.Errorf("Valor esperado para 'key': 'value', obtenido: %v", logEntry["key"])
	}
	if logEntry["level"] != "INFO" {
		t.Errorf("Nivel esperado: 'INFO', obtenido: %v", logEntry["level"])
	}
}

func TestSlogLoggerWith(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slogLogger := slog.New(handler)
	logger := NewSlogLogger(slogLogger)

	// Probar el método With
	withLogger := logger.With("context_key", "context_value")
	ctx := context.Background()
	withLogger.Info(ctx, "test message", "key", "value")

	var logEntry map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
		t.Fatalf("Error al analizar la salida del log: %v", err)
	}

	if logEntry["context_key"] != "context_value" {
		t.Errorf("Valor esperado para 'context_key': 'context_value', obtenido: %v", logEntry["context_key"])
	}
	if logEntry["key"] != "value" {
		t.Errorf("Valor esperado para 'key': 'value', obtenido: %v", logEntry["key"])
	}
}

func TestSlogLoggerLevels(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slogLogger := slog.New(handler)
	logger := NewSlogLogger(slogLogger)
	ctx := context.Background()

	tests := []struct {
		level   string
		logFunc func(context.Context, string, ...any)
		msg     string
	}{
		{"DEBUG", logger.Debug, "debug message"},
		{"INFO", logger.Info, "info message"},
		{"WARN", logger.Warn, "warn message"},
		{"ERROR", logger.Error, "error message"},
	}

	for _, tt := range tests {
		t.Run(tt.level, func(t *testing.T) {
			buf.Reset()
			tt.logFunc(ctx, tt.msg, "test_key", "test_value")

			var logEntry map[string]interface{}
			if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
				t.Fatalf("Error al analizar la salida del log: %v", err)
			}

			if !strings.EqualFold(logEntry["level"].(string), tt.level) {
				t.Errorf("Nivel esperado: '%s', obtenido: %v", tt.level, logEntry["level"])
			}
			if logEntry["msg"] != tt.msg {
				t.Errorf("Mensaje esperado: '%s', obtenido: %v", tt.msg, logEntry["msg"])
			}
			if logEntry["test_key"] != "test_value" {
				t.Errorf("Valor esperado para 'test_key': 'test_value', obtenido: %v", logEntry["test_key"])
			}
		})
	}
}

func TestSlogLoggerNilCreation(t *testing.T) {
	// Verificar que se crea un logger por defecto cuando se proporciona nil
	logger := NewSlogLogger(nil)
	if logger == nil {
		t.Fatal("NewSlogLogger(nil) retornó nil")
	}
}
