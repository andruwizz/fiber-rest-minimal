package mock

import (
	"math"

	"github.com/andruwizz/fiber-rest-minimal/internal/entity"
	"github.com/andruwizz/fiber-rest-minimal/internal/helper"
	"github.com/andruwizz/fiber-rest-minimal/internal/usecase/todo"
)

type TodoUsecase interface {
	Create(req *todo.TodoCreateParam) (*entity.Todo, error)
	List(req *todo.TodoListParam) (*entity.Pagination[entity.Todo], error)
	GetById(req *todo.TodoGetParam) (*entity.Todo, error)
	Update(req *todo.TodoUpdateParam) (*entity.Todo, error)
	Delete(req *todo.TodoDeleteParam) error
}

type mockTodoUsecase struct{}

func NewMockTodoUsecase() TodoUsecase {
	return &mockTodoUsecase{}
}

func (u *mockTodoUsecase) Create(req *todo.TodoCreateParam) (*entity.Todo, error) {
	todo := &entity.Todo{
		Id:        "398c1df7-64b1-40ac-9a04-3599b7f9a501",
		Title:     req.Title,
		Done:      req.Done,
		CreatedAt: 1793774478,
		UpdatedAt: 1793774478,
	}

	return todo, nil
}

func (u *mockTodoUsecase) List(req *todo.TodoListParam) (*entity.Pagination[entity.Todo], error) {
	page := helper.GetPage(req.Page)
	limit := helper.GetLimit(req.Limit)

	todos := []entity.Todo{
		{
			Id:        "398c1df7-64b1-40ac-9a04-3599b7f9a501",
			Title:     "Sample Todo 1",
			Done:      false,
			CreatedAt: 1793774478,
			UpdatedAt: 1793774478,
		},
		{
			Id:        "498c1df7-64b1-40ac-9a04-3599b7f9a502",
			Title:     "Sample Todo 2",
			Done:      true,
			CreatedAt: 1793774478,
			UpdatedAt: 1793774478,
		},
		{
			Id:        "598c1df7-64b1-40ac-9a04-3599b7f9a503",
			Title:     "Sample Todo 3",
			Done:      false,
			CreatedAt: 1793774478,
			UpdatedAt: 1793774478,
		},
		{
			Id:        "698c1df7-64b1-40ac-9a04-3599b7f9a504",
			Title:     "Sample Todo 4",
			Done:      true,
			CreatedAt: 1793774478,
			UpdatedAt: 1793774478,
		},
		{
			Id:        "798c1df7-64b1-40ac-9a04-3599b7f9a505",
			Title:     "Sample Todo 5",
			Done:      false,
			CreatedAt: 1793774478,
			UpdatedAt: 1793774478,
		},
		{
			Id:        "898c1df7-64b1-40ac-9a04-3599b7f9a506",
			Title:     "Sample Todo 6",
			Done:      true,
			CreatedAt: 1793774478,
			UpdatedAt: 1793774478,
		},
		{
			Id:        "998c1df7-64b1-40ac-9a04-3599b7f9a507",
			Title:     "Sample Todo 7",
			Done:      false,
			CreatedAt: 1793774478,
			UpdatedAt: 1793774478,
		},
		{
			Id:        "a98c1df7-64b1-40ac-9a04-3599b7f9a508",
			Title:     "Sample Todo 8",
			Done:      true,
			CreatedAt: 1793774478,
			UpdatedAt: 1793774478,
		},
		{
			Id:        "b98c1df7-64b1-40ac-9a04-3599b7f9a509",
			Title:     "Sample Todo 9",
			Done:      false,
			CreatedAt: 1793774478,
			UpdatedAt: 1793774478,
		},
		{
			Id:        "c98c1df7-64b1-40ac-9a04-3599b7f9a510",
			Title:     "Sample Todo 10",
			Done:      true,
			CreatedAt: 1793774478,
			UpdatedAt: 1793774478,
		},
		{
			Id:        "d98c1df7-64b1-40ac-9a04-3599b7f9a511",
			Title:     "Sample Todo 11",
			Done:      false,
			CreatedAt: 1793774478,
			UpdatedAt: 1793774478,
		},
		{
			Id:        "e98c1df7-64b1-40ac-9a04-3599b7f9a512",
			Title:     "Sample Todo 12",
			Done:      true,
			CreatedAt: 1793774478,
			UpdatedAt: 1793774478,
		},
	}

	offset := (page - 1) * limit
	totalPage := int(math.Ceil(float64(len(todos)) / float64(limit)))

	pagination := &entity.Pagination[entity.Todo]{
		Data:       todos[offset : offset+limit],
		Page:       page,
		Limit:      limit,
		TotalRows:  int64(len(todos)),
		TotalPages: int64(totalPage),
	}

	return pagination, nil
}

func (u *mockTodoUsecase) GetById(req *todo.TodoGetParam) (*entity.Todo, error) {
	todo := &entity.Todo{
		Id:        "398c1df7-64b1-40ac-9a04-3599b7f9a501",
		Title:     "Sample Todo",
		Done:      false,
		CreatedAt: 1793774478,
		UpdatedAt: 1793774478,
	}

	return todo, nil
}

func (u *mockTodoUsecase) Update(req *todo.TodoUpdateParam) (*entity.Todo, error) {
	todo := &entity.Todo{
		Id:        "398c1df7-64b1-40ac-9a04-3599b7f9a501",
		Title:     req.Title,
		Done:      req.Done,
		CreatedAt: 1793774478,
		UpdatedAt: 1793774478,
	}

	return todo, nil
}

func (u *mockTodoUsecase) Delete(req *todo.TodoDeleteParam) error {
	return nil
}
