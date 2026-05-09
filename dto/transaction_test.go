package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_valid(t *testing.T) {
	// Arrange
	request := TransactionRequest{
		TransactionType: "invalid_type",
	}
	//Act
	appError := request.Validate()

	//Assert
	if appError.Message != "Transaction type can only be deposit or withdrawal" {
		t.Error("Invalid message while testing transaction type")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction type")
	}
}

func Test_should_return_error_while_amount_is_negative(t *testing.T) {
	request := TransactionRequest{TransactionType: DEPOSIT, Amount: -100}

	appError := request.Validate()

	if appError.Message != "Amount cannot be less than zero" {
		t.Error("Invalid message while testing amount validation")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing amount validation")
	}
}
