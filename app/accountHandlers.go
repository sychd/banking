package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	dto "github.com/sychd/banking/dto/account"
	"github.com/sychd/banking/service"
	"net/http"
)

type AccountHandlers struct {
	service service.AccountService
}

func (h *AccountHandlers) newAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = vars["customer_id"]
		account, err := h.service.NewAccount(&request)

		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h *AccountHandlers) makeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = vars["customer_id"]
		request.AccountId = vars["account_id"]
		account, err := h.service.MakeTransaction(&request)

		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
