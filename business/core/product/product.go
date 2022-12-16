package core

import (
	"errors"
	"log"
	"net/http"

	"github.com/thetnaingtn/go-dermacare-service/business/data/store/product"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/validate"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidId = errors.New("id is invalid or improper form")
)

type Repository interface {
	Create(np product.NewProduct) (product.Product, error)
	Update(up product.UpdateProduct, id primitive.ObjectID) (product.Product, error)
	DeleteById(id primitive.ObjectID) (product.Product, error)
	Query(page, pageSize int) (product.Products, error)
	QueryById(id primitive.ObjectID) (product.Product, error)
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

func (c Core) QueryById(id string) (product.Product, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return product.Product{}, validate.NewRequestError(ErrInvalidId, http.StatusBadRequest)
	}
	p, err := c.store.QueryById(objectId)

	if err != nil {
		log.Println(err)
		return product.Product{}, err
	}

	return p, nil
}

func (c Core) DeleteById(id string) (product.Product, error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return product.Product{}, validate.NewRequestError(ErrInvalidId, http.StatusBadRequest)
	}

	p, err := c.store.DeleteById(objectID)

	if err != nil {
		log.Println(err)
		return product.Product{}, err
	}

	return p, nil
}
