package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitilizeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id]", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUsershandler).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUsers).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUsers).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitialMigration()
	InitilizeRouter()
}
