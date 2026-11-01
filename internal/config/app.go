package config

import (
	"github.com/andruwizz/fiber-rest-minimal/internal/delivery/handler"
	"github.com/andruwizz/fiber-rest-minimal/internal/delivery/routes"
	"github.com/andruwizz/fiber-rest-minimal/internal/repository"
	"github.com/andruwizz/fiber-rest-minimal/internal/usecase/todo"

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
