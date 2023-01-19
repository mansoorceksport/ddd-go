package aggregate

import (
	"github.com/mansoorceksport/ddd-go/entity"
	"github.com/mansoorceksport/ddd-go/valueobject"
)

// Customer Aggregate is the combination of entities and value objects.
// business logic of customer needs to inside the aggregate
type Customer struct {
	person      *entity.Person
	products    []*entity.Item
	transaction []valueobject.Transaction
}
