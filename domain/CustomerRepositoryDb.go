package domain

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/meSATYA/WowGoAPI/errs"
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

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.db.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.CustomerNotFound("Customer Not Found")
		}
		log.Println("Error while scanning customer " + err.Error())
		return nil, errs.CustomUnexpectedError("Unexpeceted Database Error")
	}
	return &c, nil
}

func NewCustomerRepositoryDb() *CustomerRepositoryDb {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return &CustomerRepositoryDb{db}
}
