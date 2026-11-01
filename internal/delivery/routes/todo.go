package routes

import (
	"go-something/internal/delivery/handler"

	"github.com/gofiber/fiber/v3"
)

func RegisterTodoRoutes(r *fiber.App, h handler.TodoHandler) {
	todo := r.Group("/todos")
	todo.Post("/", h.Create)
	todo.Get("/", h.List)
	todo.Get("/:id", h.Get)
	todo.Put("/:id", h.Update)
	todo.Delete("/:id", h.Delete)
}
