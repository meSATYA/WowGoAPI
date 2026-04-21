package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meSATYA/WowGoAPI/domain"
	"github.com/meSATYA/WowGoAPI/service"
)

func Start() {
	// Custom Multiplexer
	muxR := mux.NewRouter()

	//Wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// Defining routes
	muxR.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	//muxR.HandleFunc("/greet", greet).Methods(http.MethodGet)
	//muxR.HandleFunc("/customer", createCustomer).Methods(http.MethodPost)
	//muxR.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	// starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", muxR))
}
