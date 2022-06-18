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

func (r *Repository) GetProducts(page int, pageSize int) (products []listing.Product, count int64, err error) {
	collection := r.db.Collection("products")
	skip := (page - 1) * pageSize
	count, err = collection.CountDocuments(context.Background(), bson.D{})
	addFieldsStage := bson.D{{"$addFields", bson.D{{"critical", bson.D{{"$lte", bson.A{"$quantity", "$minimum_stock"}}}}}}}
	limitStage := bson.D{{"$limit", int64(pageSize)}}
	skipStage := bson.D{{"$skip", int64(skip)}}
	cursor, err := collection.Aggregate(context.Background(), mongo.Pipeline{addFieldsStage, skipStage, limitStage})
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

	cursor, err := collection.Find(ctx, bson.M{"_id": bson.M{"$in": ids}}, options.Find().SetProjection(bson.M{"name": 1, "price": 1, "selling_price": 1, "categories": 1}))
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

func (r *Repository) GetOrders() (orders []listing.Order, err error) {
	collection := r.db.Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	lookupStage := bson.D{{
		"$lookup", bson.D{
			{"from", "categories"},
			{"localField", "items.categories"},
			{"foreignField", "_id"},
			{"as", "categories"},
		},
	}}

	unwindStage := bson.D{{"$unwind", "$items"}}

	addFieldStage := bson.D{{"$addFields", bson.D{
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

	groupStage := bson.D{{"$group", bson.D{
		{"_id", "$_id"},
		{"name", bson.D{{"$first", "$name"}}},
		{"address", bson.D{{"$first", "$address"}}},
		{"phone_no", bson.D{{"$first", "$phone_no"}}},
		{"items", bson.D{{"$push", "$items"}}},
		{"deliver_date", bson.D{{"$first", "$deliver_date"}}},
		{"created_at", bson.D{{"$first", "$created_at"}}},
	}}}

	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage, addFieldStage, groupStage})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}
