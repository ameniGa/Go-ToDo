package repos

import (
	"context"

	entity "github.com/3almadmoon/ameni-assignment/api/entities"
)

type ToDoRepos interface {
	AddToDo(context context.Context, item entity.ToDo) error
	DeleteToDo(context context.Context, id string) (bool, error)
	UpdateToDo(context context.Context, id string, status entity.EStatus) (bool, error)
	GetAllToDo(context context.Context) ([]*entity.ToDo, error)
}
