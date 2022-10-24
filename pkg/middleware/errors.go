package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/pkg/sys/validate"
)

func Error() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		for _, err := range ctx.Errors {
			unWrapError := err.Unwrap()
			switch unWrapError.(type) {
			case validate.FieldErrors:
				ctx.JSON(http.StatusBadRequest, unWrapError)
			case *validate.RequestError:
				log.Println("")
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "server error",
				})
			}
		}
		ctx.Abort()
	}
}
