package mapstore

import (
	"errors"
	"github.com/shijuvar/gokit/examples/training-api/domain"
)

type MapStore struct {
	store map[string]domain.Customer // An in-memory store  with a map
}

func NewMapStore() *MapStore {
	return &MapStore { store: make(map[string]domain.Customer)}
}

func (m MapStore) Create(customer domain.Customer) error {
	_, ok := m.store[customer.ID]
	if !ok {
		m.store[customer.ID] = customer
	}else{
		return errors.New("customer already exists")
	}
	return nil
}

func (m MapStore) Update(s string, customer domain.Customer) error {
	_, ok := m.store[s]
	if ok {
		m.store[customer.ID] = customer
	}else {
		return errors.New("customer doesn't exists")
	}
	return nil
}

func (m MapStore) Delete(s string) error {
	_, ok := m.store[s]
	if ok{
		delete(m.store,s)
	}else{
		return errors.New("no customer with given customer ID exists")
	}
	return nil
}

func (m MapStore) GetById(s string) (domain.Customer, error) {
	customer, ok:=m.store[s]
	if ok{
		return customer, nil
	}else{
		return customer,errors.New("customer doesn't exists")
	}
}

func (m MapStore) GetAll() ([]domain.Customer, error) {
	customerList := make([]domain.Customer, 0, len(m.store))
	for  _, cus := range m.store {
		customerList = append(customerList, cus)
	}
	return customerList,nil
}
