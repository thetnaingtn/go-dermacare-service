package core

import (
	"fmt"

	"github.com/thetnaingtn/go-dermacare-service/business/data/store/user"
)

type Repository interface {
	Signup(user.NewUser) (user.User, error)
}

type Core struct {
	Store Repository
}

func NewCore(s Repository) Core {
	return Core{
		Store: s,
	}
}

func (c Core) Create(nu user.NewUser) (user.User, error) {

	usr, err := c.Store.Signup(nu)
	if err != nil {
		return user.User{}, fmt.Errorf("create: %w", err)
	}

	return usr, nil
}
