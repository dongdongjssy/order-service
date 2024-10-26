package main

import (
	"github.com/dongdongjssy/order-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/v1/orders", handlers.SaveOrder)
	server.GET("/v1/orders/:customerId", handlers.GetItemsForCustomer)
	server.GET("/v1/orders/summary", handlers.GetSummary)

	server.Run(":8000")
}
