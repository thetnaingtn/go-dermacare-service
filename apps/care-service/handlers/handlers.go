package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/apps/care-service/handlers/v1/producthandler"
	"github.com/thetnaingtn/go-dermacare-service/apps/care-service/handlers/v1/userhandler"
	productcore "github.com/thetnaingtn/go-dermacare-service/business/core/product"
	usercore "github.com/thetnaingtn/go-dermacare-service/business/core/user"
	"github.com/thetnaingtn/go-dermacare-service/business/data/store/product"
	"github.com/thetnaingtn/go-dermacare-service/business/data/store/user"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/auth"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/validate"
	"github.com/thetnaingtn/go-dermacare-service/business/web/middleware"

	"go.mongodb.org/mongo-driver/mongo"
)

type APIConfig struct {
	Auth *auth.Auth
	DB   *mongo.Database
}

func InitializeRoute(cfg APIConfig) *gin.Engine {

	router := gin.Default()
	router.Use(middleware.Error())
	// user
	userHandler := userhandler.Handlers{
		Auth: cfg.Auth,
		Core: usercore.NewCore(user.NewStore(cfg.DB)),
	}
	router.POST("/signup", validate.ErrHandler(userHandler.Signup))
	router.POST("/signin", validate.ErrHandler(userHandler.Signin))

	// product
	productHandler := producthandler.Handlers{
		Auth: cfg.Auth,
		Core: productcore.NewCore(product.NewStore(cfg.DB)),
	}

	proutes := router.Group("/products")
	proutes.Use(middleware.Authenticate(cfg.Auth))
	{
		proutes.POST("", validate.ErrHandler(productHandler.Create))
		proutes.GET("", validate.ErrHandler(productHandler.Query))
		proutes.PUT("/:id", validate.ErrHandler(productHandler.Update))
	}
	return router
}
