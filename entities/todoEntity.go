package entities

type EStatus int

const (
	UNKNOWN    EStatus = 0
	TODO       EStatus = 1
	INPROGRESS EStatus = 2
	DONE       EStatus = 3
)

type ToDo struct {
	Hash        string  `bson:"hash"`
	Title       string  `bson:"title"`
	Description string  `bson:"description"`
	Status      EStatus `bson:"status"`
}

type StatusWithError struct {
	Status bool
	Err    error
}

type ToDoWithError struct {
	ToDo *ToDo
	Err  error
}
