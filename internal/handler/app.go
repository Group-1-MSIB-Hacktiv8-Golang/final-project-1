package handler

import (
	"final-project-1/internal/app"
	"final-project-1/internal/repository"
	"final-project-1/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func StartApp() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Gin Gonic router
	r := gin.Default()

	// Initialize database connection
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbDialect := os.Getenv("DB_DIALECT")

	db, err := repository.InitDB(dbHost, dbPort, dbUser, dbPassword, dbName, dbDialect)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	// Initialize repositories and services
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)

	// Initialize application
	app := app.NewApp(todoService)

	docs.SwaggerInfo.Title = "Todo Application"
	docs.SwaggerInfo.Description = "Todo Application API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "https://final-project-1-production.up.railway.app/"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Initialize API endpoints
	todoHandler := NewTodoHandler(app)
	r.GET("/todos", todoHandler.GetTodos)
	r.GET("/todos/:id", todoHandler.GetTodoByID)
	r.POST("/todos", todoHandler.CreateTodo)
	r.PUT("/todos/:id", todoHandler.UpdateTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)

	// Start the server
	r.Run()
}
