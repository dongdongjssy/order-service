package ma

import (
	"github.com/dongdongjssy/order-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/orders", handlers.SaveOrder)
	server.GET("/items/:customerId", handlers.GetItemsForCustomer)
	server.GET("/summaries", handlers.GetSummaries)

	server.Run(":8000")
}
