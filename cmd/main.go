package main

import (
	"final-project-1/internal/app"
	"final-project-1/internal/handler"
	"final-project-1/internal/repository"
	"final-project-1/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Gin Gonic router
	r := gin.Default()

	// Initialize database connection
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbDialect := os.Getenv("DB_DIALECT")

	// DB_HOST=localhost
	// DB_PORT=5432
	// DB_USER=postgres
	// DB_PASSWORD=root
	// DB_NAME=todolist
	// DB_DIALECT=postgres

	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "root"
	dbName := "todolist"
	dbDialect := "postgres"

	db, err := repository.InitDB(dbHost, dbPort, dbUser, dbPassword, dbName, dbDialect)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	// Initialize repositories and services
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)

	// Initialize application
	app := app.NewApp(todoService)

	// Initialize API endpoints
	todoHandler := handler.NewTodoHandler(app)
	r.GET("/todos", todoHandler.GetTodos)
	r.GET("/todos/:id", todoHandler.GetTodoByID)
	r.POST("/todos", todoHandler.CreateTodo)
	r.PUT("/todos/:id", todoHandler.UpdateTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)

	// Start the server
	r.Run()
}
