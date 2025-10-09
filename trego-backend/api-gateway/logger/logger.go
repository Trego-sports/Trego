package logger

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Logger interface defines the logging methods
type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
}

// Field represents a key-value pair for structured logging
type Field struct {
	Key   string
	Value interface{}
}

// SimpleLogger is a basic implementation of the Logger interface
type SimpleLogger struct {
	fields map[string]interface{}
}

// New creates a new logger instance
func New() Logger {
	return &SimpleLogger{
		fields: make(map[string]interface{}),
	}
}

// Info logs an info message
func (l *SimpleLogger) Info(msg string, fields ...Field) {
	l.log("INFO", msg, fields...)
}

// Error logs an error message
func (l *SimpleLogger) Error(msg string, fields ...Field) {
	l.log("ERROR", msg, fields...)
}

// Debug logs a debug message
func (l *SimpleLogger) Debug(msg string, fields ...Field) {
	l.log("DEBUG", msg, fields...)
}

// Warn logs a warning message
func (l *SimpleLogger) Warn(msg string, fields ...Field) {
	l.log("WARN", msg, fields...)
}

// WithField creates a new logger with an additional field
func (l *SimpleLogger) WithField(key string, value interface{}) Logger {
	newFields := make(map[string]interface{})
	for k, v := range l.fields {
		newFields[k] = v
	}
	newFields[key] = value
	
	return &SimpleLogger{fields: newFields}
}

// WithFields creates a new logger with additional fields
func (l *SimpleLogger) WithFields(fields map[string]interface{}) Logger {
	newFields := make(map[string]interface{})
	for k, v := range l.fields {
		newFields[k] = v
	}
	for k, v := range fields {
		newFields[k] = v
	}
	
	return &SimpleLogger{fields: newFields}
}

// log is the internal logging method
func (l *SimpleLogger) log(level, msg string, fields ...Field) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	
	// Build the log message
	logMsg := fmt.Sprintf("[%s] %s %s", timestamp, level, msg)
	
	// Add fields from the logger instance
	for key, value := range l.fields {
		logMsg += fmt.Sprintf(" %s=%v", key, value)
	}
	
	// Add fields from the method call
	for _, field := range fields {
		logMsg += fmt.Sprintf(" %s=%v", field.Key, field.Value)
	}
	
	// Use standard log package for now
	log.Println(logMsg)
}

// Context key for storing logger in request context
type contextKey string

const (
	// LoggerContextKey is the key used to store logger in context
	LoggerContextKey contextKey = "logger"
	// TraceIDContextKey is the key used to store trace ID in context
	TraceIDContextKey contextKey = "trace_id"
)

// FromContext retrieves the logger from the context
func FromContext(ctx context.Context) Logger {
	if logger, ok := ctx.Value(LoggerContextKey).(Logger); ok {
		return logger
	}
	// Return a default logger if none found in context
	return New()
}

// WithLogger adds a logger to the context
func WithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, LoggerContextKey, logger)
}

// GetTraceIDFromContext retrieves the trace ID from the context
func GetTraceIDFromContext(ctx context.Context) string {
	if traceID, ok := ctx.Value(TraceIDContextKey).(string); ok {
		return traceID
	}
	return ""
}

// WithTraceID adds a trace ID to the context
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDContextKey, traceID)
}
