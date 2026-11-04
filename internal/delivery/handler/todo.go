package handler

import (
	"github.com/andruwizz/fiber-rest-minimal/internal/delivery/request"
	"github.com/andruwizz/fiber-rest-minimal/internal/delivery/response"
	"github.com/andruwizz/fiber-rest-minimal/internal/usecase/todo"

	"github.com/gofiber/fiber/v3"
)

type TodoHandler interface {
	Create(c fiber.Ctx) error
	List(c fiber.Ctx) error
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
		return response.ErrorValidation(c, fiber.StatusBadRequest, err)
	}

	dto := &todo.TodoCreateParam{
		Title: req.Title,
		Done:  req.Done,
	}

	res, err := h.usecase.Create(dto)
	if err != nil {
		return response.ErrorResource(c, err)
	}

	return response.Resource(c, fiber.StatusCreated, res)
}

func (h *todoHandler) List(c fiber.Ctx) error {
	req := new(request.TodoList)
	if err := c.Bind().Query(req); err != nil {
		return response.ErrorValidation(c, fiber.StatusBadRequest, err)
	}

	dto := &todo.TodoListParam{
		Limit: req.Limit,
		Page:  req.Page,
	}

	res, err := h.usecase.List(dto)
	if err != nil {
		return response.ErrorResource(c, err)
	}

	return response.Collection(c, fiber.StatusOK, res)
}

func (h *todoHandler) Get(c fiber.Ctx) error {
	dto := &todo.TodoGetParam{
		Id: c.Params("id"),
	}

	res, err := h.usecase.GetById(dto)
	if err != nil {
		return response.ErrorResource(c, err)
	}

	return response.Resource(c, fiber.StatusOK, res)
}

func (h *todoHandler) Update(c fiber.Ctx) error {
	req := new(request.TodoUpdate)
	if err := c.Bind().Body(req); err != nil {
		return response.ErrorValidation(c, fiber.StatusBadRequest, err)
	}

	dto := &todo.TodoUpdateParam{
		Id:    c.Params("id"),
		Title: req.Title,
		Done:  req.Done,
	}

	res, err := h.usecase.Update(dto)
	if err != nil {
		return response.ErrorResource(c, err)
	}

	return response.Resource(c, fiber.StatusOK, res)
}

func (h *todoHandler) Delete(c fiber.Ctx) error {
	dto := &todo.TodoDeleteParam{
		Id: c.Params("id"),
	}

	err := h.usecase.Delete(dto)
	if err != nil {
		return response.ErrorResource(c, err)
	}

	return response.Empty(c, fiber.StatusOK)
}
