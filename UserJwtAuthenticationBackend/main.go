package main

import (
	"fmt"
	"log"
	"net/http"

	"./controller"
	"./middleware"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Started")
	router := mux.NewRouter()
	router.HandleFunc("/signup", controller.RegistrationHandler).Methods("POST")
	router.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	router.HandleFunc("/fetch", middleware.Middleware).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}
