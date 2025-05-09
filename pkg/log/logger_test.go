package log

import (
	"context"
	"testing"
)

// MockLogger implementa la interfaz Logger para propósitos de prueba
type MockLogger struct {
	debugCalled bool
	infoCalled  bool
	warnCalled  bool
	errorCalled bool
	fatalCalled bool
	withCalled  bool
	lastMsg     string
	lastFields  []any
}

func NewMockLogger() *MockLogger {
	return &MockLogger{}
}

func (m *MockLogger) With(fields ...any) Logger {
	m.withCalled = true
	m.lastFields = fields
	return m
}

func (m *MockLogger) Debug(ctx context.Context, msg string, fields ...any) {
	m.debugCalled = true
	m.lastMsg = msg
	m.lastFields = fields
}

func (m *MockLogger) Info(ctx context.Context, msg string, fields ...any) {
	m.infoCalled = true
	m.lastMsg = msg
	m.lastFields = fields
}

func (m *MockLogger) Warn(ctx context.Context, msg string, fields ...any) {
	m.warnCalled = true
	m.lastMsg = msg
	m.lastFields = fields
}

func (m *MockLogger) Error(ctx context.Context, msg string, fields ...any) {
	m.errorCalled = true
	m.lastMsg = msg
	m.lastFields = fields
}

func (m *MockLogger) Fatal(ctx context.Context, msg string, fields ...any) {
	m.fatalCalled = true
	m.lastMsg = msg
	m.lastFields = fields
}

func TestLoggerInterface(t *testing.T) {
	mock := NewMockLogger()
	ctx := context.Background()

	// Probar que todos los métodos del logger son invocados correctamente

	mock.Debug(ctx, "debug message", "key", "value")
	if !mock.debugCalled || mock.lastMsg != "debug message" {
		t.Error("Debug no fue invocado correctamente")
	}

	mock.Info(ctx, "info message", "key", "value")
	if !mock.infoCalled || mock.lastMsg != "info message" {
		t.Error("Info no fue invocado correctamente")
	}

	mock.Warn(ctx, "warn message", "key", "value")
	if !mock.warnCalled || mock.lastMsg != "warn message" {
		t.Error("Warn no fue invocado correctamente")
	}

	mock.Error(ctx, "error message", "key", "value")
	if !mock.errorCalled || mock.lastMsg != "error message" {
		t.Error("Error no fue invocado correctamente")
	}

	// No probamos Fatal porque terminaría la ejecución del programa

	// Probar With
	newLogger := mock.With("with_key", "with_value")
	if !mock.withCalled {
		t.Error("With no fue invocado")
	}

	// Verificar que With retorna un logger válido
	if newLogger == nil {
		t.Error("With retornó un logger nil")
	}
}

func TestLoggerContextInteraction(t *testing.T) {
	mock := NewMockLogger()
	ctx := context.Background()

	// Probar la integración con el contexto
	ctx = ContextWithLogger(ctx, mock)
	logger := FromContext(ctx)

	if logger != mock {
		t.Error("El logger recuperado del contexto no es el mismo que se almacenó")
	}

	// Probar las funciones ayudantes con el contexto
	Info(ctx, "context test")
	if !mock.infoCalled || mock.lastMsg != "context test" {
		t.Error("La función ayudante Info no invocó el logger del contexto correctamente")
	}
}
