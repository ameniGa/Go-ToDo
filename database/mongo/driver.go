package mongo

import (
	"context"
	"errors"
	"github.com/3almadmoon/ameni-assignment/config"
	entity "github.com/3almadmoon/ameni-assignment/entities"
	"github.com/3almadmoon/ameni-assignment/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

type MongoHandler struct {
	*mongo.Collection
}

var rwMutex sync.RWMutex

// NewMongoDBhandler creates a mongoDB handler
func NewMongoDBhandler(conf *config.Config) (*MongoHandler, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Database.Uri))
	if err != nil {
		return nil, err
	}
	log.Println("successfully connected to MongoDB")
	collection := client.Database(conf.Database.Name).Collection(conf.Database.Collection)

	return &MongoHandler{collection}, nil
}

// AddToDo adds a to do item to database
// returns error
func (db *MongoHandler) AddToDo(context context.Context, item entity.ToDo) error {
	err := helpers.CheckTimeout(context)
	if err != nil {
		return err
	}
	if (item == entity.ToDo{}) || helpers.IsEmpty(item.Hash) || helpers.IsEmpty(item.Title) {
		return errors.New("id or title should not be empty")
	}

	rwMutex.Lock()
	defer rwMutex.Unlock()
	_, err = db.InsertOne(context, item)
	if err != nil {
		return err
	}
	return nil
}

// DeleteToDo deletes item by hash from database
// returns boolean true:success, false:fail and error
func (db *MongoHandler) DeleteToDo(context context.Context, hash string) (bool, error) {
	err := helpers.CheckTimeout(context)
	if err != nil {
		return false, err
	}
	if helpers.IsEmpty(hash) {
		return false, errors.New("id or title should not be empty")
	}

	filter := bson.D{{Key: "hash", Value: hash}}
	rwMutex.Lock()
	res, err := db.DeleteOne(context, filter)
	rwMutex.Unlock()
	return handleResponse(err, res)
}

// UpdateToDo updates ,by hash, the status of to do item
// returns boolean true:success, false:fail and error
func (db *MongoHandler) UpdateToDo(context context.Context, hash string, status entity.EStatus) (bool, error) {
	err := helpers.CheckTimeout(context)
	if err != nil {
		return false, err
	}
	if helpers.IsEmpty(hash) {
		return false, errors.New("id or title should not be empty")
	}

	filter := bson.D{{Key: "hash", Value: hash}}
	log.Println(hash)
	update := bson.D{{"$set", bson.D{{Key: "status", Value: status}}}}
	rwMutex.Lock()
	res, err := db.UpdateOne(context, filter, update)
	log.Println(res, err)
	rwMutex.Unlock()
	return handleResponse(err, res)
}

// GetAllToDo finds all items in collection to do
// returns array pf To Do struct and error
func (db *MongoHandler) GetAllToDo(context context.Context, ch chan<- entity.ToDoWithError) {
	defer close(ch)

	err := helpers.CheckTimeout(context)
	if err != nil {
		ch <- entity.ToDoWithError{
			ToDo: nil,
			Err:  err,
		}
		return
	}
	cursor, err := db.Find(context, bson.D{})
	log.Println(err)
	if err != nil || cursor == nil{
		ch <- entity.ToDoWithError{
			ToDo: nil,
			Err:  err,
		}
		return
	}

	for cursor.Next(context) {
		var elem entity.ToDo
		err = cursor.Decode(&elem)
		if err != nil {
			ch <- entity.ToDoWithError{
				ToDo: nil,
				Err:  err,
			}
			continue
		}
		ch <- entity.ToDoWithError{
			ToDo: &elem,
			Err:  nil,
		}
	}
	cursor.Close(context)
	return
}

// handleResponse check query response
// returns boolean true:success, false:fail and error
func handleResponse(err error, res interface{}) (bool, error) {
	if err != nil {
		return false, err
	}
	var field int64
	switch r := res.(type) {
	case *mongo.DeleteResult:
		field = r.DeletedCount
	case *mongo.UpdateResult:
		field = r.MatchedCount
	default:
		field = 0
	}
	if field == 0 {
		return false, errors.New("item not found")
	}
	return true, nil
}
