package order

import (
	"context"
	"log"
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

func (s Store) Create(order Order) (Order, error) {
	collection := s.DB.Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := collection.InsertOne(ctx, order)

	if err != nil {
		log.Println(err)
		return Order{}, err
	}

	id, _ := res.InsertedID.(primitive.ObjectID)

	order.Id = id.Hex()

	return order, nil
}
