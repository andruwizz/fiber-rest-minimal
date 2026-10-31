package config

import (
	"go-something/internal/delivery/handler"
	"go-something/internal/delivery/routes"
	"go-something/internal/repository"
	"go-something/internal/usecase/todo"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB  *gorm.DB
	App *fiber.App
}

func Bootstrap(config *BootstrapConfig) {
	// Setup Repository
	todoRepo := repository.NewTodoRepository(config.DB)

	// Setup Services / Usecase
	todoService := todo.NewTodoUsecase(todoRepo)

	// Setup Handler
	todoHandler := handler.NewTodoHandler(todoService)

	// Setup Router
	routes.RegisterTodoRoutes(config.App, todoHandler)

}
