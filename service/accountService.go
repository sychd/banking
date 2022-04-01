package service

import (
	"github.com/sychd/banking/domain"
	dto "github.com/sychd/banking/dto/account"
	"github.com/sychd/banking/errs"
	"time"
)

// AccountService Port
type AccountService interface {
	NewAccount(*dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(req *dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
	GetAccountById(id string) (*dto.AccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repository domain.AccountRepository // it is a dependency, not concrete implementation
}

func (s DefaultAccountService) GetAccountById(id string) (*dto.AccountResponse, *errs.AppError) {
	newAccount, err := s.repository.ById(id)
	if err != nil {
		return nil, err
	}
	res := newAccount.ToAccountResponseDto()

	return &res, nil
}

func (s DefaultAccountService) MakeTransaction(req *dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.IsWithdrawal() {
		account, err := s.repository.ById(req.AccountId)
		if err != nil {
			return nil, err
		}

		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Account has not enough credits for this operation")
		}
	}

	t := domain.Transaction{
		TransactionId:   "",
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	transaction, appErr := s.repository.SaveTransaction(t)
	if appErr != nil {
		return nil, appErr
	}
	transDto := transaction.ToTransactionResponseDto()

	return &transDto, nil
}

func (s DefaultAccountService) NewAccount(req *dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repository.Save(a)

	if err != nil {
		return nil, err
	}

	res := newAccount.ToNewAccountResponseDto()

	return &res, nil
}

// NewAccountService Helper function to instantiate this service.
// NewAccountService takes an argument that is a dependency that we can inject into our customer service
func NewAccountService(repository domain.AccountRepository) AccountService {
	return DefaultAccountService{repository: repository}
}
