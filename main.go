package main

import (
	"github.com/dongdongjssy/order-service/handlers"
	"github.com/gin-gonic/gin"
)

const (
	ENDPOINT_ORDERS_TRANSFORM = "/v1/orders/transform"
)

func setupRouter() *gin.Engine {
	server := gin.Default()
	server.POST(ENDPOINT_ORDERS_TRANSFORM, handlers.TransformOrders)
	return server
}

func main() {
	server := setupRouter()
	server.Run(":8080")
}
