package dto

import (
	"github.com/sychd/banking/errs"
	"strings"
)

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	CustomerId      string  `json:"customer_id"`
	TransactionType string  `json:"transaction_type"` // "withdrawal" | "deposit"
	Amount          float64 `json:"amount"`
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.Amount < 0 {
		return errs.NewValidationError("Amount should be more than 0")
	}

	if strings.ToLower(r.TransactionType) != "withdrawal" && strings.ToLower(r.TransactionType) != "deposit" {
		return errs.NewValidationError("Wrong transaction type provided")
	}

	return nil
}

func (r TransactionRequest) IsWithdrawal() bool {
	return r.TransactionType == "withdrawal"
}

func (r TransactionRequest) UpdateAmount(amount float64) {
	if strings.ToLower(r.TransactionType) != "withdrawal" {
		r.Amount = r.Amount - amount
	} else {
		r.Amount = r.Amount + amount
	}
}
