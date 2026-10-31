package request

type TodoCreate struct {
	Title string `json:"title" form:"title" validate:"required"`
	Done  bool   `json:"done" form:"done" validate:"required"`
}

type TodoGet struct {
	Id string `json:"id" form:"id" validate:"required"`
}

type TodoUpdate struct {
	Id    string `json:"id" form:"id" validate:"required"`
	Title string `json:"title" form:"title" validate:"required"`
	Done  bool   `json:"done" form:"done" validate:"required"`
}

type TodoDelete struct {
	Id string `json:"id" form:"id" validate:"required"`
}
