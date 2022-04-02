package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_type_is_not_deposit_or_withdrawal(t *testing.T) {
	request := TransactionRequest{
		TransactionType: "invalid type",
	}

	appError := request.Validate()

	if appError.Message != "Wrong transaction type provided" {
		t.Error("Invalid message while testing tr. type")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing tr. type")
	}

	if appError == nil {
		t.Error("No error happened")
	}
}
func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	request := TransactionRequest{
		TransactionType: "deposit",
		Amount:          -1000,
	}

	appError := request.Validate()

	if appError.Message != "Amount should be more than 0" {
		t.Error("Invalid message while testing tr. type")
	}

	if appError == nil {
		t.Error("No error happened")
	}

}
