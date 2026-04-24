package service

import (
	"time"

	"github.com/meSATYA/WowGoAPI/domain"
	"github.com/meSATYA/WowGoAPI/dto"
	"github.com/meSATYA/WowGoAPI/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountID:   "",
		CustomerID:  request.CustomerID,
		OpeningDate: time.Now().Format(time.RFC3339),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

func (s DefaultAccountService) MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	//incoming request validation
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	// service side validation for checking the available balance in the account
	if request.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(request.AccountId)
		if err != nil {
			return nil, err
		}

		if !account.CanWithdraw(request.Amount) {
			return nil, errs.AccountValidationError("Insufficient balance for withdrawal")
		}
	}

	// if everything is fine, build the domain object and save the transaction
	t := domain.Transaction{
		AccountID:       request.AccountId,
		Amount:          request.Amount,
		TransactionType: request.TransactionType,
		TransactionDate: time.Now().Format(time.RFC3339),
	}
	transaction, appErr := s.repo.SaveTransaction(t)
	if appErr != nil {
		return nil, appErr
	}
	response := transaction.ToMakeTransactionResponseDto()
	return &response, nil
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repository}
}
