package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Satyajit", Email: "satyajit@gmail.com", Phone: "1234567890", City: "Berlin", Zip: "13569", DOB: "12-12-1986", Status: "active"},
		{Id: "1002", Name: "Dambro", Email: "dambroooo@gmail.com", Phone: "9876543210", City: "Magdeburg", Zip: "24040", DOB: "22-04-1977", Status: "inactive"},
	}
	return CustomerRepositoryStub{customers: customers}
}
