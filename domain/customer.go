package domain

import "github.com/meSATYA/WowGoAPI/errs"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
