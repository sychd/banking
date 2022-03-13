package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Customer struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Marko", City: "Dodo"},
		{Name: "Donald", City: "Dodo"},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func getTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loc, err := time.LoadLocation(vars["time_zone"])
	dt := time.Now()

	if err != nil {
		fmt.Fprint(w, dt.Format(time.RFC1123))
		return
	}

	fmt.Fprint(w, dt.In(loc).Format(time.RFC1123))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received...")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
