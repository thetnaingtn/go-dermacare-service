package producthandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	core "github.com/thetnaingtn/go-dermacare-service/business/core/product"
	"github.com/thetnaingtn/go-dermacare-service/business/data/store/product"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/auth"
	"github.com/thetnaingtn/go-dermacare-service/pkg/sys/validate"
)

type Handlers struct {
	Auth *auth.Auth
	Core core.Core
}

func (h Handlers) Create(ctx *gin.Context) error {
	var np product.NewProduct

	if err := ctx.ShouldBind(&np); err != nil {
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
