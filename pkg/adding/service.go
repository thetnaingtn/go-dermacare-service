package adding

type Repository interface {
	AddProduct(payload Product) (string, error)
}

type Service interface {
	AddProduct(payload Product) (string, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddProduct(product Product) (string, error) {
	return s.r.AddProduct(product)
}
