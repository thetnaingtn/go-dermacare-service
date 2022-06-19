package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
	"github.com/thetnaingtn/go-dermacare-service/pkg/listing"
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

		if err == adding.ErrOutOfStock {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": err.Error(),
			})
			return
		}

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

func GetOrders(service listing.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orders, err := service.GetOrders()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Can't list order",
			})
			log.Println(err)
			return
		}
		ctx.JSON(200, gin.H{
			"message": "Successfully retrive orders",
			"orders":  orders,
		})
	}
}
