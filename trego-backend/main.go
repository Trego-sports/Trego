package main

import (
	"fmt"
	"log"

	"trego-backend/api-gateway/config"
	"trego-backend/api-gateway/logger"
	"trego-backend/api-gateway/web"
	"trego-backend/database"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

// main is the entry point of the Trego API Gateway service
func main() {
	// Load configuration
	conf := config.New()

	// Create logger
	logger := logger.New()

	// Initialize database connection
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Run database migrations
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Run the server
	run(conf, logger)
}

// run sets up and starts an HTTP server with the given configurations
// It blocks program execution while the server is running
func run(conf *config.Config, logger logger.Logger) {
	// Set Gin mode
	gin.SetMode(conf.GinMode)

	// Create Gin engine
	ginEngine := gin.New()

	// Setup router with configuration
	web.SetupRouter(ginEngine.Group("/"),
		func(opt *web.Options) {
			opt.Config = conf
		},
		func(opt *web.Options) {
			opt.Logger = logger
		},
	)

	// Start server with graceful shutdown
	log.Printf("Starting Trego API Gateway server on port %s", conf.Port)
	err := endless.ListenAndServe(fmt.Sprintf(":%s", conf.Port), ginEngine)
	log.Printf("Server closed, error: %v", err)
}
