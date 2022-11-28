package user

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordNotMatch = errors.New("Passwords not match")
)

type Store struct {
	db *mongo.Database
}

func NewStore(db *mongo.Database) Store {
	return Store{db}
}

func (s Store) Signup(u NewUser) (User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return User{}, err
	}

	collection := s.db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	usr := User{
		Name:      u.Name,
		Password:  string(hash),
		Email:     u.Email,
		Roles:     u.Roles,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := collection.InsertOne(ctx, usr)
	if err != nil {
		return User{}, err
	}
	id := res.InsertedID.(primitive.ObjectID)
	usr.Id = id.String()

	return usr, nil

}
