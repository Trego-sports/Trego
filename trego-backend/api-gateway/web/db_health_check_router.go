package web

import (
	"trego-backend/api-gateway/config"

	"github.com/gin-gonic/gin"
)

const (
	dbHealthCheckURL = "/dbHealthCheck"
)

func setupDbHealthCheckHandler(routerGroup *gin.RouterGroup, conf *config.Config, middlewares ...gin.HandlerFunc) {
	for _, m := range middlewares {
		routerGroup.Use(m)
	}

	handler := &dbHealthCheckAPIHandler{Conf: conf}
	routerGroup.GET(dbHealthCheckURL, handler.dbHealthCheck)
}
