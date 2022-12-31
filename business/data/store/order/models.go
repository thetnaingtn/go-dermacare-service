package order

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// order payload that client sent
type NewOrder struct {
	Name    string `json:"name,omitempty" binding:"required"`
	Address string `json:"address,omitempty" binding:"required"`
	PhoneNo string `json:"phone_no,omitempty" binding:"required"`
	Items   []struct {
		Id       string `json:"_id,omitempty" binding:"required"`
		Quantity int    `json:"quantity" binding:"required"`
	} `json:"items" binding:"required"`
	DeliverDate time.Time `json:"deliver_date,omitempty" binding:"required"`
}

type OrderItem struct {
	Id           string               `json:"_id"`
	Name         string               `json:"name"`
	Price        int                  `json:"price"`
	SellingPrice int                  `json:"selling_price" bson:"selling_price"`
	Categories   []primitive.ObjectID `json:"categories,omitempty" bson:"categories,omitempty"`
	Quantity     int                  `json:"quantity" bson:"quantity"`
}

// order type marshal to store in DB, unmarshal to send it back to client
type Order struct {
	Id          string      `json:"_id" bson:"_id,omitempty"`
	Name        string      `json:"name" bson:"name"`
	Address     string      `json:"address,omitempty" bson:"address,omitempty"`
	PhoneNo     string      `json:"phone_no,omitempty" bson:"phone_no,omitempty"`
	Items       []OrderItem `json:"items" bson:"items,omitempty"`
	DeliverDate time.Time   `json:"deliver_date,omitempty" bson:"deliver_date,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Orders struct {
	Orders []Order `json:"orders"`
	Total  int64   `json:"total"`
}
