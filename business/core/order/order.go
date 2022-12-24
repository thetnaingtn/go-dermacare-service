package order

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/thetnaingtn/go-dermacare-service/business/data/store/order"
	"github.com/thetnaingtn/go-dermacare-service/business/data/store/product"
	"github.com/thetnaingtn/go-dermacare-service/business/sys/validate"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Core struct {
	product product.Store
	order   order.Store
}

func NewCore(o order.Store, p product.Store) Core {
	return Core{product: p, order: o}
}

func (c Core) Create(no order.NewOrder) (order.Order, error) {
	//product ids
	pids := make([]primitive.ObjectID, 0, len(no.Items))
	for _, item := range no.Items {
		id, _ := primitive.ObjectIDFromHex(item.Id)
		pids = append(pids, id)
	}

	products, _ := c.product.QueryByIds(pids, []string{"name", "categories", "price", "selling_price", "quantity"})

	productQtyMap := make(map[string]int)
	for _, item := range no.Items {
		productQtyMap[item.Id] = item.Quantity
	}

	outStockProducts := product.GetOutOfStockProduct(products, productQtyMap)
	if outStockProducts != nil {
		err := validate.NewRequestError(fmt.Errorf("Out of stock product: %s", strings.Join(outStockProducts, ",")), http.StatusConflict)
		return order.Order{}, err
	}

	// replace product's quantity with requested item's quantity to update instock product.
	for i := range products {
		products[i].Quantity = productQtyMap[products[i].Id]
	}

	if err := c.product.UpdateInStockProduct(products); err != nil {
		log.Println(err)
		return order.Order{}, err
	}

	orderItems := make([]order.OrderItem, 0, len(no.Items))

	for _, p := range products {
		item := order.OrderItem{
			Id:           p.Id,
			Name:         p.Name,
			Price:        p.Price,
			SellingPrice: p.SellingPrice,
			Categories:   p.Categories,
			Quantity:     p.Quantity,
		}

		orderItems = append(orderItems, item)
	}

	o := order.Order{
		Name:        no.Name,
		Address:     no.Address,
		PhoneNo:     no.PhoneNo,
		DeliverDate: no.DeliverDate,
		Items:       orderItems,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	ord, err := c.order.Create(o)
	if err != nil {
		log.Println(err)
		return order.Order{}, err
	}

	return ord, nil

}
