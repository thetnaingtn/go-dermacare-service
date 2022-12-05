package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/auth"
	"github.com/thetnaingtn/go-dermacare-service/pkg/sys/validate"
	"go.mongodb.org/mongo-driver/bson"
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
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return User{}, fmt.Errorf("can't convert inserted id")
	}

	usr.Id = id.Hex()

	return usr, nil

}

func (s Store) Authenticate(email, password string) (auth.Claim, error) {
	var user User

	collection := s.db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	collection.FindOne(ctx, bson.M{"email": email}, nil).Decode(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return auth.Claim{}, validate.ErrIncorrectPassword
	}

	claim := auth.Claim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Dermacare service",
			Subject:   user.Id,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	return claim, nil

}
