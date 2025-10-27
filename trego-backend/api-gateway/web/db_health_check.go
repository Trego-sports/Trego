package web

import (
	"context"
	"net/http"
	"time"

	"trego-backend/api-gateway/config"
	ginmiddleware "trego-backend/api-gateway/internal/middleware"
	"trego-backend/api-gateway/logger"
	"trego-backend/database"

	"github.com/gin-gonic/gin"
)

type dbHealthCheckAPIHandler struct {
	Conf *config.Config
}

// @Summary		Database Health Check API
// @Description	This API is used to check the database connectivity and health
// @Tags			Database
// @Router			/dbHealthCheck [get]
// @Accept			json
// @Produce		json
// @Success		200	{object}	string	"{"status": "ok", "database": "ok", "timestamp": "2024-01-15T10:30:45Z"}"
// @Failure		503	{object}	string	"{"status": "error", "database": "error", "timestamp": "2024-01-15T10:30:45Z"}"
func (h *dbHealthCheckAPIHandler) dbHealthCheck(ctx *gin.Context) {
	// Get the logger from context (includes trace ID)
	log := ginmiddleware.GetLoggerFromContext(ctx)

	// Log the request
	log.Info("Database health check requested",
		logger.Field{Key: "endpoint", Value: "/dbHealthCheck"},
		logger.Field{Key: "method", Value: ctx.Request.Method},
		logger.Field{Key: "user_agent", Value: ctx.Request.UserAgent()},
	)

	// Check database health with timeout
	dbCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	dbStatus := "ok"
	var dbError string
	if err := database.HealthCheck(dbCtx); err != nil {
		dbStatus = "error"
		dbError = err.Error()
		log.Error("Database health check failed",
			logger.Field{Key: "error", Value: err.Error()},
		)
	}

	// Log database status for debugging
	log.Debug("Database health check response",
		logger.Field{Key: "database_status", Value: dbStatus},
	)

	// Determine HTTP status code
	status := http.StatusOK
	if dbStatus == "error" {
		status = http.StatusServiceUnavailable
	}

	// Prepare response
	response := gin.H{
		"status":    dbStatus,
		"database":  dbStatus,
		"timestamp": time.Now().UTC(),
	}

	// Include error details if there's an error
	if dbError != "" {
		response["error"] = dbError
	}

	ctx.JSON(status, response)

	// Log completion
	if dbStatus == "ok" {
		log.Info("Database health check completed successfully")
	} else {
		log.Warn("Database health check completed with errors")
	}
}
