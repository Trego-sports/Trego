# Trego API Gateway

A well-structured API Gateway service built with Go, Gin, and following SOLID principles.

## Architecture

The API Gateway follows a clean architecture pattern with clear separation of concerns:

```
api-gateway/
├── main.go              # Entry point and server startup
├── config/              # Configuration management
│   └── config.go
├── handlers/            # HTTP request handlers
│   └── health.go
├── middleware/          # HTTP middleware
│   └── middleware.go
└── web/                 # Router setup and route configuration
    └── web.go
```

## Features

- **Health Check Endpoints**: `/health`, `/health/ready`, `/health/live`
- **Graceful Shutdown**: Using endless for zero-downtime deployments
- **CORS Support**: Cross-origin resource sharing enabled
- **Request Logging**: Comprehensive request/response logging
- **Request ID Tracking**: Unique request IDs for tracing
- **Configuration Management**: Environment-based configuration

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

## Running the Service

```bash
# From the trego-backend directory
go run main.go

# Or run the api-gateway directly
go run api-gateway/main.go
```

## Health Check Endpoints

- `GET /health` - Basic health check
- `GET /health/ready` - Readiness check (for Kubernetes)
- `GET /health/live` - Liveness check (for Kubernetes)

## API Endpoints

- `GET /api/v1/ping` - Simple ping endpoint

## Future Extensions

The architecture is designed to easily accommodate:

- Authentication middleware
- Rate limiting
- Request validation
- Service discovery
- Load balancing
- Circuit breakers
- Metrics collection
