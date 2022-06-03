package mongo

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Database
}

func NewRepository(client *mongo.Client) *Repository {
	if err := godotenv.Load(); err != nil {
		panic("Can't load env variables")
	}
	db := client.Database(os.Getenv("DB_NAME"))
	return &Repository{db}
}

func (r *Repository) AddProduct(product adding.Product) string {
	productCollection := r.db.Collection("products")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	result, err := productCollection.InsertOne(ctx, product)
	if err != nil {
		panic(err)
	}

	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		return id.Hex()
	}

	return ""

}
