package user

import (
	"errors"
	"net/http"
	"strconv"

	"trego-backend/api-gateway/config"
	ginmiddleware "trego-backend/api-gateway/internal/middleware"
	"trego-backend/api-gateway/logger"
	"trego-backend/models"

	"github.com/gin-gonic/gin"
)

// Handler handles user-related HTTP requests
type Handler struct {
	config  *config.Config
	service *Service
}

// NewHandler creates a new user handler
func NewHandler(config *config.Config, service *Service) *Handler {
	return &Handler{
		config:  config,
		service: service,
	}
}

// GetUserByEmail handles GET /api/v1/user/email/:email
// Query user by email (used as username identifier)
func (h *Handler) GetUserByEmail(ctx *gin.Context) {
	log := ginmiddleware.GetLoggerFromContext(ctx)

	// Extract email from URL parameter
	email := ctx.Param("email")
	if email == "" {
		log.Warn("GetUserByEmail called with empty email",
			logger.Field{Key: "endpoint", Value: "/api/v1/user/:email"},
			logger.Field{Key: "method", Value: ctx.Request.Method},
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email parameter is required"})
		return
	}

	// Validate email format
	if err := validateEmail(email); err != nil {
		log.Warn("Invalid email format",
			logger.Field{Key: "email", Value: email},
			logger.Field{Key: "error", Value: err.Error()},
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid email format"})
		return
	}

	log.Info("GetUserByEmail request received",
		logger.Field{Key: "email", Value: email},
	)

	// Call service layer to get user
	user, err := h.service.GetUserByEmail(log, email)
	if err != nil {
		log.Error("Failed to get user",
			logger.Field{Key: "email", Value: email},
			logger.Field{Key: "error", Value: err.Error()},
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user"})
		return
	}

	if user == nil {
		log.Info("User not found",
			logger.Field{Key: "email", Value: email},
		)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	log.Info("GetUserByEmail completed successfully",
		logger.Field{Key: "email", Value: email},
		logger.Field{Key: "user_id", Value: user.UserID},
	)

	ctx.JSON(http.StatusOK, user)
}

// CreateUser handles POST /api/v1/user
// Add a new user
func (h *Handler) CreateUser(ctx *gin.Context) {
	log := ginmiddleware.GetLoggerFromContext(ctx)

	// Parse and validate request body
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Warn("Invalid request body",
			logger.Field{Key: "error", Value: err.Error()},
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	// Additional validation
	if req.Name == "" {
		log.Warn("Name is required")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	if err := validateEmail(req.Email); err != nil {
		log.Warn("Invalid email format",
			logger.Field{Key: "error", Value: err.Error()},
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid email format"})
		return
	}

	log.Info("CreateUser request received",
		logger.Field{Key: "email", Value: req.Email},
		logger.Field{Key: "name", Value: req.Name},
	)

	// Call service layer to create user
	user, err := h.service.CreateUser(log, &req)
	if err != nil {
		log.Error("Failed to create user",
			logger.Field{Key: "email", Value: req.Email},
			logger.Field{Key: "error", Value: err.Error()},
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	log.Info("CreateUser completed successfully",
		logger.Field{Key: "user_id", Value: user.UserID},
		logger.Field{Key: "email", Value: user.Email},
	)

	ctx.JSON(http.StatusCreated, user)
}

// UpdateUser handles PUT /api/v1/user/:user_id
// Modify user data
func (h *Handler) UpdateUser(ctx *gin.Context) {
	log := ginmiddleware.GetLoggerFromContext(ctx)

	// Extract user_id from URL parameter
	userID := ctx.Param("user_id")
	if userID == "" {
		log.Warn("UpdateUser called with empty user_id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id parameter is required"})
		return
	}

	// Parse and validate request body
	var req models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Warn("Invalid request body",
			logger.Field{Key: "error", Value: err.Error()},
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	// Validate that at least one field is being updated
	if req.Name == nil && req.PhoneNumber == nil && req.Location == nil && req.PictureURL == nil {
		log.Warn("No fields to update")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "at least one field must be provided for update"})
		return
	}

	log.Info("UpdateUser request received",
		logger.Field{Key: "user_id", Value: userID},
	)

	// Call service layer to update user
	user, err := h.service.UpdateUser(log, userID, &req)
	if err != nil {
		log.Error("Failed to update user",
			logger.Field{Key: "user_id", Value: userID},
			logger.Field{Key: "error", Value: err.Error()},
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	if user == nil {
		log.Info("User not found",
			logger.Field{Key: "user_id", Value: userID},
		)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	log.Info("UpdateUser completed successfully",
		logger.Field{Key: "user_id", Value: user.UserID},
	)

	ctx.JSON(http.StatusOK, user)
}

// GetUserByID handles GET /api/v1/user/:user_id
// Query user by user_id (alternative to GetUserByEmail)
func (h *Handler) GetUserByID(ctx *gin.Context) {
	log := ginmiddleware.GetLoggerFromContext(ctx)

	// Extract user_id from URL parameter
	userID := ctx.Param("user_id")
	if userID == "" {
		log.Warn("GetUserByID called with empty user_id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id parameter is required"})
		return
	}

	log.Info("GetUserByID request received",
		logger.Field{Key: "user_id", Value: userID},
	)

	// Call service layer to get user
	user, err := h.service.GetUserByID(log, userID)
	if err != nil {
		log.Error("Failed to get user",
			logger.Field{Key: "user_id", Value: userID},
			logger.Field{Key: "error", Value: err.Error()},
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user"})
		return
	}

	if user == nil {
		log.Info("User not found",
			logger.Field{Key: "user_id", Value: userID},
		)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	log.Info("GetUserByID completed successfully",
		logger.Field{Key: "user_id", Value: user.UserID},
		logger.Field{Key: "email", Value: user.Email},
	)

	ctx.JSON(http.StatusOK, user)
}

// ListUsers handles GET /api/v1/users
// List all users with optional pagination
func (h *Handler) ListUsers(ctx *gin.Context) {
	log := ginmiddleware.GetLoggerFromContext(ctx)

	// Extract query parameters
	limitStr := ctx.DefaultQuery("limit", "10")
	offsetStr := ctx.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > 100 {
		log.Warn("Invalid limit parameter",
			logger.Field{Key: "limit", Value: limitStr},
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "limit must be a number between 1 and 100"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		log.Warn("Invalid offset parameter",
			logger.Field{Key: "offset", Value: offsetStr},
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "offset must be a non-negative number"})
		return
	}

	log.Info("ListUsers request received",
		logger.Field{Key: "limit", Value: limit},
		logger.Field{Key: "offset", Value: offset},
	)

	// Call service layer to list users
	users, err := h.service.ListUsers(log, limit, offset)
	if err != nil {
		log.Error("Failed to list users",
			logger.Field{Key: "error", Value: err.Error()},
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
		return
	}

	log.Info("ListUsers completed successfully",
		logger.Field{Key: "count", Value: len(users)},
	)

	ctx.JSON(http.StatusOK, gin.H{
		"users":  users,
		"limit":  limit,
		"offset": offset,
	})
}

// validateEmail performs basic email format validation
func validateEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}

	// Basic email validation (contains @ and .)
	hasAt := false
	hasDot := false
	for _, char := range email {
		if char == '@' {
			hasAt = true
		}
		if char == '.' && hasAt {
			hasDot = true
		}
	}

	if !hasAt || !hasDot {
		return errors.New("invalid email format")
	}

	return nil
}
