package service

import (
	"errors"

	"github.com/VictorArruda22/api-golang/internal/entities"
)

type CustomerService struct {
	rp entities.CustomerRepository
}

func NewCustomerService(rp entities.CustomerRepository) *CustomerService {
	return &CustomerService{rp: rp}
}

func (s *CustomerService) GetAll() ([]entities.Customer, error) {
	return s.rp.GetAll()
}

func (s *CustomerService) GetByID(id int) ([]entities.Customer, error) {
	return s.rp.GetByID(id)
}

func (s *CustomerService) Create(customer entities.Customer) ([]entities.Customer, error) {
	return s.rp.Create(customer)
}

func (s *CustomerService) Update(customer entities.Customer) ([]entities.Customer, error) {
	//Verifica se o customer existe
	existingCustomer, err := s.rp.GetByID(customer.Id)
	if err != nil {
		return nil, err
	}
	if len(existingCustomer) == 0 {
		return nil, errors.New("Customer não encontrado.")
	}
	return s.rp.Update(customer)
}

func (s *CustomerService) Delete(id int) error {
	//Verifica se o customer existe
	existingCustomer, err := s.rp.GetByID(id)
	if err != nil {
		return err
	}
	if len(existingCustomer) == 0 {
		return errors.New("Nenhum customer encontrado.")
	}
	return s.rp.Delete(id)
}
