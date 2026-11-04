package repository

import (
	"errors"

	"github.com/andruwizz/fiber-rest-minimal/internal/entity"
	"github.com/andruwizz/fiber-rest-minimal/internal/helper"
	"github.com/andruwizz/fiber-rest-minimal/internal/model"

	"gorm.io/gorm"
)

type Todo interface {
	Create(todo *entity.Todo) error
	List(page, limit int) ([]entity.Todo, int64, error)
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
	if err := r.db.Create(m).Error; err != nil {
		return err
	}
	*todo = *m.ToTodoEntity()
	return nil
}

func (r *todoRepository) List(page, limit int) ([]entity.Todo, int64, error) {
	var total int64
	if err := r.db.Model(&model.Todo{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var models []model.Todo
	if err := r.db.Scopes(helper.Paginate(limit, page)).Find(&models).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	todos := make([]entity.Todo, len(models))
	for i, m := range models {
		todos[i] = *m.ToTodoEntity()
	}
	return todos, total, nil
}

func (r *todoRepository) GetById(id string, todo *entity.Todo) error {
	m := model.FromTodoEntity(todo)
	if err := r.db.Where("id = ?", id).Take(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.ErrRecordNotFound
		}
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
	*todo = *m.ToTodoEntity()
	return nil
}

func (r *todoRepository) Delete(todo *entity.Todo) error {
	m := model.FromTodoEntity(todo)
	return r.db.Delete(m).Error
}
