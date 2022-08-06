package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var productSchema = bson.M{
	"bsonType": "object",
	"required": []string{"name", "quantity", "price", "selling_price", "minimum_stock"},
	"properties": bson.M{
		"name": bson.M{
			"bsonType":    "string",
			"description": "product name is required",
		},
		"quantity": bson.M{
			"bsonType":    "int",
			"description": "quantity is required",
		},
		"price": bson.M{
			"bsonType":    "long",
			"description": "price is required",
		},
		"selling_price": bson.M{
			"bsonType":    "long",
			"description": "selling price is required",
		},
		"minimum_stock": bson.M{
			"bsonType":    "int",
			"description": "minimum stock is required",
		},
	},
}

func getSchemaValidation(name string) *options.CreateCollectionOptions {
	var schema primitive.M
	if name == "products" {
		schema = productSchema
	}

	validator := bson.M{
		"$jsonSchema": schema,
	}

	return options.CreateCollection().SetValidator(validator)
}
