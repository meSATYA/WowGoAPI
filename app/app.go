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
	muxR.HandleFunc("/greet", greet)
	muxR.HandleFunc("/customers", getAllCustomers)
	// starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", muxR))
}
