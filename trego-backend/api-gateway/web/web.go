package web

import (
	"trego-backend/api-gateway/config"
	ginmiddleware "trego-backend/api-gateway/internal/middleware"
	"trego-backend/api-gateway/logger"

	"github.com/gin-gonic/gin"
)

// Options holds configuration and dependencies for the web server
type Options struct {
	Config *config.Config
	Logger logger.Logger
}

// SetupRouter configures and sets up all routes for the API Gateway
func SetupRouter(routerGroup *gin.RouterGroup, optFuncs ...func(*Options)) {
	opt := Options{}
	for _, f := range optFuncs {
		f(&opt)
	}

	// Setup basic middlewares
	setupBasicMiddlewares(routerGroup, opt.Logger)

	// Setup health check routes
	setupHealthCheckHandler(routerGroup, opt.Config)

	// Setup database health check routes
	setupDbHealthCheckHandler(routerGroup, opt.Config)

	// Setup API routes
	setupAPIRoutes(routerGroup, opt.Config)
}

// setupBasicMiddlewares configures common middlewares for all routes
func setupBasicMiddlewares(routerGroup *gin.RouterGroup, logger logger.Logger) {
	// Get ordered middleware
	defaultMiddlewares := ginmiddleware.GetOrderedMiddleware(logger)
	for _, middleware := range defaultMiddlewares {
		routerGroup.Use(middleware)
	}
}

// setupHealthCheckRoutes configures health check endpoints (legacy - kept for reference)
// func setupHealthCheckRoutes(routerGroup *gin.RouterGroup) {
// 	healthHandler := handlers.NewHealthHandler()

// 	// Health check endpoints
// 	routerGroup.GET("/health", healthHandler.Check)
// 	routerGroup.GET("/health/ready", healthHandler.Readiness)
// 	routerGroup.GET("/health/live", healthHandler.Liveness)
// }

// setupAPIRoutes configures API routes
func setupAPIRoutes(routerGroup *gin.RouterGroup, cfg *config.Config) {
	// API v1 routes group
	v1 := routerGroup.Group("/api/v1")
	{
		// Placeholder for future API routes
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// Auth routes
		authHandler := NewAuthHandler(cfg)
		v1.POST("/google-login", authHandler.GoogleLoginHandler)
		v1.GET("/google-callback", authHandler.GoogleCallbackHandler)
	}
}
