package listing

type Repository interface {
	GetProducts(page, pageSize int) ([]Product, int64, error)
}

type Service interface {
	GetProducts(page, pageSize int) ([]Product, int64, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetProducts(page, pageSize int) ([]Product, int64, error) {
	products, count, err := s.r.GetProducts(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return products, count, nil
}
