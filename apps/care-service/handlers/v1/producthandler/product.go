package producthandler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	core "github.com/thetnaingtn/go-dermacare-service/business/core/product"
	"github.com/thetnaingtn/go-dermacare-service/business/data/store/product"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/auth"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/validate"
)

type Handlers struct {
	Auth *auth.Auth
	Core core.Core
}

func (h Handlers) Create(ctx *gin.Context) error {
	var np product.NewProduct

	if err := ctx.ShouldBind(&np); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			fieldErrors := validate.GetFieldsValidationErrors(err)
			log.Println("fields", fieldErrors)
			return fieldErrors
		}

		log.Println(err)
		return validate.NewRequestError(validate.ErrInvalidPayload, http.StatusBadRequest)
	}

	p, err := h.Core.Create(np)
	if err != nil {
		log.Println(err)
		return err
	}

	ctx.JSON(http.StatusCreated, p)

	return nil
}

func (h Handlers) Update(ctx *gin.Context) error {
	var up product.UpdateProduct
	if err := ctx.ShouldBind(&up); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			fieldErrors := validate.GetFieldsValidationErrors(err)
			log.Println("fields", fieldErrors)
			return fieldErrors
		}

		log.Println(err)
		return validate.NewRequestError(validate.ErrInvalidPayload, http.StatusBadRequest)
	}

	id := ctx.Param("id")

	p, err := h.Core.Update(up, id)
	if err != nil {
		log.Println(err)
		return err
	}

	ctx.JSON(http.StatusOK, p)

	return nil
}

func (h Handlers) Query(ctx *gin.Context) error {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		log.Println(err)
		return validate.NewRequestError(fmt.Errorf("page doesn't have a proper value"), http.StatusBadRequest)
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "15"))
	if err != nil {
		log.Println(err)
		return validate.NewRequestError(fmt.Errorf("pageSize doesn't have a proper value"), http.StatusBadRequest)
	}

	p, err := h.Core.Query(page, pageSize)

	if err != nil {
		log.Println(err)
		return err
	}

	ctx.JSON(http.StatusOK, p)

	return nil
}
