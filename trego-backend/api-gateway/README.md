# Trego API Gateway

A well-structured API Gateway service built with Go, Gin, and following SOLID principles. Features context-based logging with UUID trace IDs, graceful shutdown, and clean middleware architecture.

## Architecture

The API Gateway follows a clean architecture pattern with clear separation of concerns:

```
api-gateway/
├── config/                    # Configuration management
│   └── config.go
├── logger/                    # Logger interface and implementation
│   └── logger.go
├── internal/
│   ├── constant/              # Application constants
│   │   ├── constant.go        # Gin context keys
│   │   └── http_header_key.go # HTTP header constants
│   └── middleware/            # HTTP middleware (modular)
│       ├── middleware.go      # Documentation
│       ├── ordered_middleware.go # Middleware ordering
│       ├── logger.go          # Context-based logger middleware
│       ├── trace_id.go        # UUID-based trace ID middleware
│       ├── cors.go            # CORS middleware
│       └── recovery.go        # Panic recovery middleware
└── web/                       # Router setup and handlers
    ├── web.go                 # Router configuration
    ├── health_check.go        # Health check handler
    └── health_check_router.go # Health check route setup
```

## Features

- **Health Check Endpoint**: `/healthCheck` with build version and timestamp
- **Context-Based Logging**: Request-scoped logging with UUID trace IDs
- **Graceful Shutdown**: Using endless for zero-downtime deployments
- **CORS Support**: Cross-origin resource sharing enabled
- **Structured Logging**: Field-based logging with trace ID correlation
- **UUID Trace IDs**: Standard UUID v4 format for distributed tracing
- **Modular Middleware**: Clean, single-responsibility middleware files
- **Configuration Management**: Environment-based configuration with build version

## SOLID Principles Implementation

### Single Responsibility Principle (SRP)
- Each package has a single, well-defined responsibility
- `config` handles configuration, `handlers` handle HTTP requests, etc.

### Open/Closed Principle (OCP)
- The router setup is open for extension through the `Options` pattern
- New handlers can be added without modifying existing code
- Middleware can be easily added or removed

### Liskov Substitution Principle (LSP)
- Handler interfaces can be easily substituted
- Configuration can be mocked for testing

### Interface Segregation Principle (ISP)
- Small, focused interfaces for different concerns
- Handlers implement only what they need

### Dependency Inversion Principle (DIP)
- High-level modules depend on abstractions (interfaces)
- Configuration is injected rather than hardcoded

## Configuration

The service can be configured using environment variables:

- `PORT`: Server port (default: 8080)
- `GIN_MODE`: Gin mode - debug, release, test (default: release)
- `LOG_LEVEL`: Logging level (default: info)
- `BUILD_VERSION`: Application build version (default: 1.0.0)

## Running the Service

```bash
# From the trego-backend directory
go run main.go

# Or run the api-gateway directly
go run api-gateway/main.go
```

## API Endpoints

### Health Check
- `GET /healthCheck` - Health check with build version and timestamp
  ```json
  {
    "buildVersion": "1.0.0",
    "time": "Mon Jan 2 15:04:05 MST 2006"
  }
  ```

### API Routes
- `GET /api/v1/ping` - Simple ping endpoint
  ```json
  {
    "message": "pong"
  }
  ```

## Logging

### Context-Based Logging
The API Gateway uses context-based logging where each request gets a unique logger instance with a UUID trace ID:

```go
// In any handler
log := ginmiddleware.GetLoggerFromContext(ctx)

// Log with structured fields
log.Info("User action", 
    logger.Field{Key: "user_id", Value: "123"},
    logger.Field{Key: "action", Value: "create_user"},
)

// Different log levels
log.Error("Database error", logger.Field{Key: "error", Value: err.Error()})
log.Debug("Debug info", logger.Field{Key: "query", Value: sqlQuery})
log.Warn("Warning message", logger.Field{Key: "retry_count", Value: 3})
```

### Trace ID Flow
1. **Request arrives** → Trace ID middleware generates/retrieves UUID
2. **Logger middleware** → Creates request-scoped logger with trace ID
3. **Handler** → Retrieves logger from context and logs with trace ID
4. **All logs** → Include the same trace ID for easy debugging

### Logger Replacement
The logger interface allows easy replacement of logging implementations:
- **Current**: SimpleLogger (stdout)
- **Future**: JSONLogger, FileLogger, ExternalLogger (ELK, Datadog, etc.)

## Middleware Architecture

The middleware system is designed with modularity in mind:

### Current Middleware Stack (in order):
1. **Trace ID Middleware** - Generates/retrieves UUID trace IDs
2. **Logger Middleware** - Creates request-scoped logger with trace ID
3. **Recovery Middleware** - Handles panics gracefully
4. **CORS Middleware** - Handles cross-origin requests

### Adding New Middleware:
```go
// Create new middleware file: internal/middleware/auth.go
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Authentication logic
        c.Next()
    }
}

// Add to ordered_middleware.go
func GetOrderedMiddleware(logger logger.Logger) []gin.HandlerFunc {
    return []gin.HandlerFunc{
        NewTraceIDMiddleware(),
        NewLoggerMiddleware(logger),
        AuthMiddleware(),        // Add here
        RecoveryMiddleware(),
        CORSMiddleware(),
    }
}
```

## Testing with Postman

### Health Check Request:
```http
GET http://localhost:8080/healthCheck
```

### Expected Response:
```json
{
  "buildVersion": "1.0.0",
  "time": "Mon Jan 2 15:04:05 MST 2006"
}
```

### Headers:
- `x-trace-id`: Optional UUID for request tracing
- `Content-Type`: application/json

## Future Extensions

The architecture is designed to easily accommodate:

- **Authentication middleware** - JWT, OAuth, API keys
- **Rate limiting** - Per-user/IP rate limiting
- **Request validation** - Input validation and sanitization
- **Service discovery** - Dynamic backend service routing
- **Load balancing** - Multiple backend instances
- **Circuit breakers** - Fault tolerance patterns
- **Metrics collection** - Prometheus, StatsD integration
- **Request/Response transformation** - Data format conversion
- **Caching** - Redis/Memcached integration
