package request

type TodoCreate struct {
	Title string `json:"title" form:"title" validate:"required"`
	Done  bool   `json:"done" form:"done,default=false" validate:"boolean"`
}

type TodoList struct {
	Limit int `json:"limit" query:"limit" validate:"numeric"`
	Page  int `json:"page" query:"page" validate:"numeric"`
}

type TodoGet struct {
	Id string `json:"id" form:"id" query:"id" validate:"required"`
}

type TodoUpdate struct {
	Id    string `json:"id" form:"id" query:"id" validate:"required"`
	Title string `json:"title" form:"title" validate:"required"`
	Done  bool   `json:"done" form:"done,default=false" validate:"boolean"`
}

type TodoDelete struct {
	Id string `json:"id" form:"id" query:"id" validate:"required"`
}
