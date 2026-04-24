package dto

import (
	"strings"

	"github.com/meSATYA/WowGoAPI/errs"
)

type NewAccountRequest struct {
	CustomerID  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.AccountValidationError("Account amount should be greater than 5000 to open a new account")
	}

	if strings.ToLower(r.AccountType) != "savings" && strings.ToLower(r.AccountType) != "current" {
		return errs.AccountValidationError("Account type should be either savings or current")
	}
	return nil
}
