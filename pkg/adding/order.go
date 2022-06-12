package adding

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	Id       string `json:"_id"`
	Quantity int    `json:"quantity"`
}
type OrderForm struct {
	Name        string    `json:"name"`
	Address     string    `json:"address,omitempty"`
	PhoneNo     int       `json:"phone_no,omitempty"`
	Items       []Item    `json:"items"`
	DeliverDate time.Time `json:"deliver_date,omitempty"`
}
type OrderItem struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Price        int64              `json:"price" bson:"price"`
	SellingPrice int64              `json:"selling_price" bson:"selling_price"`
	Category     primitive.ObjectID `json:"category,omitempty" bson:"category,omitempty"`
	Quantity     int                `json:"quantity" bson:"quantity"`
}
type Order struct {
	Name        string     `json:"name" bson:"name"`
	Address     string     `json:"address,omitempty" bson:"address,omitempty"`
	PhoneNo     int        `json:"phone_no,omitempty" bson:"phone_no,omitempty"`
	Items       OrderItems `json:"items" bson:"items"`
	DeliverDate time.Time  `json:"deliver_date,omitempty" bson:"deliver_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type OrderItems []OrderItem

func (orders OrderItems) AddQuantity(items []Item) (orderItems OrderItems) {
	for _, order := range orders {
		for _, item := range items {
			if order.Id.Hex() == item.Id {
				order.Quantity = item.Quantity
				orderItems = append(orderItems, order)
			}
		}
	}

	return orderItems
}
