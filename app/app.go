package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/api/time", getTime).Queries("tz", "{time_zone}").Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	//router.HandleFunc("/customers/{customer_id}", getCustomer).Methods(http.MethodGet) // any match
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet) // numeric only [pattern for route]

	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
