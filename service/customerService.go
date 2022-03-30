package service

import (
	"github.com/dsych/banking/domain"
	"github.com/dsych/banking/dto/customer"
	"github.com/dsych/banking/errs"
)

// CustomerService Port
type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository // it is a dependency, not concrete implementation
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	return s.repository.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repository.ById(id)
	if err != nil {
		return nil, err
	}

	return c.ToDTO(), nil
}

// NewCustomerService Helper function to instantiate this service.
// NewCustomerService takes an argument that is a dependency that we can inject into our customer service
func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repository: repository}
}
