package entity

type entityType interface {
	any
}

type Pagination[T entityType] struct {
	Data       []T   `json:"data"`
	Limit      int   `json:"limit"`
	Page       int   `json:"page"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int64 `json:"total_pages"`
}
