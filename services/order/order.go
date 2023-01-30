package order

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/mansoorceksport/tavern/domain/common/idempotent"
	memoryIdempotent "github.com/mansoorceksport/tavern/domain/common/idempotent/repository/memory"
	redisIdempotent "github.com/mansoorceksport/tavern/domain/common/idempotent/repository/redis"
	"github.com/mansoorceksport/tavern/domain/customer"
	"github.com/mansoorceksport/tavern/domain/customer/memory"
	"github.com/mansoorceksport/tavern/domain/customer/mongo"
	"github.com/mansoorceksport/tavern/domain/product"
	prodmem "github.com/mansoorceksport/tavern/domain/product/memory"
	"log"
)

var (
	ERRIdempotentIsMandatory = errors.New("idempotent is mandatory")
	ERRCustomerIsMandatory   = errors.New("customer is mandatory")
	ERRProductIsMandatory    = errors.New("product is mandatory")
)

// OrderConfiguration service configurator generator pattern
type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	idempotent idempotent.Idempotent
	customer   customer.Repository
	products   product.Repository
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

	// check all the required components is enabled
	if os.idempotent != nil {
		return nil, ERRIdempotentIsMandatory
	}

	if os.customer != nil {
		return nil, ERRCustomerIsMandatory
	}

	if os.products != nil {
		return nil, ERRProductIsMandatory
	}

	return os, nil
}

// WithCustomerRepository applies a customer repository to a OrderService
func WithCustomerRepository(cr customer.Repository) OrderConfiguration {
	// return a function that matches the order configuration alias
	return func(os *OrderService) error {
		os.customer = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		m, err := mongo.New(ctx, connectionString)
		if err != nil {
			return err
		}
		os.customer = m
		return nil
	}
}

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
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

func WithMemoryIdempotent() OrderConfiguration {
	return func(os *OrderService) error {
		i := memoryIdempotent.NewIdempotent()
		os.idempotent = i
		return nil
	}
}

func WithRedisIdempotent(conn string) OrderConfiguration {
	return func(os *OrderService) error {
		os.idempotent = redisIdempotent.NewRedis("")
		return nil
	}
}

func (o *OrderService) CreateOrder(idempotentKey string, customerId uuid.UUID, productsIDs []uuid.UUID) (float64, error) {

	// check if request is idempotent
	if err := o.idempotent.Check(idempotentKey); err != nil {
		return 0, err
	}

	// add the idempotent key to memory
	if err := o.idempotent.Add(idempotentKey); err != nil {
		return 0, err
	}

	// fetch the customer
	c, err := o.customer.Get(customerId)
	if err != nil {
		return 0, err
	}

	// Get each Product,
	var products []product.Product
	var total float64
	for _, id := range productsIDs {
		p, err := o.products.GetById(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("customer: %s has ordered %d products with IdempotentKey %s", c.GetID(), len(products), idempotentKey)

	return total, nil
}

func (o *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}
	err = o.customer.Add(c)
	if err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}
