package ginmiddleware

import (
	"trego-backend/api-gateway/logger"

	"github.com/gin-gonic/gin"
)

// GetOrderedMiddleware returns middleware in the correct order
func GetOrderedMiddleware(logger logger.Logger) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		NewTraceIDMiddleware(),
		NewLoggerMiddleware(logger),
		// RecoveryMiddleware(),
		// CORSMiddleware(),
	}
}