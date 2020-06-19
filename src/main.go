package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	authetication "src/m/v2/src/authentication"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[int]string, 0)
	data[101] = "KAVIN"
	data[102] = "RAM"
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error Occured")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func main() {
	router := mux.NewRouter()
	// Routes consist of a path and a handler function.

	// -> Register Auth Routes here ->
	auth := router.PathPrefix("/auth").Subrouter()
	authetication.RegisterAuthRoutes(auth)

	router.HandleFunc("/fetch_details", YourHandler).Methods("POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", router))
}
