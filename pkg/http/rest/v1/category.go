package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
)

func AddCategory(service adding.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var category adding.Category
		if err := ctx.ShouldBind(&category); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Can't parse incoming request data",
			})
			log.Println(err)
			return
		}
		id, err := service.AddCategory(category)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Can't create category",
			})
			log.Println(err)
			return
		}
		ctx.JSON(201, gin.H{
			"message": "Successfully create category",
			"id":      id,
		})
	}
}
