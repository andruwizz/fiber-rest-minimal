package mock

import (
	"math"
	"time"

	"github.com/andruwizz/fiber-rest-minimal/internal/entity"
	"github.com/andruwizz/fiber-rest-minimal/internal/helper"
	"github.com/andruwizz/fiber-rest-minimal/internal/usecase/todo"
	"github.com/google/uuid"
)

type TodoUsecase interface {
	Create(req *todo.TodoCreateParam) (*entity.Todo, error)
	List(req *todo.TodoListParam) (*entity.Pagination[entity.Todo], error)
	GetById(req *todo.TodoGetParam) (*entity.Todo, error)
	Update(req *todo.TodoUpdateParam) (*entity.Todo, error)
	Delete(req *todo.TodoDeleteParam) error
}

type mockTodoUsecase struct {
	db map[string]entity.Todo
}

func NewMockTodoUsecase() TodoUsecase {
	return &mockTodoUsecase{
		db: make(map[string]entity.Todo),
	}
}

func (u *mockTodoUsecase) Create(req *todo.TodoCreateParam) (*entity.Todo, error) {
	id := uuid.New().String()
	timeNow := time.Now().Unix()

	todo := &entity.Todo{
		Id:        id,
		Title:     req.Title,
		Done:      req.Done,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	u.db[id] = *todo
	return todo, nil
}

func (u *mockTodoUsecase) List(req *todo.TodoListParam) (*entity.Pagination[entity.Todo], error) {
	page := helper.GetPage(req.Page)
	limit := helper.GetLimit(req.Limit)

	todos := []entity.Todo{}
	for _, v := range u.db {
		todos = append(todos, v)
	}

	offset := (page - 1) * limit
	totalPage := int(math.Ceil(float64(len(todos)) / float64(limit)))

	pagination := &entity.Pagination[entity.Todo]{
		Data:       todos[offset:min(offset+limit, len(todos))],
		Page:       page,
		Limit:      limit,
		TotalRows:  int64(len(todos)),
		TotalPages: int64(totalPage),
	}

	return pagination, nil
}

func (u *mockTodoUsecase) GetById(req *todo.TodoGetParam) (*entity.Todo, error) {
	found, ok := u.db[req.Id]
	if !ok {
		return nil, entity.ErrRecordNotFound
	}
	todo := &found

	return todo, nil
}

func (u *mockTodoUsecase) Update(req *todo.TodoUpdateParam) (*entity.Todo, error) {
	found, ok := u.db[req.Id]
	if !ok {
		return nil, entity.ErrRecordNotFound
	}
	todo := &found

	todo.Title = req.Title
	todo.Done = req.Done
	todo.UpdatedAt = time.Now().Unix()

	u.db[req.Id] = *todo

	return todo, nil
}

func (u *mockTodoUsecase) Delete(req *todo.TodoDeleteParam) error {
	_, ok := u.db[req.Id]
	if !ok {
		return entity.ErrRecordNotFound
	}
	delete(u.db, req.Id)

	return nil
}
