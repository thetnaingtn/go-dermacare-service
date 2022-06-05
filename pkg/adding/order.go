package adding

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Price        int64              `json:"price" bson:"price"`
	SellingPrice int64              `json:"selling_price" bson:"selling_price"`
	Quantity     int                `json:"quantity" bson:"quantity"`
}

type Order struct {
	Name        string    `json:"name" bson:"name"`
	Address     string    `json:"address,omitempty" bson:"address,omitempty"`
	PhoneNo     int       `json:"phone_no,omitempty" bson:"phone_no,omitempty"`
	Items       []Item    `json:"items" bson:"items"`
	DeliverDate time.Time `json:"deliver_date" bson:"deliver_date,omitempty"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
