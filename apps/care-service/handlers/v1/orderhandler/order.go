package orderhandler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	core "github.com/thetnaingtn/go-dermacare-service/business/core/order"
	"github.com/thetnaingtn/go-dermacare-service/business/data/store/order"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/auth"
	"github.com/thetnaingtn/go-dermacare-service/pkg/sys/validate"
)

type Handlers struct {
	Auth *auth.Auth
	Core core.Core
}

func (h Handlers) Create(ctx *gin.Context) error {
	var no order.NewOrder
	if err := ctx.ShouldBind(&no); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			fieldErrors := validate.GetFieldsValidationErrors(err)
			log.Println("fields", fieldErrors)
			return fieldErrors
		}
	}

	order, err := h.Core.Create(no)
	if err != nil {
		log.Println(err)
		return err
	}

	ctx.JSON(http.StatusCreated, order)
	return err
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
