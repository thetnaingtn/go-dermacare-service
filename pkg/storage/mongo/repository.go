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
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
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

func (r *Repository) GetProductByIds(ids []primitive.ObjectID) (products adding.OrderItems, err error) {
	collection := r.db.Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"_id": bson.M{"$in": ids}}, options.Find().SetProjection(bson.M{"name": 1, "price": 1, "selling_price": 1, "category": 1}))
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &products)

	if err != nil {
		return nil, err
	}

	return products, err

}

func (r *Repository) AddCategory(category adding.Category) (string, error) {
	collection := r.db.Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	result, err := collection.InsertOne(ctx, category)
	if err != nil {
		return "", err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", nil
	}

	return id.Hex(), nil
}

func (r *Repository) AddOrder(order adding.Order) (string, error) {
	collection := r.db.Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		return "", err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", nil
	}

	return id.Hex(), nil
}
