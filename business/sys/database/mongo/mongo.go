package mongo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Setup() (*mongo.Client, func()) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DB_URI")))
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

	return client, fn
}

func CreateDatabase(client *mongo.Client) *mongo.Database {
	db := client.Database(os.Getenv("DB_NAME"))
	return db
}
