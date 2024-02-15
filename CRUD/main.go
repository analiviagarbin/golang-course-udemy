package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// go get github.com/gorilla/mux

func main() {
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/users", server.NewUser).Methods(http.MethodPost)
	// Read
	router.HandleFunc("/users", server.SearchUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.SearchUser).Methods(http.MethodGet)
	// Update
	router.HandleFunc("/users/update/{id}", server.UpdateUser).Methods(http.MethodPut)
	// Delete
	router.HandleFunc("/users/delete/{id}", server.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Listen port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
