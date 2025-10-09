package ginmiddleware

// This file is kept for backward compatibility and contains common middleware utilities.
// Individual middleware functions have been moved to separate files for better organization:
// - logger.go: Context-based logger middleware
// - trace_id.go: Trace ID generation and retrieval (using UUID)
// - cors.go: CORS middleware
// - recovery.go: Panic recovery middleware
// - ordered_middleware.go: Middleware ordering and setup
