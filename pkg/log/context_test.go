package log

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"testing"
)

func TestContextWithLogger(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slogLogger := slog.New(handler)
	logger := NewSlogLogger(slogLogger)

	// Crear contexto con logger y recuperarlo
	ctx := context.Background()
	ctxWithLogger := ContextWithLogger(ctx, logger)

	retrievedLogger := FromContext(ctxWithLogger)
	if retrievedLogger == nil {
		t.Fatal("FromContext retornó nil")
	}

	// Verificar que el logger recuperado funciona
	retrievedLogger.Info(ctx, "test message", "key", "value")

	var logEntry map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
		t.Fatalf("Error al analizar la salida del log: %v", err)
	}

	if logEntry["msg"] != "test message" {
		t.Errorf("Mensaje esperado: 'test message', obtenido: %v", logEntry["msg"])
	}
}

func TestDefaultLogger(t *testing.T) {
	// Probar que FromContext retorna el logger por defecto cuando no hay logger en el contexto
	ctx := context.Background()
	logger := FromContext(ctx)

	if logger == nil {
		t.Fatal("Logger por defecto es nil")
	}

	// Probar la función Default
	defaultLogger := Default()
	if defaultLogger == nil {
		t.Fatal("Default() retornó nil")
	}
}

func TestSetDefault(t *testing.T) {
	// Guardar el logger original
	originalLogger := Default()

	// Crear un nuevo logger
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{})
	slogLogger := slog.New(handler)
	newLogger := NewSlogLogger(slogLogger)

	// Establecer como predeterminado y verificar
	SetDefault(&newLogger)

	ctx := context.Background()
	Info(ctx, "test message", "key", "value")

	var logEntry map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
		t.Fatalf("Error al analizar la salida del log: %v", err)
	}

	if logEntry["msg"] != "test message" {
		t.Errorf("Mensaje esperado: 'test message', obtenido: %v", logEntry["msg"])
	}

	// Restaurar el logger original
	SetDefault(&originalLogger)
}

func TestWith(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{})
	slogLogger := slog.New(handler)
	logger := NewSlogLogger(slogLogger)

	ctx := context.Background()
	ctx = ContextWithLogger(ctx, logger)

	// Añadir campos al contexto
	ctxWithFields := With(ctx, "context_field", "context_value")

	// Usar el contexto para registrar
	Info(ctxWithFields, "test message", "key", "value")

	var logEntry map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
		t.Fatalf("Error al analizar la salida del log: %v", err)
	}

	if logEntry["context_field"] != "context_value" {
		t.Errorf("Valor esperado para 'context_field': 'context_value', obtenido: %v", logEntry["context_field"])
	}
	if logEntry["key"] != "value" {
		t.Errorf("Valor esperado para 'key': 'value', obtenido: %v", logEntry["key"])
	}
}

func TestHelperFunctions(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slogLogger := slog.New(handler)
	logger := NewSlogLogger(slogLogger)

	ctx := context.Background()
	ctx = ContextWithLogger(ctx, logger)

	tests := []struct {
		name    string
		logFunc func(context.Context, string, ...any)
		level   string
	}{
		{"Debug", Debug, "DEBUG"},
		{"Info", Info, "INFO"},
		{"Warn", Warn, "WARN"},
		{"Error", Error, "ERROR"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			tt.logFunc(ctx, "test "+tt.name, "key", tt.name)

			var logEntry map[string]interface{}
			if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
				t.Fatalf("Error al analizar la salida del log: %v", err)
			}

			if logEntry["level"] != tt.level {
				t.Errorf("Nivel esperado: '%s', obtenido: %v", tt.level, logEntry["level"])
			}
			if logEntry["key"] != tt.name {
				t.Errorf("Valor esperado para 'key': '%s', obtenido: %v", tt.name, logEntry["key"])
			}
		})
	}
}
