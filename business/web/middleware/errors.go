package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/validate"
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
			case validate.FieldErrors:
				re = validate.ErrorResponse{
					Error:  "Field validation errors",
					Fields: err.Fields(),
				}
				status = http.StatusBadRequest
			default:
				re = validate.ErrorResponse{
					Error: "server error",
				}
				status = http.StatusInternalServerError
			}
			ctx.JSON(status, re)
		}
	}
}
