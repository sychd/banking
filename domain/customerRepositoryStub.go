package domain

// CustomerRepositoryStub adaptor
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepository {
	customers := []Customer{
		{
			Id:      "1",
			Name:    "Marko",
			City:    "Dodo",
			Zipcode: "123",
			Birth:   "12-12-1992",
			Status:  "active",
		},
	}
	return CustomerRepositoryStub{customers}
}
