package main

import (
	"time"
)

// Todo represents a task with a title, completion status, timestamps, and a soft delete flag.
type Todo struct {
	ID        int       `json:"id"`         // Unique identifier
	Title     string    `json:"title"`      // Description of the task
	Completed bool      `json:"completed"`  // Is the task done?
	Deleted   bool      `json:"-"`          // Soft delete (hidden from API output)
	CreatedAt time.Time `json:"created_at"` // Timestamp when created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when last updated
}
