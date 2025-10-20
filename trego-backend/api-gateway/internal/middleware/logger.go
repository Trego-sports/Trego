package ginmiddleware

import (
	"trego-backend/api-gateway/internal/constant"
	"trego-backend/api-gateway/logger"

	"github.com/gin-gonic/gin"
)

// NewLoggerMiddleware creates a logger middleware that stores logger in context
func NewLoggerMiddleware(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := GetTraceIDFromContext(c)
		entry := log.WithField("trace_id", traceID)
		c.Set(constant.GinContextLoggerKey, entry)
		c.Next()
	}
}

// GetLoggerFromContext retrieves logger from Gin context
func GetLoggerFromContext(c *gin.Context) logger.Logger {
	if log, exists := c.Get(constant.GinContextLoggerKey); exists {
		if l, ok := log.(logger.Logger); ok {
			return l
		}
	}
	// Return a default logger if none found
	return logger.New()
}
