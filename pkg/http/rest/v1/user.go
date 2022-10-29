package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
	"github.com/thetnaingtn/go-dermacare-service/pkg/sys/validate"
)

func Signup(service adding.Service) validate.Handler {
	return func(ctx *gin.Context) error {
		var user adding.User
		if err := ctx.ShouldBind(&user); err != nil {
			log.Println(err)
			return validate.NewRequestError(validate.ErrInvalidPayload, http.StatusBadRequest)
		}

		if err := service.Signup(user); err != nil {
			log.Println(err)
			return err
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Successfully create user",
		})

		return nil
	}
}
