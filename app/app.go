package app

import (
	"github.com/dsych/banking/domain"
	"github.com/dsych/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	err := http.ListenAndServe("localhost:8000", router)

	if err != nil {
		log.Fatal(err)
	}
}
