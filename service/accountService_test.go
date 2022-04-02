package service

import (
	"github.com/golang/mock/gomock"
	"github.com/sychd/banking/domain"
	dto "github.com/sychd/banking/dto/account"
	"github.com/sychd/banking/errs"
	mockedDomain "github.com/sychd/banking/mocks/domain"
	"testing"
	"time"
)

var mockRepo *mockedDomain.MockAccountRepository
var service AccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = mockedDomain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepo)

	return func() {
		service = nil
		ctrl.Finish()
	}
}

func Test_should_return_error_if_validation_failed(t *testing.T) {
	request := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      0,
	}

	service := NewAccountService(nil)

	_, appError := service.NewAccount(&request)

	if appError == nil {
		t.Error("Failed while testing account validation")
	}
}

func Test_should_return_error_if_account_cannot_be_created(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}

	acc := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	mockRepo.EXPECT().Save(acc).Return(nil, errs.NewUnexpectedError("Unexpected database error"))
	_, appError := service.NewAccount(&req)

	if appError == nil {
		t.Error("Expected to have error")
	}
}

func Test_should_return_account_when_it_was_successfully_created(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}

	acc := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	accWithId := acc
	accWithId.AccountId = "1000"

	mockRepo.EXPECT().Save(acc).Return(&accWithId, nil)
	newAccount, appError := service.NewAccount(&req)

	if appError != nil {
		t.Error("Expected not to have error")
	}

	if newAccount.AccountId != accWithId.AccountId {
		t.Error("Expected to have matching account id")
	}
}
