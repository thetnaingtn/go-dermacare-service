package main

import (
	"context"
	"log"

	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
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

	repository := mongo.NewRepository(client)
	addingService := adding.NewService(repository)
	listingService := listing.NewService(repository)

	router := routes.InitializeRoute(addingService, listingService)

	log.Fatalln(router.Run(":3000"))

}
