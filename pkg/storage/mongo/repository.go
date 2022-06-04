package mongo

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
	"github.com/thetnaingtn/go-dermacare-service/pkg/listing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *Repository) AddProduct(product adding.Product) (string, error) {
	collection := r.db.Collection("products")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		return "", err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", nil
	}

	return id.Hex(), nil

}

func (r *Repository) GetProducts(page int, pageSize int) ([]listing.Product, int64, error) {
	var products []listing.Product
	collection := r.db.Collection("products")
	skip := (page - 1) * pageSize
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	cursor, err := collection.Find(context.Background(), bson.D{}, options.Find().SetSkip(int64(skip)), options.Find().SetLimit(int64(pageSize)))
	if err != nil {
		log.Fatal(err)
		return nil, 0, err
	}

	if err = cursor.All(context.Background(), &products); err != nil {
		return nil, 0, err
	}

	return products, count, nil
}
