package main

import (
	_ "github.com/dongdongjssy/order-service/docs"
	"github.com/dongdongjssy/order-service/handlers"
	"github.com/dongdongjssy/order-service/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	ENDPOINT_ORDERS_TRANSFORM = "/v1/orders/transform"
)

func setupRouter() *gin.Engine {
	server := gin.Default()

	// TODO: add needed middlewares
	// - auth check
	server.Use(middlewares.AuthMiddleware)
	// - rate limit?
	// - timeout?
	// - XSS, CORS etc.

	// register swagger ui path
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.POST(ENDPOINT_ORDERS_TRANSFORM, handlers.TransformOrders)
	return server
}

func main() {
	server := setupRouter()
	server.Run(":8080")
}
