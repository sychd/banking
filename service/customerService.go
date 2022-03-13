package service

import "github.com/dsych/banking/domain"

// CustomerService Port
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository // it is a dependency, not concrete implementation
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repository.FindAll()
}

// NewCustomerService Helper function to instantiate this service.
// NewCustomerService takes an argument that is a dependency that we can inject into our customer service
func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repository: repository}
}
