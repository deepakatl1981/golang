// main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/emps", getEmployees).Methods("GET")

	http.ListenAndServe(":8888", r)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{Sample response}")
}
