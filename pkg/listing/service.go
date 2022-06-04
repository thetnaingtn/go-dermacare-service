package listing

type Repository interface {
	GetProducts(page, pageSize int) ([]Product, int64, error)
}

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) GetProducts(page, pageSize int) ([]Product, int64, error) {
	products, count, err := s.r.GetProducts(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return products, count, nil
}
