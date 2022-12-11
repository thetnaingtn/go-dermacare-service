package product

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	DB *mongo.Database
}

func NewStore(db *mongo.Database) Store {
	return Store{db}
}

func (s Store) Create(np NewProduct) (Product, error) {
	collection := s.DB.Collection("products")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	product := Product{
		Name:         np.Name,
		Quantity:     np.Quantity,
		Price:        np.SellingPrice,
		SellingPrice: np.SellingPrice,
		Categories:   np.Categories,
		Description:  np.Description,
		Supplier:     np.Supplier,
		MinimumStock: np.MinimumStock,
		ExpiredDate:  np.ExpiredDate,
	}

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	result, err := collection.InsertOne(ctx, product)

	if err != nil {
		return Product{}, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return Product{}, fmt.Errorf("can't convert inserted id")
	}

	product.Id = id.Hex()

	return product, nil
}
