package main

import (
	"github.com/thetnaingtn/go-dermacare-service/apps/care-service/handlers"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/auth"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/database/mongo"
)

func main() {
	client, fn := mongo.Setup()
	defer fn()

	db := mongo.CreateDatabase(client)

	auth, _ := auth.New()
	apiConfig := handlers.APIConfig{
		Auth: auth,
		DB:   db,
	}

	engine := handlers.InitializeRoute(apiConfig)

	engine.Run(":3000")

}
