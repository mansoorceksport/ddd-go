package tavern

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/tavern/services/order"
	"log"
)

type TavernConfiguration func(tavern *TavernService) error

type TavernService struct {
	OrderService *order.OrderService
}

func NewTavern(configuration ...TavernConfiguration) (*TavernService, error) {
	tavern := &TavernService{}

	for _, cfg := range configuration {
		err := cfg(tavern)
		if err != nil {
			return nil, err
		}
	}
	return tavern, nil
}

func WithOrderService(orderService *order.OrderService) TavernConfiguration {
	return func(tavern *TavernService) error {
		tavern.OrderService = orderService
		return nil
	}
}

func (ts *TavernService) Order(customerID uuid.UUID, products []uuid.UUID) error {
	price, err := ts.OrderService.CreateOrder(customerID, products)
	if err != nil {
		return err
	}

	log.Printf("\nBill the customer: %0.0f\n", price)

	// Implement Billing service
	return nil
}
