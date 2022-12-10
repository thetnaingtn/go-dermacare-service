package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/auth"
)

func Authenticate(auth *auth.Auth) gin.HandlerFunc {
	f := func(ctx *gin.Context) {
		authStr := ctx.Request.Header.Get("authorization")

		parts := strings.Split(authStr, " ")

		if len(parts) != 2 || parts[0] != "bearer" {
			err := errors.New("expect authorization format: bearer <token>")
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		claim, err := auth.ValidateToken(parts[1])
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		ctx.Set("claims", claim)
		ctx.Next()

	}

	return f

}
