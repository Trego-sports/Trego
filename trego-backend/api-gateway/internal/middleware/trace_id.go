package ginmiddleware

import (
	"trego-backend/api-gateway/internal/constant"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// NewTraceIDMiddleware creates a middleware that generates and stores trace ID
func NewTraceIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.Request.Header.Get(constant.HTTPHeaderTraceID)
		if traceID == "" {
			traceID = uuid.New().String()
			c.Request.Header.Set(constant.HTTPHeaderTraceID, traceID)
		}
		c.Set(constant.GinContextTraceIDKey, traceID)
		c.Next()
	}
}

// GetTraceIDFromContext retrieves trace ID from Gin context
func GetTraceIDFromContext(c *gin.Context) string {
	if traceID, exists := c.Get(constant.GinContextTraceIDKey); exists {
		if id, ok := traceID.(string); ok {
			return id
		}
	}
	return ""
}
