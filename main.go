package main

import (
	_ "github.com/dongdongjssy/order-service/docs"
	"github.com/dongdongjssy/order-service/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	if err := server.Run(":8080"); err != nil {
		log.Fatal("server fails to start duet to ", err)
	} else {
		log.Info("server is running on port 8080...")
	}
}
