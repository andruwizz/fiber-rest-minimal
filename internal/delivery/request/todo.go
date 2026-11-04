package request

type TodoCreate struct {
	Title string `json:"title" form:"title" validate:"required"`
	Done  bool   `json:"done" form:"done,default=false" validate:"boolean"`
}

type TodoList struct {
	Limit int `json:"limit" query:"limit" validate:"numeric,max=100"`
	Page  int `json:"page" query:"page" validate:"numeric,max=100"`
}

type TodoGet struct {
	Id string `uri:"id" validate:"required"`
}

type TodoUpdate struct {
	Id    string `uri:"id" validate:"required"`
	Title string `json:"title" form:"title" validate:"required"`
	Done  bool   `json:"done" form:"done,default=false" validate:"boolean"`
}

type TodoDelete struct {
	Id string `uri:"id" validate:"required"`
}
