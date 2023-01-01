package main

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/thetnaingtn/go-dermacare-service/apps/care-service/handlers"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/auth"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/database/mongo"
	"github.com/thetnaingtn/go-dermacare-service/foundation/web"
)

func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(web.ValidatorTagNameFunc)
	}

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
