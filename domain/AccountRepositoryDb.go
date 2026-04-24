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

func (d AccountRepositoryDb) SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError) {
	// starting the database transaction block
	tx, err := d.db.Begin()
	if err != nil {
		logger.Error("Error while starting transaction " + err.Error())
		return nil, errs.CustomUnexpectedError("Unexpeceted Database Error")
	}

	// inserting the bank account transaction
	result, _ := tx.Exec(`insert into transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)`,
		transaction.AccountID, transaction.Amount, transaction.TransactionType, transaction.TransactionDate)

	// updating the account balance
	if transaction.IsWithdrawal() {
		_, err = tx.Exec(`update accounts set amount = amount - ? where account_id = ?`, transaction.Amount, transaction.AccountID)
	} else {
		_, err = tx.Exec(`update accounts set amount = amount + ? where account_id = ?`, transaction.Amount, transaction.AccountID)
	}

	// in case of any error, roll back the transaction from both the tables
	if err != nil {
		logger.Error("Error while inserting transaction " + err.Error())
		tx.Rollback()
		return nil, errs.CustomUnexpectedError("Unexpeceted Database Error")
	}

	// if all good, commit the transaction
	err = tx.Commit()
	if err != nil {
		logger.Error("Error while committing transaction " + err.Error())
		return nil, errs.CustomUnexpectedError("Unexpeceted Database Error")
	}

	// getting the last transaction id from the transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted id " + err.Error())
		return nil, errs.CustomUnexpectedError("Unexpeceted Database Error")
	}

	// Getting the latest account information from the account table
	account, appErr := d.FindBy(transaction.AccountID)
	if appErr != nil {
		return nil, appErr
	}

	transaction.TransactionID = strconv.FormatInt(transactionId, 10)

	// udating the transation struct with latest account balance
	transaction.Amount = account.Amount
	return &transaction, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account
	err := d.db.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.CustomUnexpectedError("Unexpected database error")
	}
	return &account, nil
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
