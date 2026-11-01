package entity

type entityType interface {
	Todo
}

type Pagination[T entityType] struct {
	Data      []T   `json:"data"`
	Limit     int   `json:"limit"`
	Page      int   `json:"page"`
	TotalRows int64 `json:"total_rows"`
	TotalPage int   `json:"total_page"`
}
