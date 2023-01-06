package category

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	DB *mongo.Database
}

func NewStore(db *mongo.Database) Store {
	return Store{DB: db}
}

func (s Store) Create(nc NewCategory) (Category, error) {
	var category Category
	collection := s.DB.Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	category = Category{
		Name:        nc.Name,
		Description: nc.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	res, err := collection.InsertOne(ctx, category)

	if err != nil {
		log.Println(err)
		return Category{}, fmt.Errorf("Error when insert new category: %w", err)
	}

	id, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return Category{}, fmt.Errorf("Inserted id can't assert primitive.ObjectId")
	}

	category.ID = id.Hex()

	return category, nil
}
