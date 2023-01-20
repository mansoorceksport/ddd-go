// Package memory is a in-memory implementation of customer repository
package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mansoorceksport/ddd-go/aggregate"
	"github.com/mansoorceksport/ddd-go/domain/customer"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {

	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	// make sure the map function is initialised.
	// aggregate.customer factory function must have ensured the customer map is initialised.
	// this is for extra safety
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	// make sure customer is not already in repo
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}

func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exists: %w", customer.ErrUpdateCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
