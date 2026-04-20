package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// Custom Multiplexer
	muxR := mux.NewRouter()
	// Defining routes
	muxR.HandleFunc("/greet", greet).Methods(http.MethodGet)
	muxR.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	muxR.HandleFunc("/customer", createCustomer).Methods(http.MethodPost)
	muxR.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	// starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", muxR))
}
