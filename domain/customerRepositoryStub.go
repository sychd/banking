package domain

import "github.com/dsych/banking/errs"

// CustomerRepositoryStub adaptor
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll(status string) ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) ById(id string) (*Customer, *errs.AppError) {
	return &s.customers[0], nil
}

func NewCustomerRepositoryStub() CustomerRepository {
	customers := []Customer{
		{
			Id:          "1",
			Name:        "Marko",
			City:        "Dodo",
			Zipcode:     "123",
			DateOfBirth: "12-12-1992",
			Status:      "active",
		},
	}
	return CustomerRepositoryStub{customers}
}
