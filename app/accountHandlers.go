package app

import (
	"encoding/json"
	dto "github.com/dsych/banking/dto/account"
	"github.com/dsych/banking/service"
	"github.com/gorilla/mux"
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
