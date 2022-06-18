package listing

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id           primitive.ObjectID   `json:"_id" bson:"_id"`
	Name         string               `json:"name" bson:"name"`
	Quantity     int                  `json:"quantity" bson:"quantity"`
	Price        int64                `json:"price" bson:"price"`
	SellingPrice int64                `json:"selling_price" bson:"selling_price"`
	Categories   []primitive.ObjectID `json:"categories,omitempty" bson:"categories,omitempty"`
	Description  string               `json:"description,omitempty" bson:"description,omitempty"`
	Supplier     primitive.ObjectID   `json:"supplier,omitempty" bson:"supplier,omitempty"`
	MinimumStock int                  `json:"miminum_quantity,omitempty" bson:"minimum_quantity,omitempty"`
	ExpiredDate  time.Time            `json:"expired_date,omitempty" bson:"expired_date,omitempty"`
	CreatedAt    time.Time            `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time            `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
