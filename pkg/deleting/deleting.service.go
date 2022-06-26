package deleting

import "go.mongodb.org/mongo-driver/bson/primitive"

type Repository interface {
	DeleteProductById(id primitive.ObjectID) (Product, error)
}

type Service interface {
	DeleteProduct(id primitive.ObjectID) (Product, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteProduct(id primitive.ObjectID) (Product, error) {
	return s.r.DeleteProductById(id)
}
