package ginmiddleware

import (
	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware provides panic recovery
func RecoveryMiddleware() gin.HandlerFunc {
	return gin.Recovery()
}
