package domain

import (
	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/sychd/banking/dto/account"
)

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

const WITHDRAWAL = "withdrawal"

func (t Transaction) ToTransactionResponseDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		NewAmount:       t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

func (t Transaction) IsWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

type TransactionRepository interface {
	Save(request Transaction) (*Transaction, *errs.AppError)
}
