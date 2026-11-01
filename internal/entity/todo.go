package entity

type Todo struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Done      bool   `json:"done"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type PaginatedTodo struct {
	Data      []Todo `json:"data"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	TotalRows int64  `json:"total_rows"`
	TotalPage int    `json:"total_page"`
}
