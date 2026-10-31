package handler

import (
	"go-something/internal/delivery/request"
	"go-something/internal/usecase/todo"

	"github.com/gofiber/fiber/v3"
)

type TodoHandler struct {
	usecase todo.TodoUsecase
}

func NewTodoHandler(u todo.TodoUsecase) *TodoHandler {
	return &TodoHandler{usecase: u}
}

func (h *TodoHandler) Create(c fiber.Ctx) error {
	req := new(request.TodoCreate)
	if err := c.Bind().Body(req); err != nil {
		return err
	}

	dto := &todo.TodoCreateParam{
		Title: req.Title,
		Done:  req.Done,
	}

	err := h.usecase.Create(dto)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}
