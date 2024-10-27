package main

import (
	"github.com/dongdongjssy/order-service/handlers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	server := gin.Default()
	server.POST("/v1/orders/transform", handlers.TransformOrders)
	return server
}

func main() {
	server := setupRouter()
	server.Run(":8080")
}
