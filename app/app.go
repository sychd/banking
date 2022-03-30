package app

import (
	"fmt"
	"github.com/dsych/banking/domain"
	"github.com/dsych/banking/logger"
	"github.com/dsych/banking/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func Start() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("godotenv error: %s", err)
		return
	}

	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb(os.Getenv("CLEARDB_DATABASE_URL")))}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Queries("status", "{status:active|inactive}").Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	address := fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT"))
	logger.Infof("Server started at %s", address)
	err = http.ListenAndServe(address, router)

	if err != nil {
		log.Fatal(err)
	}
}
