package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/meSATYA/WowGoAPI/errs"
	"github.com/meSATYA/WowGoAPI/logger"
)

type CustomerRepositoryDb struct {
	db *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.db.Select(&customers, findAllSql)
		//rows, err = d.db.Query(findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.db.Select(&customers, findAllSql, status)
		//rows, err = d.db.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.CustomUnexpectedError("Unexpeceted Database Error")
	}

	//err = sqlx.StructScan(rows, &customers)
	//
	//if err != nil {
	//	logger.Error("Error while scanning customers " + err.Error())
	//	return nil, errs.CustomUnexpectedError("Unexpeceted Database Error")
	//}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	//row := d.db.QueryRow(customerSql, id)
	var c Customer
	err := d.db.Get(&c, customerSql, id)

	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.CustomerNotFound("Customer Not Found")
		}
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errs.CustomUnexpectedError("Unexpeceted Database Error")
	}
	return &c, nil
}

func NewCustomerRepositoryDb(db *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{db}
}
