package ma

import (
	"github.com/dongdongjssy/order-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/orders", handlers.SaveOrder)
	server.GET("/orders/:customerId", handlers.GetItemsForCustomer)
	server.GET("/orders/summary", handlers.GetSummary)

	server.Run(":8000")
}
