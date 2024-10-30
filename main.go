package main

import (
	"context"
	"net/http"
	"strings"
	"time"

	_ "github.com/dongdongjssy/order-service/docs"
	"github.com/dongdongjssy/order-service/handlers"
	"github.com/dongdongjssy/order-service/model"
	"github.com/gin-contrib/gzip"
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
	server.Use(authMiddleware())
	// - compression
	server.Use(gzip.Gzip(gzip.DefaultCompression))
	// - rate limit? golang.org/x/time/rate
	// - timeout?
	// server.Use(timeoutMiddleware(5 * time.Second))
	// - XSS? github.com/microcosm-cc/bluemonday
	// - CORS? github.com/gin-contrib/cors

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
func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtToken := ctx.GetHeader("authorization")
		if !strings.HasPrefix(jwtToken, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, model.Response{
				Code:    http.StatusUnauthorized,
				Message: "unauthorized",
				Errors:  []string{"Invalid JWT token, must start with 'Bearer '"},
			})
			return
		}

		slicedToken := jwtToken[strings.Index(jwtToken, " ")+1:]
		if slicedToken == "abc" {
			ctx.Next()
		} else {
			// call auth service to verify token
			ctx.JSON(http.StatusUnauthorized, model.Response{
				Code:    http.StatusUnauthorized,
				Message: "unauthorized",
				Errors:  []string{"Invalid JWT token"},
			})
		}
	}
}

func timeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Update the request with the new context
		c.Request = c.Request.WithContext(ctx)
		done := make(chan struct{}, 1)

		// Execute the handler in a goroutine
		go func() {
			c.Next()
			done <- struct{}{}
		}()

		select {
		case <-done:
		// Handler completed within the timeout
		case <-ctx.Done():
			// Timeout reached, respond with 504 Gateway Timeout
			c.AbortWithStatus(http.StatusGatewayTimeout)
		}
	}
}
