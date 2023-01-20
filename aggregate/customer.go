package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mansoorceksport/ddd-go/entity"
	"github.com/mansoorceksport/ddd-go/valueobject"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have valid name")
)

// Customer Aggregate is the combination of entities and value objects.
// business logic of customer needs to be inside the aggregate
type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

// NewCustomer is a factory to create a new customer aggregate.
// it will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	return Customer{
		person: &entity.Person{
			ID:   uuid.New(),
			Name: name,
		},
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}
