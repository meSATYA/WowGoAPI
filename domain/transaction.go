package domain

import "github.com/meSATYA/WowGoAPI/dto"

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionID   string  `db:"transaction_id"`
	AccountID       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}

func (t Transaction) ToMakeTransactionResponseDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionID:   t.TransactionID,
		AccountID:       t.AccountID,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
