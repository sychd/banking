package domain

import (
	dto "github.com/sychd/banking/dto/customer"
	"github.com/sychd/banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

var CustomerStatusDict = map[string]string{
	"inactive": "0",
	"active":   "1",
}

// CustomerRepository Interface (port)
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

func (c Customer) statusAsText() string {
	for k, v := range CustomerStatusDict {
		if v == c.Status {
			return k
		}
	}
	return "unknown"
}

func (c Customer) ToDTO() *dto.CustomerResponse {
	return &dto.CustomerResponse{
		Id: c.Id, Name: c.Name, City: c.City, Zipcode: c.Zipcode, DateOfBirth: c.DateOfBirth, Status: c.statusAsText(),
	}
}
