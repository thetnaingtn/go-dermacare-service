package adding

type Repository interface {
	AddProduct(payload Product) (string, error)
}

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) AddProduct(product Product) (string, error) {
	return s.r.AddProduct(product)
}
