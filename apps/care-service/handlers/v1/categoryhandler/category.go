package categoryhandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	categoryCore "github.com/thetnaingtn/go-dermacare-service/business/core/category"
	"github.com/thetnaingtn/go-dermacare-service/business/data/store/category"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/validate"
)

type Handler struct {
	Core categoryCore.Core
}

func (h Handler) Create(ctx *gin.Context) error {
	var nc category.NewCategory

	err := ctx.ShouldBind(&nc)

	_, ok := err.(validator.ValidationErrors)
	if ok {
		fieldErrors := validate.GetFieldsValidationErrors(err)
		log.Println(fieldErrors)
		return fieldErrors
	}

	if err != nil {
		log.Println(err)
		return validate.NewRequestError(err, http.StatusBadRequest)
	}

	category, err := h.Core.Create(nc)

	if err != nil {
		return err
	}

	ctx.JSON(http.StatusCreated, category)

	return nil
}
