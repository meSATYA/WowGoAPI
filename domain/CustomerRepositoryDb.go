package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	db *sql.DB
}

func (d *CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.db.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var cst Customer
		err := rows.Scan(&cst.Id, &cst.Name, &cst.City, &cst.Zipcode, &cst.DateOfBirth, &cst.Status)
		if err != nil {
			log.Println("Error while scanning customers " + err.Error())
			return nil, err
		}
		customers = append(customers, cst)
	}
	return customers, nil
}

func NewCustomerRepositoryDb() *CustomerRepositoryDb {
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return &CustomerRepositoryDb{db}
}
