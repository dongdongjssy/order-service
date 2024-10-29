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

	// TODO: add needed middlewares
	// - auth check
	server.Use(authMiddleware)
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

// dummy auth middleware
func authMiddleware(ctx *gin.Context) {
	// jwtToken := ctx.GetHeader("authorization")
	// if !strings.HasPrefix(jwtToken, "Bearer ") {
	// 	ctx.JSON(http.StatusUnauthorized, model.Response{
	// 		Code:    http.StatusUnauthorized,
	// 		Message: "unauthorized",
	// 		Errors:  []string{"Invalid JWT token, must start with 'Bearer '"},
	// 	})
	// 	return
	// }

	// slicedToken := jwtToken[strings.Index(jwtToken, " "):]
	// call auth service to verify token

	ctx.Next()
}
