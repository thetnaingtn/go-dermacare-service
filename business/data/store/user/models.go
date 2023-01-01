package user

import "time"

type User struct {
	Id        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type NewUser struct {
	Name            string   `json:"name" binding:"required"`
	Email           string   `json:"email" binding:"required,email"`
	Password        string   `json:"password" binding:"required"`
	PasswordConfirm string   `json:"password_confirm" binding:"required,eqfield=Password"`
	Roles           []string `json:"roles"`
}
