package memory

import (
	"errors"
	"github.com/mansoorceksport/tavern/domain/common/idempotent"
	"sync"
)

var (
	ErrKeyCannotBeEmpty = errors.New("Idempotent-Key cannot be empty")
	ErrDuplicateRequest = errors.New("create order is a idempotent, cannot handle duplicate request")
)

type Memory struct {
	sync.Mutex
	keys map[string]bool
}

func NewIdempotent() idempotent.Idempotent {
	return &Memory{
		keys: map[string]bool{},
	}
}

func (i Memory) Check(k string) error {
	if k == "" {
		return ErrKeyCannotBeEmpty
	}

	if _, ok := i.keys[k]; ok {
		return ErrDuplicateRequest
	}

	return nil
}

func (i Memory) Add(k string) error {
	if k == "" {
		return ErrKeyCannotBeEmpty
	}

	if _, ok := i.keys[k]; ok {
		return ErrDuplicateRequest
	}

	i.Lock()
	defer i.Unlock()
	i.keys[k] = true
	return nil
}
