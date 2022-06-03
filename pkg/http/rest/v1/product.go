package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
)

func AddProduct(service *adding.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product adding.Product
		if err := ctx.ShouldBind(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Can't parse incoming request data",
			})
			return
		}

		id := service.AddProduct(product)
		ctx.JSON(201, gin.H{
			"message": "Successfully create product",
			"id":      id,
		})
	}
}
