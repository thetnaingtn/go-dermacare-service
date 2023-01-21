package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBConfig struct {
	Host string
	Name string
}

func CreateDatabase(cfg DBConfig) (*mongo.Database, func()) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", cfg.Host)))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	fn := func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}

	db := client.Database(cfg.Name)
	return db, fn
}
