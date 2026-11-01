package todo

type TodoCreateParam struct {
	Title string
	Done  bool
}

type TodoGetParam struct {
	Id string
}

type TodoUpdateParam struct {
	Id    string
	Title string
	Done  bool
}

type TodoDeleteParam struct {
	Id string
}
