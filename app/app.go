package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sychd/banking/domain"
	"github.com/sychd/banking/logger"
	"github.com/sychd/banking/service"
	"log"
	"net/http"
	"os"
	"time"
)

func createDbClient(dbUrl string) *sqlx.DB {
	client, err := sqlx.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
func Start() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("godotenv error: %s", err)
		return
	}

	router := mux.NewRouter()
	dbClient := createDbClient(os.Getenv("CLEARDB_DATABASE_URL"))

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb(dbClient))}
	ah := AccountHandlers{service.NewAccountService(domain.NewAccountRepositoryDb(dbClient))}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet).Name("GetAllCustomers")
	router.HandleFunc("/customers", ch.getAllCustomers).Queries("status", "{status:active|inactive}").Methods(http.MethodGet).Name("GetAllCustomers")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet).Name("GetCustomer")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.newAccount).Methods(http.MethodPost).Name("NewAccount")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.makeTransaction).Methods(http.MethodPost).Name("NewTransaction")

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	// starting server
	address := fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT"))
	logger.Infof("Server started at %s", address)
	err = http.ListenAndServe(address, router)

	if err != nil {
		log.Fatal(err)
	}
}
