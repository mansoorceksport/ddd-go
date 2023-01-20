package aggregate_test

import (
	"errors"
	"github.com/mansoorceksport/ddd-go/aggregate"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	// Table driven Test cases
	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "valid name",
			name:        "Mansoor",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v got %v", tc.expectedErr, err)
			}
		})
	}
}
