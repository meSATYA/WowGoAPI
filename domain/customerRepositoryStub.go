package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Satyajit", City: "Berlin", Zipcode: "13569", DateOfBirth: "12-12-1986", Status: "1"},
		{Id: "1002", Name: "Dambro", City: "Magdeburg", Zipcode: "24040", DateOfBirth: "22-04-1977", Status: "0"},
	}
	return CustomerRepositoryStub{customers: customers}
}
