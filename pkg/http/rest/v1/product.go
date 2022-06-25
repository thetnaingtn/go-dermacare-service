package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
	"github.com/thetnaingtn/go-dermacare-service/pkg/editing"
	"github.com/thetnaingtn/go-dermacare-service/pkg/listing"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProduct(service adding.Service) gin.HandlerFunc {
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

func GetProduct(service listing.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Can't process the product id",
			})
			log.Println(err)
			return
		}

		product, err := service.GetProduct(id)
		if err != nil && err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Document not found",
			})
			return
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Can't get the product",
			})
			log.Println(err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Successfully retrieve product",
			"product": product,
		})

	}
}

func GetProducts(service listing.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := ctx.DefaultQuery("page", "1")
		size := ctx.DefaultQuery("pageSize", "10")

		page, err := strconv.Atoi(p)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Couldn't parse the incoming page no",
			})
			log.Println(err)
			return
		}
		pageSize, err := strconv.Atoi(size)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Couldn't parse the incoming page size",
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

func UpdateProduct(service editing.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productEditForm editing.ProductEditForm
		productId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
		if err != nil {
			// TODO:need to check valid hex or not.
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Can't process the product id",
			})
			log.Println(err)
			return
		}
		err = ctx.ShouldBind(&productEditForm)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Can't process the request body",
			})
			log.Println(err)
			return
		}

		updatedProduct, err := service.UpdateProduct(productId, productEditForm)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "Document not found",
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Can't update the document",
			})
			log.Println(err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Successfully update document",
			"product": updatedProduct,
		})

	}
}
