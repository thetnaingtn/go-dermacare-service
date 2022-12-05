package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/apps/care-service/handlers/v1/userhandler"
	usercore "github.com/thetnaingtn/go-dermacare-service/business/core/user"
	"github.com/thetnaingtn/go-dermacare-service/business/data/store/user"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/auth"
	"github.com/thetnaingtn/go-dermacare-service/pkg/sys/validate"

	"go.mongodb.org/mongo-driver/mongo"
)

type APIConfig struct {
	Auth *auth.Auth
	DB   *mongo.Database
}

func InitializeRoute(cfg APIConfig) *gin.Engine {

	router := gin.Default()

	// user
	userHandler := userhandler.Handlers{
		Auth: cfg.Auth,
		Core: usercore.NewCore(user.NewStore(cfg.DB)),
	}
	router.POST("/signup", validate.ErrHandler(userHandler.Signup))
	router.POST("/signin", validate.ErrHandler(userHandler.Signin))

	return router
}
