package adding

type Repository interface {
	AddProduct(payload Product) (string, error)
	AddCategory(payload Category) (string, error)
	AddOrder(payload Order) (string, error)
}

type Service interface {
	AddProduct(payload Product) (string, error)
	AddCategory(payload Category) (string, error)
	AddOrder(payload Order) (string, error)
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

func (s *service) AddCategory(category Category) (string, error) {
	return s.r.AddCategory(category)
}

func (s *service) AddOrder(order Order) (string, error) {
	return s.r.AddOrder(order)
}
