package adding

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	AddProduct(payload Product) (string, error)
	AddCategory(payload Category) (string, error)
	AddOrder(payload Order) (string, error)
	GetProductByIds(ids []primitive.ObjectID, fields []string) (OrderItems, error)
}

type Service interface {
	AddProduct(payload Product) (string, error)
	AddCategory(payload Category) (string, error)
	AddOrder(payload OrderForm) (string, error)
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

func (s *service) AddOrder(orderForm OrderForm) (string, error) {
	var productIds []primitive.ObjectID
	for _, item := range orderForm.Items {
		objectId, _ := primitive.ObjectIDFromHex(item.Id)
		productIds = append(productIds, objectId)
	}
	products, err := s.r.GetProductByIds(productIds, []string{"name", "categories", "price", "selling_price"})

	if err != nil {
		return "", err
	}

	orderItems := products.AddQuantity(orderForm.Items)

	order := Order{
		Name:        orderForm.Name,
		Address:     orderForm.Address,
		PhoneNo:     orderForm.PhoneNo,
		DeliverDate: orderForm.DeliverDate,
		Items:       orderItems,
	}

	return s.r.AddOrder(order)
}
