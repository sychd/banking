package service

import (
	"github.com/sychd/banking/domain"
	"github.com/sychd/banking/dto/customer"
	"github.com/sychd/banking/errs"
)

// CustomerService Port
//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package service github.com/sychd/banking/service CustomerService
type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository // it is a dependency, not concrete implementation
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	customers, appErr := s.repository.FindAll(status)
	if appErr != nil {
		return nil, appErr
	}

	dtos := make([]dto.CustomerResponse, len(customers))
	for i, value := range customers {
		dtos[i] = *value.ToDTO()
	}

	return dtos, nil
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
