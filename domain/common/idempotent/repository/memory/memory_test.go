package memory

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/tavern/domain/common/idempotent"
	"testing"
)

type testCase struct {
	idempotent.Idempotent
	test          string
	expectedError error
	key           string
}

func TestMemory_Add(t *testing.T) {
	k := uuid.NewString()
	i := Memory{
		keys: map[string]bool{
			k: true,
		},
	}

	testCases := []testCase{{
		Idempotent:    i,
		test:          "empty key",
		expectedError: ErrKeyCannotBeEmpty,
		key:           "",
	}, {
		Idempotent:    i,
		test:          "duplicate key",
		expectedError: ErrDuplicateRequest,
		key:           k,
	}, {
		Idempotent:    i,
		test:          "success",
		expectedError: nil,
		key:           uuid.NewString(),
	}}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := i.Add(tc.key)
			if err != tc.expectedError {
				t.Fatalf("Expected Error %v got %v", tc.expectedError, err)
			}
		})
	}
}

func TestMemory_Check(t *testing.T) {
	k := uuid.NewString()
	i := NewIdempotent()
	i.Add(k)

	testCases := []testCase{{
		Idempotent:    i,
		test:          "empty key",
		expectedError: ErrKeyCannotBeEmpty,
		key:           "",
	}, {
		Idempotent:    i,
		test:          "duplicate key",
		expectedError: ErrDuplicateRequest,
		key:           k,
	}, {
		Idempotent:    i,
		test:          "Success",
		expectedError: nil,
		key:           uuid.NewString(),
	}}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := tc.Idempotent.Check(tc.key)
			if err != tc.expectedError {
				t.Fatalf("Expected expectedError %v got %v", tc.expectedError, err)
			}
		})
	}
}
