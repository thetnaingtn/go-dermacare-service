package order

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func (s Store) Query(page, pageSize int) (Orders, error) {
	o := []Order{}

	collection := s.DB.Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	total, _ := collection.CountDocuments(ctx, bson.D{})
	limit := bson.D{{"$limit", pageSize}}
	skip := bson.D{{"$skip", (page - 1) * pageSize}}

	lookup := bson.D{{
		"$lookup", bson.D{
			{"from", "categories"},
			{"localField", "items.categories"},
			{"foreignField", "_id"},
			{"as", "categories"},
		},
	}}

	unwind := bson.D{{"$unwind", "$items"}}

	addDefaultCategories := bson.D{{"$addFields", bson.D{
		{"items.categories", bson.D{{
			"$cond", bson.D{{
				"if", bson.D{{
					"$ne", bson.A{bson.D{{"$type", "$items.categories"}}, "array"}}},
			}, {"then", bson.A{}}, {"else", "$items.categories"}},
		}}},
	}}}

	addField := bson.D{{"$addFields", bson.D{
		{"items.categories", bson.D{{
			"$map", bson.D{{
				"input", bson.D{
					{"$filter", bson.D{
						{"input", "$categories"},
						{"as", "category"},
						{"cond", bson.D{{
							"$in", bson.A{"$$category._id", "$items.categories"},
						}}},
					}},
				},
			}, {"as", "category"}, {"in", "$$category.name"}},
		}}},
	}}}

	group := bson.D{{"$group", bson.D{
		{"_id", "$_id"},
		{"name", bson.D{{"$first", "$name"}}},
		{"address", bson.D{{"$first", "$address"}}},
		{"phone_no", bson.D{{"$first", "$phone_no"}}},
		{"items", bson.D{{"$push", "$items"}}},
		{"deliver_date", bson.D{{"$first", "$deliver_date"}}},
		{"created_at", bson.D{{"$first", "$created_at"}}},
	}}}

	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{lookup, unwind, addDefaultCategories, addField, group, skip, limit})
	if err != nil {
		log.Println(err)
		return Orders{}, err
	}

	if err := cursor.All(ctx, &o); err != nil {
		log.Println(err)
		return Orders{}, fmt.Errorf("Error when unmarshal all orders %w", err)
	}

	orders := Orders{
		Orders: o,
		Total:  total,
	}

	return orders, nil
}
