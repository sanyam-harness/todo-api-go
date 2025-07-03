# TODO API in Go

This is a simple REST API built using Go. It allows you to manage TODO items in memory without using a database.

## Features

- List all TODO items
- Create a new TODO
- Get a TODO by ID
- Update a TODO
- Soft delete a TODO (marks it as deleted without removing it from memory)

## Project Files

- `main.go` – Starts the server and sets up routes
- `handler.go` – Contains handler functions for each API endpoint
- `todo.go` – Defines the TODO item struct

## Dependencies

- gorilla/mux (for routing)
