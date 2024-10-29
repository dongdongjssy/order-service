package middlewares

import (
	"net/http"
	"strings"

	"github.com/dongdongjssy/order-service/model"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	jwtToken := ctx.GetHeader("authorization")
	if !strings.HasPrefix(jwtToken, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, model.Response{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
			Errors:  []string{"Invalid JWT token, must start with 'Bearer '"},
		})
		return
	}

	slicedToken := jwtToken[strings.Index(jwtToken, " "):]
	// fake auth for local dev use
	if slicedToken != "abc" {
		// call auth service to verify token
	}

	ctx.Next()
}
