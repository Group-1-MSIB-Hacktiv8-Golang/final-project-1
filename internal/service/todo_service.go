package service

import (
	"final-project-1/internal/domain"
	"final-project-1/internal/repository"

	"github.com/go-playground/validator/v10"
)

type todoService struct {
	TodoRepository repository.TodoRepository
	Validator      *validator.Validate
}

func NewTodoService(todoRepository repository.TodoRepository) domain.TodoService {
	return &todoService{
		TodoRepository: todoRepository,
		Validator:      validator.New(),
	}
}

func (s *todoService) GetTodos() ([]domain.Todo, error) {
	todos, err := s.TodoRepository.GetTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *todoService) GetTodoByID(id string) (*domain.Todo, error) {
	todo, err := s.TodoRepository.GetTodoByID(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *todoService) CreateTodo(todo *domain.Todo) error {
	err := s.Validator.Struct(todo)
	if err != nil {
		return err
	}
	err = s.TodoRepository.CreateTodo(todo)
	if err != nil {
		return err
	}
	return nil
}

func (s *todoService) UpdateTodoByID(id string, todo *domain.Todo) error {
	err := s.Validator.Struct(todo)
	if err != nil {
		return err
	}
	err = s.TodoRepository.UpdateTodoByID(id, todo)
	if err != nil {
		return err
	}
	return nil
}

func (s *todoService) DeleteTodoByID(id string) error {
	err := s.TodoRepository.DeleteTodoByID(id)
	if err != nil {
		return err
	}
	return nil
}
