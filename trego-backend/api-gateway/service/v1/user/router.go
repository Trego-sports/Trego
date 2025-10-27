package user

import (
	"trego-backend/api-gateway/config"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configures user-related routes
func SetupUserRoutes(routerGroup *gin.RouterGroup, conf *config.Config) {
	// Create service instance
	service := NewService()

	// Create handler instance
	handler := NewHandler(conf, service)

	// User routes
	// GET /api/v1/user/email/:email - Get user by email (username)
	routerGroup.GET("/user/email/:email", handler.GetUserByEmail)
	// GET /api/v1/user/:user_id - Get user by user_id
	routerGroup.GET("/user/:user_id", handler.GetUserByID)
	// POST /api/v1/user - Create a new user
	routerGroup.POST("/user", handler.CreateUser)
	// PUT /api/v1/user/:user_id - Update user by user_id
	routerGroup.PUT("/user/:user_id", handler.UpdateUser)

	// List users route
	// GET /api/v1/users - List all users with pagination
	routerGroup.GET("/users", handler.ListUsers)
}
