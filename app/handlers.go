package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name string `json:"name"`
	City string `json:"city"`
}


func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer {
		{Name: "Marko", City: "Dodo" },
		{Name: "Donald", City: "Dodo" },
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
