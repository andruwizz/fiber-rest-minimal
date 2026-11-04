package todo

import (
	"math"

	"github.com/andruwizz/fiber-rest-minimal/internal/entity"
	"github.com/andruwizz/fiber-rest-minimal/internal/helper"
	"github.com/andruwizz/fiber-rest-minimal/internal/repository"
)

type TodoUsecase interface {
	Create(req *TodoCreateParam) (*entity.Todo, error)
	List(req *TodoListParam) (*entity.Pagination[entity.Todo], error)
	GetById(req *TodoGetParam) (*entity.Todo, error)
	Update(req *TodoUpdateParam) (*entity.Todo, error)
	Delete(req *TodoDeleteParam) error
}

type todoUsecase struct {
	repo repository.Todo
}

func NewTodoUsecase(repo repository.Todo) TodoUsecase {
	return &todoUsecase{repo}
}

func (u *todoUsecase) Create(req *TodoCreateParam) (*entity.Todo, error) {
	todo := &entity.Todo{
		Title: req.Title,
		Done:  req.Done,
	}

	err := u.repo.Create(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (u *todoUsecase) List(req *TodoListParam) (*entity.Pagination[entity.Todo], error) {
	todos := new([]entity.Todo)
	totalRows := new(int64)

	page := helper.GetPage(req.Page)
	limit := helper.GetLimit(req.Limit)

	err := u.repo.List(page, limit, totalRows, todos)
	if err != nil {
		return nil, err
	}

	totalPage := int(math.Ceil(float64(*totalRows) / float64(limit)))
	paginated := &entity.Pagination[entity.Todo]{
		Data:      *todos,
		Page:      page,
		Limit:     limit,
		TotalRows: *totalRows,
		TotalPage: totalPage,
	}
	return paginated, nil
}

func (u *todoUsecase) GetById(req *TodoGetParam) (*entity.Todo, error) {
	todo := new(entity.Todo)
	err := u.repo.GetById(req.Id, todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoUsecase) Update(req *TodoUpdateParam) (*entity.Todo, error) {
	todo := new(entity.Todo)
	err := u.repo.GetById(req.Id, todo)
	if err != nil {
		return nil, err
	}

	todo.Title = req.Title
	todo.Done = req.Done

	err = u.repo.Update(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoUsecase) Delete(req *TodoDeleteParam) error {
	todo := new(entity.Todo)
	err := u.repo.GetById(req.Id, todo)
	if err != nil {
		return err
	}
	return u.repo.Delete(todo)
}
