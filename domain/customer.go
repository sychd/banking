package domain

import "github.com/dsych/banking/errs"

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
