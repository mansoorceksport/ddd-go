package memory

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mansoorceksport/tavern/domain/customer"
	"testing"
)

func TestMemoryRepository_GetCustomer(t *testing.T) {
	type testCase struct {
		test        string
		id          uuid.UUID
		expectedErr error
	}

	c, err := customer.NewCustomer("mansoor")
	if err != nil {
		t.Fatal(err)
	}

	id := c.GetID()

	memory := MemoryRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: c,
		},
	}

	// table drive test
	testCases := []testCase{
		{
			test:        "no customer by id",
			id:          uuid.New(),
			expectedErr: customer.ErrCustomerNotFound,
		},

		{
			test:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := memory.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
