package service

import (
	"github.com/dsych/banking/domain"
	dto "github.com/dsych/banking/dto/account"
	"github.com/dsych/banking/errs"
	"time"
)

// AccountService Port
type AccountService interface {
	NewAccount(*dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repository domain.AccountRepository // it is a dependency, not concrete implementation
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
