package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
	v1 "github.com/thetnaingtn/go-dermacare-service/pkg/http/rest/v1"
)

func InitializeRoute(a *adding.Service) *gin.Engine {
	router := gin.Default()

	productRoutes := router.Group("/product")
	{
		productRoutes.POST("/add", v1.AddProduct(a))
	}

	return router

}
