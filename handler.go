package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// In-memory store for todos and ID generator
var todos = make(map[int]*Todo)
var idCounter = 1

// List all non-deleted todos
func listTodos(w http.ResponseWriter, r *http.Request) {
	var activeTodos []*Todo
	for _, todo := range todos {
		if !todo.Deleted {
			activeTodos = append(activeTodos, todo)
		}
	}
	json.NewEncoder(w).Encode(activeTodos)
}

// Create a new todo
func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.ID = idCounter
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	todos[todo.ID] = &todo
	idCounter++
	json.NewEncoder(w).Encode(todo)
}

// Get a single todo by ID
func getTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	todo, ok := todos[id]
	if !ok || todo.Deleted {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

// Update an existing todo
func updateTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	todo, ok := todos[id]
	if !ok || todo.Deleted {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	var updatedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.Title = updatedTodo.Title
	todo.Completed = updatedTodo.Completed
	todo.UpdatedAt = time.Now()
	json.NewEncoder(w).Encode(todo)
}

// Soft delete a todo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	todo, ok := todos[id]
	if !ok || todo.Deleted {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	todo.Deleted = true
	todo.UpdatedAt = time.Now()
	w.WriteHeader(http.StatusNoContent)
}
