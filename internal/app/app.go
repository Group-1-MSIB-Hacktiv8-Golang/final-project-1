package app

import "final-project-1/internal/domain"

type App struct {
	TodoService domain.TodoService
}

func NewApp(todoService domain.TodoService) *App {
	return &App{
		TodoService: todoService,
	}
}
