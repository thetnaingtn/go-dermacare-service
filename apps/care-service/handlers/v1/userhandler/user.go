package userhandler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	user "github.com/thetnaingtn/go-dermacare-service/business/core/user"
	userstore "github.com/thetnaingtn/go-dermacare-service/business/data/store/user"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/auth"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/validate"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handlers struct {
	Auth *auth.Auth
	Core user.Core
}

func (h Handlers) Signup(ctx *gin.Context) error {
	var user userstore.NewUser
	if err := ctx.ShouldBind(&user); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			fieldErrors := validate.GetFieldsValidationErrors(err)
			log.Println("fields", fieldErrors)
			return fieldErrors
		}

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

func (h Handlers) Signin(ctx *gin.Context) error {

	var credential struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBind(&credential); err != nil {
		log.Println(err)
		return err
	}

	if credential.Email == "" || credential.Password == "" {
		return validate.NewRequestError(validate.ErrInvalidPayload, http.StatusUnauthorized)
	}

	claim, err := h.Core.Authenticate(credential.Email, credential.Password)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			er := errors.New("User not found")
			return validate.NewRequestError(er, http.StatusNotFound)
		}
		return err
	}

	token, err := h.Auth.GenerateToken(claim)
	if err != nil {
		log.Println(err)
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

	return nil
}
