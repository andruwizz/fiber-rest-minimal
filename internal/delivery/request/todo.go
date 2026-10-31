package request

type TodoCreate struct {
	Title string `json:"title" form:"title" validate:"required"`
	Done  bool   `json:"done" form:"done" validate:"required"`
}
