package web

import (
	"net/http"
	"time"

	"trego-backend/api-gateway/config"
	ginmiddleware "trego-backend/api-gateway/internal/middleware"
	"trego-backend/api-gateway/logger"

	"github.com/gin-gonic/gin"
)

type healthCheckAPIHandler struct {
	Conf *config.Config
}

// @Summary		Health Check api
// @Description	This API is used to check the service health
// @Tags			Default
// @Router			/healthCheck [get]
// @Accept			json
// @Produce		json
// @Success		200	{object}	string	"{"buildVersion": "1.0.0", "time": "Mon Jan 2 15:04:05 MST 2006"}"
func (h *healthCheckAPIHandler) healthCheck(ctx *gin.Context) {
	// Get the logger from context (includes trace ID)
	log := ginmiddleware.GetLoggerFromContext(ctx)

	// Log the request
	log.Info("Server health check requested",
		logger.Field{Key: "endpoint", Value: "/healthCheck"},
		logger.Field{Key: "method", Value: ctx.Request.Method},
		logger.Field{Key: "user_agent", Value: ctx.Request.UserAgent()},
	)

	// Log build version for debugging
	log.Debug("Server health check response",
		logger.Field{Key: "build_version", Value: h.Conf.BuildVersion},
	)

	ctx.JSON(http.StatusOK, gin.H{
		"status":       "ok",
		"buildVersion": h.Conf.BuildVersion,
		"timestamp":    time.Now().UTC(),
	})

	// Log successful response
	log.Info("Server health check completed successfully")
}
