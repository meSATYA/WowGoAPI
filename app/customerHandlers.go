package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meSATYA/WowGoAPI/service"
)

//type Customer struct {
//	Name    string `json:"name" xml:"name"`
//	City    string `json:"city" xml:"city"`
//	ZipCode string `json:"zip_code" xml:"zip_code"`
//}

type CustomerHandlers struct {
	service service.CustomerService
}

//func greet(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello World!!!")
//}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{Name: "John Doe", City: "New York", ZipCode: "10001"},
	//	{Name: "Jane Smith", City: "Los Angeles", ZipCode: "90003"},
	//}

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomerById(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

//func getCustomer(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	fmt.Fprintf(w, "Customer with name %s", vars["customer_id"])
//}
//
//func createCustomer(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Created Customer :)")
//}
