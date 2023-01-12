package category

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thetnaingtn/go-dermacare-service/business/data/store/category"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/validate"
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

func (c Core) DeleteById(id string) (category.Category, error) {
	if !primitive.IsValidObjectID(id) {
		return category.Category{}, validate.NewRequestError(validate.ErrInvalidId, http.StatusBadRequest)
	}

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return category.Category{}, fmt.Errorf("Error occur when create new object id from hex: %w", err)
	}

	cat, err := c.store.DeleteById(objId)

	if err != nil {
		return category.Category{}, err
	}

	return cat, err

}

func (c Core) Query() (category.Categories, error) {
	categories, err := c.store.Query()

	if err != nil {
		return category.Categories{}, err
	}

	return categories, nil
}
