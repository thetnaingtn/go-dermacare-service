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

func (h Handler) Update(ctx *gin.Context) error {
	var uc category.UpdateCategory
	err := ctx.ShouldBind(&uc)
	if err != nil {
		log.Println(err)
		return validate.NewRequestError(err, http.StatusBadRequest)
	}

	id := ctx.Param("id")

	category, err := h.Core.Update(id, uc)
	if err != nil {
		log.Println(err)
		return err
	}

	ctx.JSON(http.StatusOK, category)

	return nil
}

func (h Handler) DeleteById(ctx *gin.Context) error {
	id := ctx.Param("id")
	c, err := h.Core.DeleteById(id)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, c)

	return nil
}

func (h Handler) Query(ctx *gin.Context) error {
	categories, err := h.Core.Query()
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, categories)
	return nil
}
