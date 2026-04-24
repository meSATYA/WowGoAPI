package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/meSATYA/WowGoAPI/domain"
	"github.com/meSATYA/WowGoAPI/logger"
	"github.com/meSATYA/WowGoAPI/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		logger.Fatal("SERVER_ADDRESS or SERVER_PORT is not set")
	}
}

func Start() {
	sanityCheck()

	// Custom Multiplexer
	muxR := mux.NewRouter()

	//Wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Defining routes
	muxR.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	muxR.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)

	//muxR.HandleFunc("/greet", greet).Methods(http.MethodGet)
	//muxR.HandleFunc("/customer", createCustomer).Methods(http.MethodPost)
	//muxR.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	// starting the server

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), muxR))
}
