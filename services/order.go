package services

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/ddd-go/aggregate"
	"github.com/mansoorceksport/ddd-go/domain/customer"
	"github.com/mansoorceksport/ddd-go/domain/customer/memory"
	"github.com/mansoorceksport/ddd-go/domain/product"
	prodmem "github.com/mansoorceksport/ddd-go/domain/product/memory"
	"log"
)

// OrderConfiguration service configurator generator pattern
type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customer customer.CustomerRepository
	products product.ProductRepository
}

func NewOrderService(configurations ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	// loop through all the config and apply them
	for _, cfg := range configurations {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository applies a customer repository to a OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// return a function that matches the order configuration alias
	return func(os *OrderService) error {
		os.customer = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	//return func(os *OrderService) error {
	//	os.customer = memory.New()
	//	return nil
	//}
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.NewMemoryProductRepository()
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerId uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	// fetch the customer
	c, err := o.customer.Get(customerId)
	if err != nil {
		return 0, err
	}

	// Get each Product,
	var products []aggregate.Product
	var total float64
	for _, id := range productsIDs {
		p, err := o.products.GetById(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("customer: %s has ordered %d products", c.GetID(), len(products))

	return total, nil
}
