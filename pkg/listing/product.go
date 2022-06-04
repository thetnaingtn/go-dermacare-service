package listing

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id           primitive.ObjectID `json:"_id"`
	Name         string             `json:"name"`
	Quantity     int                `json:"quantity"`
	Price        int64              `json:"price"`
	SellingPrice int64              `json:"selling_price"`
	Category     primitive.ObjectID `json:"category"`
	Description  string             `json:"description,omitempty"`
	Supplier     primitive.ObjectID `json:"supplier,omitempty"`
	ExpiredDate  time.Time          `json:"expired_date,omitempty"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}
