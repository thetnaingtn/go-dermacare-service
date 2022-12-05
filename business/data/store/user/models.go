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
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Password        string   `json:"password"`
	PasswordConfirm string   `json:"password_confirm"`
	Roles           []string `json:"roles"`
}
