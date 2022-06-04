package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
	v1 "github.com/thetnaingtn/go-dermacare-service/pkg/http/rest/v1"
	"github.com/thetnaingtn/go-dermacare-service/pkg/listing"
)

func InitializeRoute(a *adding.Service, l *listing.Service) *gin.Engine {
	router := gin.Default()

	router.POST("/products", v1.AddProduct(a))
	router.GET("/products", v1.GetProducts(l))

	return router

}
