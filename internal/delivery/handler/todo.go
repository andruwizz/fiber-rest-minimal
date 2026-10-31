package handler

import (
	"go-something/internal/delivery/request"
	"go-something/internal/usecase/todo"

	"github.com/gofiber/fiber/v3"
)

type TodoHandler interface {
	Create(c fiber.Ctx) error
	Get(c fiber.Ctx) error
	Update(c fiber.Ctx) error
	Delete(c fiber.Ctx) error
}

type todoHandler struct {
	usecase todo.TodoUsecase
}

func NewTodoHandler(u todo.TodoUsecase) TodoHandler {
	return &todoHandler{usecase: u}
}

func (h *todoHandler) Create(c fiber.Ctx) error {
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

func (h *todoHandler) Get(c fiber.Ctx) error {
	req := new(request.TodoGet)
	if err := c.Bind().URI(req); err != nil {
		return err
	}

	dto := &todo.TodoGetParam{
		Id: req.Id,
	}

	res, err := h.usecase.GetById(dto)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *todoHandler) Update(c fiber.Ctx) error {
	req := new(request.TodoUpdate)
	req.Id = c.Params("id")

	if err := c.Bind().Body(req); err != nil {
		return err
	}

	dto := &todo.TodoUpdateParam{
		Id:    req.Id,
		Title: req.Title,
		Done:  req.Done,
	}

	err := h.usecase.Update(dto)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *todoHandler) Delete(c fiber.Ctx) error {
	req := new(request.TodoDelete)
	if err := c.Bind().URI(req); err != nil {
		return err
	}

	dto := &todo.TodoDeleteParam{
		Id: req.Id,
	}

	err := h.usecase.Delete(dto)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
