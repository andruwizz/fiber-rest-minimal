package model

import (
	"github.com/andruwizz/fiber-rest-minimal/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	Id        string `gorm:"type:uuid;primaryKey"`
	Title     string `gorm:"not null"`
	Done      bool   `gorm:"default:false"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
}

func (t *Todo) BeforeCreate(tx *gorm.DB) error {
	newUUID := uuid.New()
	t.Id = newUUID.String()

	return nil
}

func FromTodoEntity(e *entity.Todo) *Todo {
	return &Todo{
		Id:        e.Id,
		Title:     e.Title,
		Done:      e.Done,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
func (t *Todo) ToTodoEntity() *entity.Todo {
	return &entity.Todo{
		Id:        t.Id,
		Title:     t.Title,
		Done:      t.Done,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
