package tavern

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/tavern/domain/product"
	"github.com/mansoorceksport/tavern/services/order"
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

func TestTavernService_Order(t *testing.T) {
	products := init_products(t)
	//ctx := context.Background()
	//orderService, err := order.NewOrderService(order.WithMongoCustomerRepository(ctx, "mongodb://localhost:9000"), order.WithMemoryProductRepository(products))
	orderService, err := order.NewOrderService(order.WithMemoryIdempotent(), order.WithMemoryCustomerRepository(), order.WithMemoryProductRepository(products))
	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(orderService))

	uid, err := tavern.OrderService.AddCustomer("mawan")
	if err != nil {
		t.Fatal(err)
	}

	customerId := uid
	orders := []uuid.UUID{products[0].GetId(), products[1].GetId()}
	idempotentKey := uuid.NewString()
	err = tavern.Order(idempotentKey, customerId, orders)
	if err != nil {
		t.Fatal(err)
	}
}
