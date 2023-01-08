package category

import (
	"fmt"
	"log"

	"github.com/thetnaingtn/go-dermacare-service/business/data/store/category"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Core struct {
	store category.Store
}

func NewCore(s category.Store) Core {
	return Core{store: s}
}

func (c Core) Create(nc category.NewCategory) (category.Category, error) {

	cat, err := c.store.Create(nc)
	if err != nil {
		return category.Category{}, err
	}

	return cat, nil
}

func (c Core) Update(id string, uc category.UpdateCategory) (category.Category, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return category.Category{}, fmt.Errorf("Error occur when object id created from hex: %w", err)
	}

	cat, err := c.store.Update(objId, uc)
	if err != nil {
		log.Println(err)
		return category.Category{}, err
	}

	return cat, nil
}
