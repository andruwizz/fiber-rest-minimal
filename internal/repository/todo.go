package repository

import (
	"go-something/internal/entity"
	"go-something/internal/model"

	"gorm.io/gorm"
)

type Todo interface {
	Create(todo *entity.Todo) error
	GetById(id string, todo *entity.Todo) error
	Update(todo *entity.Todo) error
	Delete(todo *entity.Todo) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) Todo {
	return &todoRepository{db}
}

func (r *todoRepository) Create(todo *entity.Todo) error {
	m := model.FromTodoEntity(todo)
	return r.db.Create(m).Error
}

func (r *todoRepository) GetById(id string, todo *entity.Todo) error {
	m := model.FromTodoEntity(todo)
	if err := r.db.Where("id = ?", id).Take(m).Error; err != nil {
		return err
	}
	*todo = *m.ToTodoEntity()
	return nil
}

func (r *todoRepository) Update(todo *entity.Todo) error {
	m := model.FromTodoEntity(todo)
	if err := r.db.Save(m).Error; err != nil {
		return err
	}
	return nil
}

func (r *todoRepository) Delete(todo *entity.Todo) error {
	m := model.FromTodoEntity(todo)
	return r.db.Delete(m).Error
}
