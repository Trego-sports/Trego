package config

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Config holds all configuration for the API Gateway
type Config struct {
	Port               string
	GinMode            string
	LogLevel           string
	BuildVersion       string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURI  string
	FrontendURL        string
}

// New creates a new configuration instance with default values
// and overrides them with environment variables if present
func New() *Config {
	config := &Config{
		Port:               getEnv("PORT", "8080"),
		GinMode:            getEnv("GIN_MODE", gin.ReleaseMode),
		LogLevel:           getEnv("LOG_LEVEL", "info"),
		BuildVersion:       getEnv("BUILD_VERSION", "1.0.0"),
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleRedirectURI:  getEnv("GOOGLE_REDIRECT_URI", "http://localhost:8080/api/v1/google-callback"),
		FrontendURL:        getEnv("FRONTEND_URL", "http://localhost:3000"),
	}

	return config
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt gets an environment variable as integer with a fallback default value
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
