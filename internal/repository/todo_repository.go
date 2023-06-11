package repository

import (
	"final-project-1/internal/domain"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type TodoRepository interface {
	GetTodos() ([]domain.Todo, error)
	GetTodoByID(id string) (*domain.Todo, error)
	CreateTodo(todo *domain.Todo) error
	UpdateTodoByID(id string, todo *domain.Todo) error
	DeleteTodoByID(id string) error
}

type todoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		DB: db,
	}
}

func (r *todoRepository) GetTodos() ([]domain.Todo, error) {
	var todos []domain.Todo
	err := r.DB.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepository) GetTodoByID(id string) (*domain.Todo, error) {
	var todo domain.Todo
	err := r.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) CreateTodo(todo *domain.Todo) error {
	todo.ID = uuid.New().String()
	err := r.DB.Create(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *todoRepository) UpdateTodoByID(id string, todo *domain.Todo) error {
	err := r.DB.Model(&domain.Todo{}).Where("id = ?", id).Updates(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *todoRepository) DeleteTodoByID(id string) error {
	err := r.DB.Where("id = ?", id).Delete(&domain.Todo{}).Error
	if err != nil {
		return err
	}
	return nil
}
