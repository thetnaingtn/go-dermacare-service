package category

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	DB *mongo.Database
}

func NewStore(db *mongo.Database) Store {
	return Store{DB: db}
}

var collection = "categories"

func (s Store) Create(nc NewCategory) (Category, error) {
	var category Category
	collection := s.DB.Collection(collection)
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
		return Category{}, fmt.Errorf("Error when insert new category: %w", err)
	}

	id, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return Category{}, fmt.Errorf("Inserted id can't assert primitive.ObjectId")
	}

	category.ID = id.Hex()

	return category, nil
}

func (s Store) Update(id primitive.ObjectID, uc UpdateCategory) (Category, error) {
	var category Category
	collection := s.DB.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	filterDoc := bson.M{"_id": id}
	updateDoc := bson.M{"$set": uc}
	options := options.FindOneAndUpdate().SetReturnDocument(1)

	res := collection.FindOneAndUpdate(ctx, filterDoc, updateDoc, options)

	if err := res.Decode(&category); err != nil {
		return Category{}, fmt.Errorf("Error occur when unmarshal updated result: %w", err)
	}

	return category, nil
}
