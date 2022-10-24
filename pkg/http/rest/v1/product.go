package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
	"github.com/thetnaingtn/go-dermacare-service/pkg/deleting"
	"github.com/thetnaingtn/go-dermacare-service/pkg/editing"
	"github.com/thetnaingtn/go-dermacare-service/pkg/listing"
	"github.com/thetnaingtn/go-dermacare-service/pkg/sys/validate"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProduct(service adding.Service) validate.Handler {
	return func(ctx *gin.Context) error {
		var product adding.Product

		if err := ctx.ShouldBind(&product); err != nil {
			if _, ok := err.(validator.ValidationErrors); ok {
				fieldErrors := validate.GetFieldsValidationErrors(err)
				log.Println(fieldErrors)
				return fieldErrors
			}

			log.Println(err)
			return validate.NewRequestError(validate.ErrInvalidPayload, http.StatusBadRequest)
		}

		id, err := service.AddProduct(product)
		if err != nil {
			log.Println(err)
			return err
		}

		ctx.JSON(201, gin.H{
			"message": "Successfully create product",
			"id":      id,
		})

		return nil

	}
}

func GetProduct(service listing.Service) validate.Handler {
	return func(ctx *gin.Context) error {
		id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
		if err != nil {
			log.Println(err)
			return validate.NewRequestError(validate.ErrInvalidId, http.StatusBadRequest)
		}

		product, err := service.GetProduct(id)
		if err != nil && err == mongo.ErrNoDocuments {
			return validate.NewRequestError(validate.ErrNotFound, http.StatusNotFound)
		}

		if err != nil {
			log.Println(err)
			return err
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Successfully retrieve product",
			"product": product,
		})

		return nil
	}
}

func GetProducts(service listing.Service) validate.Handler {
	return func(ctx *gin.Context) error {
		p := ctx.DefaultQuery("page", "1")
		size := ctx.DefaultQuery("pageSize", "10")

		page, err := strconv.Atoi(p)
		if err != nil {
			log.Println(err)
			return validate.NewRequestError(validate.ErrInvalidPayload, http.StatusBadRequest)
		}

		pageSize, err := strconv.Atoi(size)

		if err != nil {
			log.Println(err)
			return validate.NewRequestError(validate.ErrInvalidPayload, http.StatusBadRequest)
		}

		products, count, err := service.GetProducts(page, pageSize)
		if err != nil {
			log.Println(err)
			return err
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Successfully retrieve products",
			"products": products,
			"total":    count,
		})

		return nil
	}
}

func UpdateProduct(service editing.Service) validate.Handler {
	return func(ctx *gin.Context) error {
		var productEditForm editing.ProductEditForm
		productId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
		if err != nil {
			log.Println(err)
			return validate.NewRequestError(validate.ErrInvalidId, http.StatusBadRequest)
		}

		err = ctx.ShouldBind(&productEditForm)
		if err != nil {
			log.Println(err)
			return validate.NewRequestError(validate.ErrInvalidPayload, http.StatusBadRequest)
		}

		updatedProduct, err := service.UpdateProduct(productId, productEditForm)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return validate.NewRequestError(validate.ErrNotFound, http.StatusNotFound)
			}

			log.Println(err)
			return err
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Successfully update document",
			"product": updatedProduct,
		})

		return nil
	}
}

func DeleteProduct(service deleting.Service) validate.Handler {
	return func(ctx *gin.Context) error {
		id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
		if err != nil {
			log.Println(err)
			return validate.NewRequestError(validate.ErrInvalidId, http.StatusBadRequest)

		}

		product, err := service.DeleteProduct(id)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				log.Println(err)
				return validate.NewRequestError(validate.ErrNotFound, http.StatusNotFound)
			}

			log.Println(err)
			return err
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Successfully deleted product",
			"product": product,
		})

		return nil
	}
}
