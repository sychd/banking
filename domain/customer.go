package domain

type Customer struct {
	Id      string
	Name    string
	City    string
	Zipcode string
	Birth   string
	Status  string
}

// CustomerRepository Interface
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
