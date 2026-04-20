package app

import (
	"log"
	"net/http"
)

func Start() {
	// Go built-in multiplexer
	// Defining routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
