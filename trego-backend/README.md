# Trego Backend

A Go backend service built with the Gin web framework.

## Prerequisites

- Go 1.21 or higher
- Git

## Getting Started

### Installation

1. Install dependencies:

```bash
go mod download
```

2. Run the server:

```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Available Endpoints

- `GET /health` - Health check endpoint
- `GET /api/v1/ping` - Test endpoint that returns "pong"

## Development

### Running in development mode

```bash
go run main.go
```

### Building the application

```bash
go build -o bin/trego-backend
```

### Running the built binary

```bash
./bin/trego-backend
```

## Project Structure

```
trego-backend/
├── main.go           # Application entry point
├── go.mod            # Go module dependencies
└── README.md         # This file
```

## Environment Variables

Currently no environment variables are required. As the project grows, add them here.

## License

[Add your license here]
