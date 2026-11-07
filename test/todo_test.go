package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"

	"testing"

	"github.com/andruwizz/fiber-rest-minimal/internal/config"
	"github.com/andruwizz/fiber-rest-minimal/internal/delivery/handler"
	"github.com/andruwizz/fiber-rest-minimal/internal/delivery/request"
	"github.com/andruwizz/fiber-rest-minimal/internal/delivery/response"
	"github.com/andruwizz/fiber-rest-minimal/internal/delivery/routes"
	"github.com/andruwizz/fiber-rest-minimal/internal/entity"
	"github.com/andruwizz/fiber-rest-minimal/internal/usecase/todo"
	"github.com/andruwizz/fiber-rest-minimal/internal/usecase/todo/mock"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func NewTestSetup() (*fiber.App, todo.TodoUsecase) {
	todoService := mock.NewMockTodoUsecase()
	todoHandler := handler.NewTodoHandler(todoService)

	app := config.NewFiberApp()
	routes.RegisterTodoRoutes(app, todoHandler)

	return app, todoService
}

func TestCreateTodo(t *testing.T) {
	// Arrange
	app, _ := NewTestSetup()

	payload := request.TodoCreate{Title: "Test Todo", Done: true}
	reqPayload, _ := json.Marshal(payload)

	// Act
	req := httptest.NewRequest(fiber.MethodPost, "/todos", bytes.NewReader(reqPayload))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	// Assert
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var res response.DataResponse[entity.Todo]
	json.Unmarshal(body, &res)

	assert.NoError(t, resp.Body.Close())
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	assert.Equal(t, payload.Title, res.Data.Title)
	assert.Equal(t, payload.Done, res.Data.Done)
	assert.NotNil(t, res.Data.Id)
	assert.NotNil(t, res.Data.CreatedAt)
	assert.NotNil(t, res.Data.UpdatedAt)
}

func TestGetTodo(t *testing.T) {
	// Arrange
	app, mockUsecase := NewTestSetup()
	dto := &todo.TodoCreateParam{Title: "Test Todo", Done: true}
	res, _ := mockUsecase.Create(dto)

	// Act
	req := httptest.NewRequest(fiber.MethodGet, "/todos/"+res.Id, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	// Assert
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var resData response.DataResponse[entity.Todo]
	json.Unmarshal(body, &resData)

	assert.NoError(t, resp.Body.Close())
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.Equal(t, dto.Title, resData.Data.Title)
	assert.Equal(t, dto.Done, resData.Data.Done)
	assert.NotNil(t, resData.Data.Id)
	assert.NotNil(t, resData.Data.CreatedAt)
	assert.NotNil(t, resData.Data.UpdatedAt)
}

func TestUpdateTodo(t *testing.T) {
	// Arrange
	app, mockUsecase := NewTestSetup()
	dto := &todo.TodoCreateParam{Title: "Test Todo", Done: true}
	res, _ := mockUsecase.Create(dto)

	updatePayload := request.TodoUpdate{Title: "Updated Todo", Done: false}
	reqPayload, _ := json.Marshal(updatePayload)

	// Act
	req := httptest.NewRequest(fiber.MethodPut, "/todos/"+res.Id, bytes.NewReader(reqPayload))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	// Assert
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var resData response.DataResponse[entity.Todo]
	json.Unmarshal(body, &resData)

	assert.NoError(t, resp.Body.Close())
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.Equal(t, updatePayload.Title, resData.Data.Title)
	assert.Equal(t, updatePayload.Done, resData.Data.Done)
	assert.NotNil(t, resData.Data.Id)
	assert.NotNil(t, resData.Data.CreatedAt)
	assert.NotNil(t, resData.Data.UpdatedAt)
}
func TestDeleteTodo(t *testing.T) {
	// Arrange
	app, mockUsecase := NewTestSetup()
	dto := &todo.TodoCreateParam{Title: "Test Todo", Done: true}
	res, _ := mockUsecase.Create(dto)

	// Act
	req := httptest.NewRequest(fiber.MethodDelete, "/todos/"+res.Id, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	// Assert
	var resData response.EmptyResponse

	assert.NoError(t, resp.Body.Close())
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NotNil(t, resData.Success)
}
