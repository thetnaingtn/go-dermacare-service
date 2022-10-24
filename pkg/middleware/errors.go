package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/pkg/sys/validate"
)

func Error() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		for _, errMsg := range ctx.Errors {
			unWrapErr := errMsg.Unwrap()
			var re validate.ErrorResponse
			var status int
			switch err := unWrapErr.(type) {
			case *validate.RequestError:
				re = validate.ErrorResponse{
					Error: err.Error(),
				}
				status = err.Status
			default:
				re = validate.ErrorResponse{
					Error: "server error",
				}
				status = http.StatusInternalServerError
			}
			ctx.JSON(status, re)
		}
		ctx.Abort()
	}
}
