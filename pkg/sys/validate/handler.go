package validate

import "github.com/gin-gonic/gin"

type Handler func(*gin.Context) error

// function which accept handler that return error and transform it into a handler that gin accept.
func ErrHandler(handler Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := handler(ctx); err != nil {
			ctx.Error(err)
		}
	}
}
