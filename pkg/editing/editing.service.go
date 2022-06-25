package editing

import (
	"github.com/thetnaingtn/go-dermacare-service/pkg/listing"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	UpdateProduct(id primitive.ObjectID, form ProductEditForm) (listing.Product, error)
}

type Service interface {
	UpdateProduct(id primitive.ObjectID, form ProductEditForm) (listing.Product, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateProduct(id primitive.ObjectID, form ProductEditForm) (listing.Product, error) {
	return s.r.UpdateProduct(id, form)
}
