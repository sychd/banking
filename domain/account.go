package domain

import (
	dto "github.com/sychd/banking/dto/account"
	"github.com/sychd/banking/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

func (a Account) ToAccountResponseDto() dto.AccountResponse {
	return dto.AccountResponse{
		AccountId:   a.AccountId,
		CustomerId:  a.CustomerId,
		OpeningDate: a.OpeningDate,
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      a.Status,
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}

type AccountRepository interface {
	Save(request Account) (*Account, *errs.AppError)
	SaveTransaction(t Transaction) (*Transaction, *errs.AppError)
	ById(id string) (*Account, *errs.AppError)
}
