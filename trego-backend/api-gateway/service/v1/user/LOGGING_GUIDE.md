# Logging Guide

This document explains how logging works across the different layers of the user API.

## Overview

Logging is done using the `logger.Logger` interface from the `trego-backend/api-gateway/logger` package. The logger supports structured logging with key-value pairs and maintains trace IDs for request tracking.

## Logging Flow

```
Request → Handler → Service → Repository (future)
         ↙         ↙        ↙
      logger     logger   logger
```

### 1. Handler Layer (handler.go)

**Getting the logger:**
```go
log := ginmiddleware.GetLoggerFromContext(ctx)
```

The logger is retrieved from the Gin context and automatically includes:
- Trace ID (for request tracking)
- Request metadata

**Example usage:**
```go
log.Info("Creating user",
    logger.Field{Key: "email", Value: req.Email},
    logger.Field{Key: "name", Value: req.Name},
)

log.Error("Failed to create user",
    logger.Field{Key: "email", Value: req.Email},
    logger.Field{Key: "error", Value: err.Error()},
)
```

**Passing to service layer:**
```go
user, err := h.service.CreateUser(log, &req)
```

### 2. Service Layer (service.go)

**Receiving the logger:**
```go
func (s *Service) CreateUser(log logger.Logger, req *models.CreateUserRequest) (*models.User, error) {
    // Use the logger here
    log.Debug("Business logic started",
        logger.Field{Key: "email", Value: req.Email},
    )
    
    // TODO: Call repository
    // return s.repo.CreateUser(log, userData)
}
```

**Key points:**
- Logger is passed as the first parameter
- Maintains trace ID from the request
- Can log business logic operations

**Example usage:**
```go
log.Debug("Checking if email exists",
    logger.Field{Key: "email", Value: email},
)

log.Warn("Email already exists",
    logger.Field{Key: "email", Value: email},
)

log.Error("Database query failed",
    logger.Field{Key: "query", Value: "SELECT * FROM users..."},
    logger.Field{Key: "error", Value: err.Error()},
)
```

### 3. Repository Layer (repo.go - future)

**Pattern:**
```go
func (r *Repository) GetUserByEmail(log logger.Logger, email string) (*models.User, error) {
    log.Debug("Executing database query",
        logger.Field{Key: "query", Value: "SELECT * FROM users WHERE email = $1"},
        logger.Field{Key: "email", Value: email},
    )
    
    // Execute query
    // ...
    
    log.Debug("Query completed",
        logger.Field{Key: "rows_returned", Value: 1},
        logger.Field{Key: "duration_ms", Value: 42},
    )
    
    return user, nil
}
```

## Log Levels

Use the appropriate log level:

- **Debug** - Detailed information for debugging (e.g., "Executing database query", "Merging update fields")
- **Info** - General informational messages (e.g., "User created successfully", "Request received")
- **Warn** - Warning messages (e.g., "Feature not implemented", "Retrying connection")
- **Error** - Error messages (e.g., "Database connection failed", "Validation failed")

## Best Practices

### 1. Always Pass Logger

✅ **Correct:**
```go
func (s *Service) CreateUser(log logger.Logger, req *CreateUserRequest) {}
```

❌ **Incorrect:**
```go
func (s *Service) CreateUser(req *CreateUserRequest) {}  // Missing logger
```

### 2. Use Structured Logging

✅ **Correct:**
```go
log.Info("User retrieved",
    logger.Field{Key: "user_id", Value: userID},
    logger.Field{Key: "email", Value: user.Email},
)
```

❌ **Incorrect:**
```go
log.Info(fmt.Sprintf("User retrieved: %s", userID))  // Loses structured fields
```

### 3. Include Relevant Context

✅ **Correct:**
```go
log.Error("Failed to update user",
    logger.Field{Key: "user_id", Value: userID},
    logger.Field{Key: "update_fields", Value: updatedFields},
    logger.Field{Key: "error", Value: err.Error()},
)
```

❌ **Incorrect:**
```go
log.Error("Failed to update user")  // Not enough context
```

### 4. Maintain Trace ID

The trace ID is automatically included when you:
1. Get logger from context in the handler
2. Pass that logger to service methods
3. Pass that logger to repository methods

This allows you to trace a single request across all layers:
```
[2024-01-15 10:30:15] INFO Request received trace_id=abc123
[2024-01-15 10:30:15] DEBUG Getting user by email trace_id=abc123 email=user@example.com
[2024-01-15 10:30:15] DEBUG Executing database query trace_id=abc123 query=SELECT...
[2024-01-15 10:30:15] INFO Request completed trace_id=abc123
```

## Complete Example

```go
// Handler
func (h *Handler) CreateUser(ctx *gin.Context) {
    log := ginmiddleware.GetLoggerFromContext(ctx)
    
    var req models.CreateUserRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        log.Warn("Invalid request",
            logger.Field{Key: "error", Value: err.Error()},
        )
        ctx.JSON(400, gin.H{"error": "invalid request"})
        return
    }
    
    log.Info("Creating user",
        logger.Field{Key: "email", Value: req.Email},
    )
    
    user, err := h.service.CreateUser(log, &req)
    if err != nil {
        log.Error("Failed to create user",
            logger.Field{Key: "email", Value: req.Email},
            logger.Field{Key: "error", Value: err.Error()},
        )
        ctx.JSON(500, gin.H{"error": "failed to create user"})
        return
    }
    
    log.Info("User created successfully",
        logger.Field{Key: "user_id", Value: user.UserID},
    )
    
    ctx.JSON(201, user)
}

// Service
func (s *Service) CreateUser(log logger.Logger, req *models.CreateUserRequest) (*models.User, error) {
    log.Debug("Validating email uniqueness",
        logger.Field{Key: "email", Value: req.Email},
    )
    
    // Check if email exists
    exists, err := s.repo.EmailExists(log, req.Email)
    if err != nil {
        log.Error("Failed to check email",
            logger.Field{Key: "error", Value: err.Error()},
        )
        return nil, err
    }
    
    if exists {
        log.Warn("Email already exists",
            logger.Field{Key: "email", Value: req.Email},
        )
        return nil, ErrEmailExists
    }
    
    log.Debug("Generating user ID")
    userID := generateUUID()
    
    user := &models.User{
        UserID:     userID,
        Name:       req.Name,
        Email:      req.Email,
        // ...
    }
    
    log.Debug("Saving user to database",
        logger.Field{Key: "user_id", Value: userID},
    )
    
    createdUser, err := s.repo.CreateUser(log, user)
    if err != nil {
        log.Error("Failed to save user",
            logger.Field{Key: "user_id", Value: userID},
            logger.Field{Key: "error", Value: err.Error()},
        )
        return nil, err
    }
    
    log.Info("User created successfully",
        logger.Field{Key: "user_id", Value: createdUser.UserID},
    )
    
    return createdUser, nil
}

// Repository
func (r *Repository) CreateUser(log logger.Logger, user *models.User) (*models.User, error) {
    log.Debug("Inserting user into database",
        logger.Field{Key: "user_id", Value: user.UserID},
    )
    
    query := "INSERT INTO users (...) VALUES (...)"
    
    err := r.db.QueryRow(query, ...).Scan(...)
    if err != nil {
        log.Error("Database insert failed",
            logger.Field{Key: "query", Value: query},
            logger.Field{Key: "error", Value: err.Error()},
        )
        return nil, err
    }
    
    log.Debug("User inserted successfully",
        logger.Field{Key: "user_id", Value: user.UserID},
    )
    
    return user, nil
}
```

## Summary

1. **Handler**: Get logger from context with `ginmiddleware.GetLoggerFromContext(ctx)`
2. **Service**: Receive logger as first parameter, pass to repository
3. **Repository**: Receive logger as first parameter, use for query logging
4. **Always** pass the logger through the call chain to maintain trace IDs
5. **Use** structured logging with `logger.Field`
6. **Choose** appropriate log levels (Debug, Info, Warn, Error)

