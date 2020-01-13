package database

import (
	"context"
	"github.com/3almadmoon/ameni-assignment/config"
	entity"github.com/3almadmoon/ameni-assignment/entities"
	"log"
)

type DBhandler interface {
	AddToDo(ctx context.Context, item entity.ToDo) error
	DeleteToDo(ctx context.Context, hash string) (bool, error)
	UpdateToDo(ctx context.Context, hash string, status entity.EStatus) (bool, error)
	GetAllToDo(ctx context.Context) ([]*entity.ToDo, error)
}

func CreateDBhandler(config *config.Config) (*DBhandler,error) {
	var db DBhandler
	var err error
	switch config.Database.Type {
	case "mongo" :
		db, err = NewMongoDBhandler(config)
		if err != nil {
			return nil, err
		}
	default:
		log.Panicf("%v not supported",config.Database.Type)
	}
	return &db,nil
}