package userhandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	user "github.com/thetnaingtn/go-dermacare-service/business/core/user"
	userstore "github.com/thetnaingtn/go-dermacare-service/business/data/store/user"
	"github.com/thetnaingtn/go-dermacare-service/pkg/sys/validate"
)

type Handlers struct {
	Core user.Core
}

func (h Handlers) Signup(ctx *gin.Context) error {
	var user userstore.NewUser
	if err := ctx.ShouldBind(&user); err != nil {
		return validate.NewRequestError(validate.ErrInvalidPayload, http.StatusBadRequest)
	}

	usr, err := h.Core.Create(user)

	if err != nil {
		log.Println(err)
		return err
	}

	ctx.JSON(http.StatusOK, usr)

	return nil
}
