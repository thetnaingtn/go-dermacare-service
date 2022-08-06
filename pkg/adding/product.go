package adding

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Name         string               `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	Quantity     int                  `json:"quantity,omitempty" bson:"quantity,omitempty" binding:"required"`
	Price        int64                `json:"price,omitempty" bson:"price,omitempty" binding:"required"`
	SellingPrice int64                `json:"selling_price,omitempty" bson:"selling_price,omitempty" binding:"required"`
	Categories   []primitive.ObjectID `json:"categories,omitempty" bson:"categories,omitempty"`
	Description  string               `json:"description,omitempty" bson:"description,omitempty"`
	Supplier     primitive.ObjectID   `json:"supplier,omitempty" bson:"supplier,omitempty"`
	MinimumStock int                  `json:"minimum_stock,omitempty" bson:"minimum_stock,omitempty" binding:"required"`
	ExpiredDate  time.Time            `json:"expired_date,omitempty" bson:"expired_date,omitempty"`
	CreatedAt    time.Time            `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time            `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
