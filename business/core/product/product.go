package core

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thetnaingtn/go-dermacare-service/business/data/store/product"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/validate"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Create(np product.NewProduct) (product.Product, error)
	Update(up product.UpdateProduct, id primitive.ObjectID) (product.Product, error)
	Delete(id primitive.ObjectID) (product.Product, error)
	Query(page, pageSize int) (product.Products, error)
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

func (c Core) Query(page, pageSize int) (product.Products, error) {
	products, err := c.store.Query(page, pageSize)
	if err != nil {
		return product.Products{}, err
	}

	return products, nil
}

func (c Core) Delete(id string) (product.Product, error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return product.Product{}, validate.NewRequestError(fmt.Errorf("id doesn't have proper form"), http.StatusBadRequest)
	}

	p, err := c.store.Delete(objectID)

	if err != nil {
		log.Println(err)
		return product.Product{}, err
	}

	return p, nil
}
