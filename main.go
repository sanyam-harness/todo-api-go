package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define API routes and link them to handlers
	r.HandleFunc("/todos", listTodos).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", getTodo).Methods("GET")
	r.HandleFunc("/todos/{id:[0-9]+}", updateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id:[0-9]+}", deleteTodo).Methods("DELETE")

	// Start the HTTP server
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
