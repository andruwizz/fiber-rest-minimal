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

type TodoGetParam struct {
	Id string
}

type TodoDeleteParam struct {
	Id string
}
