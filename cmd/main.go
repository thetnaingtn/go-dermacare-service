package main

import (
	"context"
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
	"github.com/thetnaingtn/go-dermacare-service/pkg/deleting"
	"github.com/thetnaingtn/go-dermacare-service/pkg/editing"
	routes "github.com/thetnaingtn/go-dermacare-service/pkg/http/rest"
	"github.com/thetnaingtn/go-dermacare-service/pkg/listing"
	"github.com/thetnaingtn/go-dermacare-service/pkg/storage/mongo"
)

func main() {
	client := mongo.Setup()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	repository := mongo.NewRepository(client)
	addingService := adding.NewService(repository)
	listingService := listing.NewService(repository)
	editingService := editing.NewService(repository)
	deletingService := deleting.NewService(repository)

	router := routes.InitializeRoute(addingService, listingService, editingService, deletingService)

	log.Fatalln(router.Run(":3000"))

}
