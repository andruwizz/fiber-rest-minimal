package todo

type TodoCreateParam struct {
	Title string
	Done  bool
}

type TodoUpdateParam struct {
	Id    string
	Title string
	Done  bool
}

type TodoGetByIdParam struct {
	ID string
}

type TodoDeleteParam struct {
	ID string
}
