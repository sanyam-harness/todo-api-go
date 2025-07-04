package main

import (
	"errors"
	"time"
)

type TodoService struct {
	todos     map[int]*Todo
	idCounter int
}

func NewTodoService() *TodoService {
	return &TodoService{
		todos:     make(map[int]*Todo),
		idCounter: 1,
	}
}

func (s *TodoService) ListTodos() []*Todo {
	var result []*Todo
	for _, todo := range s.todos {
		if !todo.Deleted {
			result = append(result, todo)
		}
	}
	return result
}

func (s *TodoService) CreateTodo(todo *Todo) *Todo {
	todo.ID = s.idCounter
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	s.todos[todo.ID] = todo
	s.idCounter++
	return todo
}

func (s *TodoService) GetTodo(id int) (*Todo, error) {
	todo, ok := s.todos[id]
	if !ok || todo.Deleted {
		return nil, errors.New("todo not found")
	}
	return todo, nil
}

func (s *TodoService) UpdateTodo(id int, updated *Todo) (*Todo, error) {
	todo, ok := s.todos[id]
	if !ok || todo.Deleted {
		return nil, errors.New("todo not found")
	}
	todo.Title = updated.Title
	todo.Completed = updated.Completed
	todo.UpdatedAt = time.Now()
	return todo, nil
}

func (s *TodoService) DeleteTodo(id int) error {
	todo, ok := s.todos[id]
	if !ok || todo.Deleted {
		return errors.New("todo not found")
	}
	todo.Deleted = true
	todo.UpdatedAt = time.Now()
	return nil
}
