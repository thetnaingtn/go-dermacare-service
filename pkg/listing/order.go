package listing

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	Name         string   `json:"name"`
	Price        int64    `json:"price"`
	SellingPrice int64    `json:"selling_price" bson:"selling_price"` // need to specify bson tag otherwise mongo-db driver will marshal the field as lowercase(sellingprice).
	Categories   []string `json:"categories"`
	Quantity     int      `json:"quantity"`
}

type Order struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	PhoneNo     string             `json:"phone_no,omitempty" bson:"phone_no,omitempty"`
	Items       []Item             `json:"items" bson:"items"`
	DeliverDate time.Time          `json:"deliver_date,omitempty" bson:"deliver_date,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
