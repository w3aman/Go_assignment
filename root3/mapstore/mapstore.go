package mapstore

import (
	"errors"
	"fmt"
	"root3/domain"
)

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore{
		store: make(map[string]domain.Customer)}
}

func (m *MapStore) Create(c domain.Customer) error {
	if c.ID == "" {
		return errors.New("id not found")
	}

	m.store[c.ID] = c
	fmt.Println("user id", c.ID)
	return nil
}

func (m *MapStore) Update(id string, c domain.Customer) error {
	if id == "" {
		return errors.New("id not found")
	}
	m.store[id] = c
	fmt.Println("updated id", id)
	return nil
}

func (m *MapStore) GetById(id string) (domain.Customer, error) {
	if id == "" {
		return domain.Customer{}, errors.New("id not found")
	}

	c, ok := m.store[id]
	if !ok {
		return domain.Customer{}, errors.New("customer not found")
	}
	return c, nil
}

func (m *MapStore) GetAll() ([]domain.Customer, error) {
	if len(m.store) == 0 {
		return []domain.Customer{}, errors.New("store is empty")
	}
	clist := []domain.Customer{}
	for _, c := range m.store {
		clist = append(clist, c)

	}
	return clist, nil
}

func (m *MapStore) Delete(id string) error {
	if id == "" {
		return errors.New("id not found")
	}
	delete(m.store, id)
	return nil
}
