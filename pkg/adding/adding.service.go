package adding

import (
	"errors"
	"io"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	AddProduct(payload Product) (string, error)
	AddCategory(payload Category) (string, error)
	AddOrder(payload Order) (string, error)
	Signup(payload User) error
	GetProductByIds(ids []primitive.ObjectID, fields []string) (OrderItems, error)
	UpdateInStockProduct(items []Item) error
	GetUserByEmail(email string) (User, error)
}

type Service interface {
	AddProduct(payload Product) (string, error)
	AddCategory(payload Category) (string, error)
	AddOrder(payload OrderForm) (string, error)
	Signup(payload User) error
	Signin(payload User) error
	GenerateToken(payload User) (string, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

var ErrOutOfStock = errors.New("Out of Stock")

func (s *service) AddProduct(product Product) (string, error) {
	return s.r.AddProduct(product)
}

func (s *service) AddCategory(category Category) (string, error) {
	return s.r.AddCategory(category)
}

func (s *service) AddOrder(orderForm OrderForm) (string, error) {
	productIds := make([]primitive.ObjectID, len(orderForm.Items))
	for _, item := range orderForm.Items {
		objectId, _ := primitive.ObjectIDFromHex(item.Id)
		productIds = append(productIds, objectId)
	}
	// Get all interested products
	products, err := s.r.GetProductByIds(productIds, []string{"name", "categories", "price", "selling_price", "quantity"})

	if err != nil {
		return "", err
	}

	// check whether they are in stock
	inStockProducts := products.GetInStockProduct(orderForm.Items)
	var outOfStock []string

	for product, inStock := range inStockProducts {
		if !inStock {
			outOfStock = append(outOfStock, product)
		}
	}

	// if not return ErrOutOfStock
	if msg := strings.Join(outOfStock, ","); msg != "" {
		// TODO: add proper error msg to ErrOutOfStock
		return "", ErrOutOfStock
	}

	orderItems := products.AddQuantity(orderForm.Items)

	if err = s.r.UpdateInStockProduct(orderForm.Items); err != nil {
		return "", err
	}

	order := Order{
		Name:        orderForm.Name,
		Address:     orderForm.Address,
		PhoneNo:     orderForm.PhoneNo,
		DeliverDate: orderForm.DeliverDate,
		Items:       orderItems,
	}

	return s.r.AddOrder(order)
}

func (s *service) Signup(u User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hash)

	if err != nil {
		return err
	}

	return s.r.Signup(u)
}

func (s *service) Signin(u User) error {
	user, err := s.r.GetUserByEmail(u.Email)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GenerateToken(u User) (string, error) {
	file, err := os.Open("key/private.pem")
	if err != nil {
		return "", err
	}

	privatePEM, err := io.ReadAll(io.LimitReader(file, 1024*1024))
	if err != nil {
		return "", err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privatePEM))
	if err != nil {
		return "", err
	}

	claims := struct {
		jwt.RegisteredClaims
		User string
	}{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "subject",
			Issuer:    "derma care service",
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(8760 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
		User: u.Email,
	}

	method := jwt.GetSigningMethod("RS256")
	token := jwt.NewWithClaims(method, claims)

	tokenstr, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenstr, nil

}
