package response

import (
	"github.com/andruwizz/fiber-rest-minimal/internal/entity"
	"github.com/gofiber/fiber/v3"
)

type ListMeta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type DataResponse[T entityType] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type ListResponse[T entityType] struct {
	Success bool     `json:"success"`
	Data    []T      `json:"data"`
	Meta    ListMeta `json:"meta"`
}

type EmptyResponse struct {
	Success bool `json:"success"`
}

type entityType interface {
	any
}

func Empty(ctx fiber.Ctx, code int) error {
	res := EmptyResponse{
		Success: true,
	}

	return ctx.Status(code).JSON(res)
}

func Resource[T entityType](ctx fiber.Ctx, code int, e *T) error {
	res := DataResponse[T]{
		Success: true,
		Data:    *e,
	}

	return ctx.Status(code).JSON(res)
}

func Collection[T entityType](ctx fiber.Ctx, code int, p *entity.Pagination[T]) error {
	col := ListResponse[T]{
		Success: true,
		Data:    p.Data,
		Meta: ListMeta{
			Page:       p.Page,
			PerPage:    p.Limit,
			Total:      int(p.TotalRows),
			TotalPages: int(p.TotalPages),
		},
	}

	return ctx.Status(code).JSON(col)
}
