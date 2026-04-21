package domain

type Customer struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	City   string `json:"city"`
	Zip    string `json:"zip"`
	DOB    string `json:"dob"`
	Status string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
