package memory

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/ddd-go/aggregate"
	"github.com/mansoorceksport/ddd-go/domain/product"
	"sync"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{products: map[uuid.UUID]aggregate.Product{}}
}

func (m *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product
	for _, p := range m.products {
		products = append(products, p)
	}
	return products, nil
}

func (m *MemoryProductRepository) GetById(id uuid.UUID) (aggregate.Product, error) {
	if p, ok := m.products[id]; ok {
		return p, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound

}

func (m *MemoryProductRepository) Add(p aggregate.Product) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.products[p.GetId()]; ok {
		return product.ErrProductAlreadyExists
	}
	m.products[p.GetId()] = p
	return nil
}

func (m *MemoryProductRepository) Update(p aggregate.Product) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.products[p.GetId()]; !ok {
		return product.ErrProductNotFound
	}
	m.products[p.GetId()] = p
	return nil
}

func (m *MemoryProductRepository) Delete(id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(m.products, id)
	return nil
}
