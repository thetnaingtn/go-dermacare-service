package core

import (
	"log"

	"github.com/thetnaingtn/go-dermacare-service/business/data/store/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Create(np product.NewProduct) (product.Product, error)
	Update(up product.UpdateProduct, id primitive.ObjectID) (product.Product, error)
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

func (c Core) Update(up product.UpdateProduct, id string) (product.Product, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return product.Product{}, err
	}

	p, err := c.store.Update(up, objId)

	if err != nil {
		log.Println(err)
		return product.Product{}, err
	}

	return p, nil
}
