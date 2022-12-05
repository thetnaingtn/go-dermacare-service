package core

import (
	"fmt"

	"github.com/thetnaingtn/go-dermacare-service/business/data/store/user"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/auth"
)

type Repository interface {
	Signup(user.NewUser) (user.User, error)
	Authenticate(email, passowrd string) (auth.Claim, error)
}

type Core struct {
	store Repository
}

func NewCore(s Repository) Core {
	return Core{
		store: s,
	}
}

func (c Core) Create(nu user.NewUser) (user.User, error) {

	usr, err := c.store.Signup(nu)
	if err != nil {
		return user.User{}, fmt.Errorf("create: %w", err)
	}

	return usr, nil
}

func (c Core) Authenticate(email, password string) (auth.Claim, error) {
	claim, err := c.store.Authenticate(email, password)
	if err != nil {
		return auth.Claim{}, err
	}

	return claim, nil
}
