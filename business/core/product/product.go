package core

import (
	"log"

	"github.com/thetnaingtn/go-dermacare-service/business/data/store/product"
)

type Repository interface {
	Create(np product.NewProduct) (product.Product, error)
}

type Core struct {
	store Repository
}

func NewCore(s Repository) Core {
	return Core{store: s}
}

func (c Core) Create(np product.NewProduct) (product.Product, error) {

	p, err := c.store.Create(np)
	if err != nil {
		log.Println(err)
		return product.Product{}, err
	}

	return p, nil
}
