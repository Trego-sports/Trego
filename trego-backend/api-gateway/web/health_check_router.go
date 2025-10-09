package web

import (
	"trego-backend/api-gateway/config"
	"github.com/gin-gonic/gin"
)

const (
	healthCheckURL = "/healthCheck"
)

func setupHealthCheckHandler(routerGroup *gin.RouterGroup, conf *config.Config, middlewares ...gin.HandlerFunc) {
	for _, m := range middlewares {
		routerGroup.Use(m)
	}

	handler := &healthCheckAPIHandler{Conf: conf}
	routerGroup.GET(healthCheckURL, handler.healthCheck)
}