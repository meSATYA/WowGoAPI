package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/meSATYA/WowGoAPI/domain"
	"github.com/meSATYA/WowGoAPI/logger"
	"github.com/meSATYA/WowGoAPI/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		logger.Fatal("SERVER_ADDRESS or SERVER_PORT is not set")
	}
}

func getDbClient() *sqlx.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	dbClient, err := sqlx.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}

	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxIdleConns(10)
	dbClient.SetMaxOpenConns(10)

	return dbClient
}

func Start() {
	sanityCheck()

	// Custom Multiplexer
	muxR := mux.NewRouter()

	//Wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub()) }

	dbClient := getDbClient()
	customerRespositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRespositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// Defining routes
	muxR.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	muxR.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)
	muxR.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	//muxR.HandleFunc("/greet", greet).Methods(http.MethodGet)
	//muxR.HandleFunc("/customer", createCustomer).Methods(http.MethodPost)
	//muxR.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	// starting the server

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), muxR))
}
