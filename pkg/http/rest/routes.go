package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
	"github.com/thetnaingtn/go-dermacare-service/pkg/deleting"
	"github.com/thetnaingtn/go-dermacare-service/pkg/editing"
	v1 "github.com/thetnaingtn/go-dermacare-service/pkg/http/rest/v1"
	"github.com/thetnaingtn/go-dermacare-service/pkg/listing"
	"github.com/thetnaingtn/go-dermacare-service/pkg/middleware"
	"github.com/thetnaingtn/go-dermacare-service/pkg/sys/validate"
)

func InitializeRoute(a adding.Service, l listing.Service, e editing.Service, d deleting.Service) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Error())

	router.POST("/products", validate.ErrHandler(v1.AddProduct(a)))
	router.GET("/products", validate.ErrHandler(v1.GetProducts(l)))
	router.GET("/products/:id", validate.ErrHandler(v1.GetProduct(l)))
	router.PUT("/products/:id", validate.ErrHandler(v1.UpdateProduct(e)))
	router.DELETE("/products/:id", validate.ErrHandler(v1.DeleteProduct(d)))

	router.POST("/categories", v1.AddCategory(a))

	router.POST("/orders", v1.AddOrder(a))
	router.GET("/orders", v1.GetOrders(l))

	router.POST("/users", v1.AddUser(a))

	return router

}
