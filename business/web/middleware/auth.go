package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/auth"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/validate"
)

func Authenticate(auth *auth.Auth) gin.HandlerFunc {
	f := func(ctx *gin.Context) {
		authStr := ctx.Request.Header.Get("authorization")

		parts := strings.Split(authStr, " ")

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, validate.ErrorResponse{Error: "expect authorization format: Bearer <token>"})
			return
		}

		claim, err := auth.ValidateToken(parts[1])

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, validate.ErrorResponse{Error: err.Error()})
			return
		}

		ctx.Set("claims", claim)
		ctx.Next()

	}

	return f

}
