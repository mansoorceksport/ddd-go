package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/mansoorceksport/ddd-go/aggregate"
	"testing"
)

func TestTavernService_Order(t *testing.T) {
	products := init_products(t)
	ctx := context.Background()
	orderService, err := NewOrderService(WithMongoCustomerRepository(ctx, "mongodb://localhost:9000"), WithMemoryProductRepository(products))
	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(orderService))
	customer, err := aggregate.NewCustomer("mawan")
	if err != nil {
		t.Fatal(err)
	}

	err = tavern.OrderService.customer.Add(customer)
	if err != nil {
		t.Fatal(err)
	}

	customerId := customer.GetID()
	orders := []uuid.UUID{products[0].GetId(), products[1].GetId()}
	err = tavern.Order(customerId, orders)
	if err != nil {
		t.Fatal(err)
	}
}
