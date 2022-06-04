package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
	"github.com/thetnaingtn/go-dermacare-service/pkg/listing"
)

func AddProduct(service *adding.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product adding.Product
		if err := ctx.ShouldBind(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Can't parse incoming request data",
			})
			log.Println(err)
			return
		}

		id, err := service.AddProduct(product)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Can't create product",
			})
			log.Println(err)
			return
		}

		ctx.JSON(201, gin.H{
			"message": "Successfully create product",
			"id":      id,
		})
	}
}

func GetProducts(service *listing.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := ctx.Query("page")
		size := ctx.DefaultQuery("pageSize", "10")
		if p == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Page should provide in request",
			})
			return
		}

		page, err := strconv.Atoi(p)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Couldn't parse the incoming page no",
			})
			log.Println(err)
			return
		}
		pageSize, err := strconv.Atoi(size)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Couldn't parse the incoming page no",
			})
			log.Println(err)
			return
		}

		products, count, err := service.GetProducts(page, pageSize)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Couldn't retrieve products",
			})
			log.Println(err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Successfully retrieve products",
			"products": products,
			"total":    count,
		})

	}
}
