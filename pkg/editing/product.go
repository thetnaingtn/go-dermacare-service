package editing

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// need to add omitempty for bson. Otherwise, all the empty form payload are encoded into zero value(false, 0, "") and will update the db with those zero value.
type ProductEditForm struct {
	Name         string               `json:"name" bson:"name,omitempty"`
	Quantity     int                  `json:"quantity" bson:"quantity,omitempty"`
	Price        int64                `json:"price" bson:"price,omitempty"`
	SellingPrice int64                `json:"selling_price" bson:"selling_price,omitempty"`
	Categories   []primitive.ObjectID `json:"categories" bson:"categories,omitempty"`
	Description  string               `json:"description" bson:"description,omitempty"`
	Supplier     primitive.ObjectID   `json:"supplier" bson:"supplier,omitempty"`
	MinimumStock int                  `json:"miminum_stock" bson:"minimum_stock,omitempty"`
	ExpiredDate  time.Time            `json:"expired_date" bson:"expired_date,omitempty"`
	CreatedAt    time.Time            `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt    time.Time            `json:"updated_at" bson:"updated_at,omitempty"`
}
