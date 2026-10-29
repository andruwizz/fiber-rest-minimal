package todo

import (
	"go-something/internal/entity"
	"go-something/internal/repository"
)

type TodoUsecase interface {
	Create(todo *TodoCreateParam) error
	GetById(req *TodoGetByIdParam) (*entity.Todo, error)
	Update(todo *TodoUpdateParam) error
	Delete(todo *TodoDeleteParam) error
}

type todoUsecase struct {
	repo repository.Todo
}

func NewTodoUsecase(repo repository.Todo) TodoUsecase {
	return &todoUsecase{repo}
}

func (u *todoUsecase) Create(req *TodoCreateParam) error {
	todo := &entity.Todo{
		Title: req.Title,
		Done:  req.Done,
	}

	err := u.repo.Create(todo)
	if err != nil {
		return err
	}

	return err
}

func (u *todoUsecase) GetById(req *TodoGetByIdParam) (*entity.Todo, error) {
	todo := new(entity.Todo)
	err := u.repo.GetById(req.ID, todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoUsecase) Update(req *TodoUpdateParam) error {
	todo := new(entity.Todo)
	err := u.repo.GetById(req.Id, todo)
	if err != nil {
		return err
	}

	todo.Title = req.Title
	todo.Done = req.Done

	err = u.repo.Update(todo)
	if err != nil {
		return err
	}
	return nil
}

func (u *todoUsecase) Delete(req *TodoDeleteParam) error {
	todo := new(entity.Todo)
	err := u.repo.GetById(req.ID, todo)
	if err != nil {
		return err
	}
	return u.repo.Delete(todo)
}
