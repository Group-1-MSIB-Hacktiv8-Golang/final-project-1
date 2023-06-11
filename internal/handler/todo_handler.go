package handler

import (
	"final-project-1/internal/app"
	"final-project-1/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	App *app.App
}

func NewTodoHandler(app *app.App) *TodoHandler {
	return &TodoHandler{
		App: app,
	}
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.App.TodoService.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	todo, err := h.App.TodoService.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.App.TodoService.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo domain.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.App.TodoService.UpdateTodoByID(id, &todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	err := h.App.TodoService.DeleteTodoByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}
	c.Status(http.StatusNoContent)
}
