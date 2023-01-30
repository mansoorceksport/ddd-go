package order

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/tavern/domain/customer"
	"github.com/mansoorceksport/tavern/domain/product"
	"testing"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	peanuts, err := product.NewProduct("Peanuts", "Snacks", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := product.NewProduct("Wine", "nasty drink", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{beer, peanuts, wine}
}

func TestOrderService_CreateOrder(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(WithMemoryIdempotent(), WithMemoryCustomerRepository(), WithMemoryProductRepository(products))
	if err != nil {
		t.Fatal(err)
	}

	cust, err := customer.NewCustomer("mansoor")
	err = os.customer.Add(cust)

	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetId(),
		products[1].GetId(),
	}

	idempotentKey := uuid.NewString()
	_, err = os.CreateOrder(idempotentKey, cust.GetID(), order)
	if err != nil {
		t.Fatal(err)
	}

}
