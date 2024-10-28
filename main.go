package main

import (
	_ "github.com/dongdongjssy/order-service/docs"
	"github.com/dongdongjssy/order-service/handlers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	ENDPOINT_ORDERS_TRANSFORM = "/v1/orders/transform"
)

func setupRouter() *gin.Engine {
	server := gin.Default()

	// register swagger ui path
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.POST(ENDPOINT_ORDERS_TRANSFORM, handlers.TransformOrders)
	return server
}

func main() {
	server := setupRouter()
	server.Run(":8080")
}
