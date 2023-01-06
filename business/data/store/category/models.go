package category

import "time"

type Category struct {
	ID          string    `json:"_id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type NewCategory struct {
	Name        string `json:"name" bson:"name" binding:"required"`
	Description string `json:"description" bson:"description" binding:"required"`
}
