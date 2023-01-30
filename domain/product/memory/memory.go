package memory

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/tavern/domain/product"
	"sync"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{products: map[uuid.UUID]product.Product{}}
}

func (m *MemoryProductRepository) GetAll() ([]product.Product, error) {
	var products []product.Product
	for _, p := range m.products {
		products = append(products, p)
	}
	return products, nil
}

func (m *MemoryProductRepository) GetById(id uuid.UUID) (product.Product, error) {
	if p, ok := m.products[id]; ok {
		return p, nil
	}
	return product.Product{}, product.ErrProductNotFound

}

func (m *MemoryProductRepository) Add(p product.Product) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.products[p.GetId()]; ok {
		return product.ErrProductAlreadyExists
	}
	m.products[p.GetId()] = p
	return nil
}

func (m *MemoryProductRepository) Update(p product.Product) error {
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
