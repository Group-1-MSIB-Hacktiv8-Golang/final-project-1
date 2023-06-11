package domain

import "time"

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TodoService interface {
	GetTodos() ([]Todo, error)
	GetTodoByID(id string) (*Todo, error)
	CreateTodo(todo *Todo) error
	UpdateTodoByID(id string, todo *Todo) error
	DeleteTodoByID(id string) error
}
