package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/meSATYA/WowGoAPI/errs"
	"github.com/meSATYA/WowGoAPI/logger"
)

type AccountRepositoryDb struct {
	db *sqlx.DB
}

func (d AccountRepositoryDb) Save(account Account) (*Account, *errs.AppError) {
	sqlInsert := "insert into accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := d.db.Exec(sqlInsert, account.CustomerID, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logger.Error("Error while creating new account " + err.Error())
		return nil, errs.CustomUnexpectedError("Unexpeceted Database Error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted id " + err.Error())
		return nil, errs.CustomUnexpectedError("Unexpeceted Database Error")
	}
	account.AccountID = strconv.FormatInt(id, 10)
	return &account, nil
}

func NewAccountRepositoryDb(db *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{db: db}
}
