package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/3almadmoon/ameni-assignment/config"
	"github.com/3almadmoon/ameni-assignment/database/mongo"
	entity "github.com/3almadmoon/ameni-assignment/entities"
)

// DBhandler holds the functions of database stuff
type DBhandler interface {
	// AddToDo adds a to do item to database
	AddToDo(ctx context.Context, item entity.ToDo) error
	// DeleteToDo deletes item by hash from database
	DeleteToDo(ctx context.Context, hash string) (bool, error)
	// UpdateToDo updates ,by hash, the status of to do item
	UpdateToDo(ctx context.Context, hash string, status entity.EStatus) (bool, error)
	// GetAllToDo returns all to do items
	GetAllToDo(ctx context.Context, ch chan<- entity.ToDoWithError)
}

// CreateDBhandler creates a database handler from the given database type
func CreateDBhandler(config *config.Config) (DBhandler,error) {
	var db DBhandler
    var err error
	switch config.Database.Type {
	case "mongo" :
		db,err = mongo.NewMongoDBhandler(config)
		if err != nil{
			return nil,err
		}
	default:
		return nil, errors.New(fmt.Sprintf("%v not supported",config.Database.Type))
	}
	return db,nil
}