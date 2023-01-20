package customer

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mansoorceksport/ddd-go/aggregate"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer      = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid uuid.UUID) (aggregate.Customer, error)
	Add(customer aggregate.Customer) error
	Update(customer aggregate.Customer) error
}
