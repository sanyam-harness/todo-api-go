# TODO API in Go

This is a simple REST API built using Go and Gorilla Mux.  
It manages TODO items in memory (no database), and follows a clean layered architecture with a service layer.

---

## 🔧 Features

- Create a new TODO
- List all active (non-deleted) TODOs
- Get TODO by ID
- Update TODO
- Soft delete TODO (marks as deleted but does not remove)

---

## 🧱 Project Structure

main.go → Entry point and route setup
handler.go → Handles HTTP requests
service.go → Business logic (CRUD operations)
todo.go → Data model (Todo struct)


---

## 🚀 How to Run

Make sure you have [Go installed](https://golang.org/dl/), then:

```bash
go run main.go handler.go service.go todo.go

The server will run on: http://localhost:8080

## 🧪 How to Test the API with curl

You can test the TODO API using `curl` commands from your terminal.

### 1. Create a TODO

```bash
curl -X POST http://localhost:8080/todos \
-H "Content-Type: application/json" \
-d '{"title": "Learn Go"}'

### 2. List all TODOs

'''bash
curl http://localhost:8080/todos

### 3. Get a TODO by ID

'''bash
curl http://localhost:8080/todos/1

### 4. Update a TODO 

'''bash
curl -X PUT http://localhost:8080/todos/1 \
-H "Content-Type: application/json" \
-d '{"title": "Learn Go in depth", "completed": true}'

### 5. Soft Delete a TODO

'''bash
curl -X DELETE http://localhost:8080/todos/1

### 6. Confirm Deletion (Should Show Empty or Not Found)

'''bash
curl http://localhost:8080/todos/1



