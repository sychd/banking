package domain

import (
	dto "github.com/dsych/banking/dto/account"
	"github.com/dsych/banking/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(request Account) (*Account, *errs.AppError)
}
