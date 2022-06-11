package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
)

func AddOrder(service adding.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var order adding.OrderForm
		if err := ctx.ShouldBind(&order); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Can't parse incoming request data",
			})
			log.Println(err)
			return
		}

		id, err := service.AddOrder(order)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Can't create order",
			})
			log.Println(err)
			return
		}

		ctx.JSON(201, gin.H{
			"message": "Successfully create order",
			"id":      id,
		})
	}
}
