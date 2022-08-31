package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(service adding.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user adding.User
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Can't parse request payload",
			})
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		user.Password = string(hash)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Can't hash password",
			})
			return
		}

		if err := service.AddUser(user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Can't create user",
			})
			return
		}

		log.Printf("%+v", user)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Successfully create user",
		})
	}
}
