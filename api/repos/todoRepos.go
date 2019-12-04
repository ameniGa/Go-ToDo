package repos

import (
	"context"

	entity "github.com/3almadmoon/ameni-assignment/api/entities"
)

//ToDoRepos interface of the functions to implement in buisness logic
type ToDoRepos interface {
	AddToDo(context context.Context, item entity.ToDo) error
	DeleteToDo(context context.Context, hash string) (bool, error)
	UpdateToDo(context context.Context, hash string, status entity.EStatus) (bool, error)
	GetAllToDo(context context.Context) ([]*entity.ToDo, error)
}
