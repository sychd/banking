package app

import (
	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	dto "github.com/sychd/banking/dto/customer"
	"github.com/sychd/banking/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
var ch CustomerHandlers

// for mock usage we need to
// 1. install> go install github.com/golang/mock/mockgen@v1.6.0
// 2. write "metadata comment" for generation above CustomerService interface
// 3. run go generate ./... in root

var mockedService *service.MockCustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedService = service.NewMockCustomerService(ctrl)

	ch = CustomerHandlers{mockedService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)

	return func() {
		ctrl.Finish()
		router = nil
	}
}

func Test_should_return_customers_with_code_200(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	dummyCustomers := []dto.CustomerResponse{
		{
			Id:          "1",
			Name:        "Marko",
			City:        "Dodo",
			Zipcode:     "123",
			DateOfBirth: "12-12-1992",
			Status:      "active",
		},
	}
	mockedService.EXPECT().GetAllCustomers("").Return(dummyCustomers, nil)
	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Wrong status code")
	}
}

func Test_should_return_code_500_with_error_message(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	mockedService.EXPECT().GetAllCustomers("").Return(nil, errs.NewUnexpectedError("Error"))
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Wrong status code")
	}
}
