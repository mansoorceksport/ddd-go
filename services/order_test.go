package services

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/ddd-go/aggregate"
	"testing"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	peanuts, err := aggregate.NewProduct("Peanuts", "Snacks", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := aggregate.NewProduct("Wine", "nasty drink", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{beer, peanuts, wine}
}

func TestOrderService_CreateOrder(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(WithMemoryCustomerRepository(), WithMemoryProductRepository(products))
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("mansoor")
	err = os.customer.Add(cust)

	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetId(),
		products[1].GetId(),
	}

	err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Fatal(err)
	}

}
