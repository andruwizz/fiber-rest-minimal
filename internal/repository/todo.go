package repository

import (
	"github.com/andruwizz/fiber-rest-minimal/internal/entity"
	"github.com/andruwizz/fiber-rest-minimal/internal/helper"
	"github.com/andruwizz/fiber-rest-minimal/internal/model"

	"gorm.io/gorm"
)

type Todo interface {
	Create(todo *entity.Todo) error
	List(page, limit int, total *int64, todos *[]entity.Todo) error
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

func (r *todoRepository) List(page, limit int, total *int64, todos *[]entity.Todo) error {
	if err := r.db.Model(&model.Todo{}).Count(total).Error; err != nil {
		return err
	}

	if err := r.db.Scopes(helper.Paginate(limit, page)).Find(todos).Error; err != nil {
		return err
	}

	return nil
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
