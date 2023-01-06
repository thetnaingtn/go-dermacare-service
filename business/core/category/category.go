package category

import "github.com/thetnaingtn/go-dermacare-service/business/data/store/category"

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
