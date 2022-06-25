package listing

import "go.mongodb.org/mongo-driver/bson/primitive"

type Repository interface {
	GetProducts(page, pageSize int) ([]Product, int64, error)
	GetProductById(id primitive.ObjectID) (Product, error)
	GetOrders() ([]Order, error)
}

type Service interface {
	GetProducts(page, pageSize int) ([]Product, int64, error)
	GetProduct(id primitive.ObjectID) (Product, error)
	GetOrders() ([]Order, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetProduct(id primitive.ObjectID) (Product, error) {
	return s.r.GetProductById(id)
}

func (s *service) GetProducts(page, pageSize int) ([]Product, int64, error) {
	products, count, err := s.r.GetProducts(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return products, count, nil
}

func (s *service) GetOrders() ([]Order, error) {
	return s.r.GetOrders()
}
