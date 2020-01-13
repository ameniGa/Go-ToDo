package database

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
	"strings"
	"sync"
	"time"
)

type MongoDBhandler struct {
	*mongo.Collection
}

var (
	rwMutex sync.RWMutex
	wg      sync.WaitGroup
)

//TODO change params with config
func NewMongoDBhandler(conf *config.Config) (*MongoDBhandler, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Database.Uri))
	if err != nil {
		return nil, err
	}
	log.Println("successfully connected to MongoDB")
	//TODO create index hash
	collection := client.Database(conf.Database.Name).Collection(conf.Database.Collection)
	return &MongoDBhandler{collection}, nil
}

//AddToDo add a todo item to database
//returns error
func (db *MongoDBhandler) AddToDo(context context.Context, item entity.ToDo) error {
	err := helpers.CheckTimeout(context)
	if err != nil {
		return err
	}
	log.Println("add ",item.Hash)
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

//DeleteToDo delete item by hash from database
//returns boolean true:success, false:fail and error
func (db *MongoDBhandler) DeleteToDo(context context.Context, hash string) (bool, error) {
	err := helpers.CheckTimeout(context)
	if err != nil {
		return false, err
	}
	if helpers.IsEmpty(hash) {
		return false, errors.New("id or title should not be empty")
	}
	log.Println("add ",hash)

	filter := bson.D{{Key: "hash", Value: hash}}
	rwMutex.Lock()
	res, err := db.DeleteOne(context, filter)
	rwMutex.Unlock()
	return handleResponse(err, res)
}

//UpdateToDo update ,by hash, the status of todo item
//returns boolean true:success, false:fail and error
func (db *MongoDBhandler) UpdateToDo(context context.Context, hash string, status entity.EStatus) (bool, error) {
	err := helpers.CheckTimeout(context)
	if err != nil {
		return false, err
	}
	if helpers.IsEmpty(hash) {
		return false, errors.New("id or title should not be empty")
	}
	log.Println("add ",hash)

	filter := bson.D{{Key: "hash", Value: hash}}
	log.Println(hash)
	update := bson.D{{"$set", bson.D{{Key: "status", Value: status}}}}
	rwMutex.Lock()
	res, err := db.UpdateOne(context, filter, update)
	log.Println(res, err)
	rwMutex.Unlock()
	return handleResponse(err, res)
}

//GetAllToDo finds all items in collection todo
//returns array pf ToDo struct and error
func (db *MongoDBhandler) GetAllToDo(context context.Context) ([]*entity.ToDo, error) {
	err := helpers.CheckTimeout(context)
	if err != nil {
		return nil, err
	}
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	cursor, err := db.Find(context, bson.D{})
	if err != nil {
		return nil, err
	}
	var res []*entity.ToDo
	var errs []string
	for cursor.Next(context) {
		var elem entity.ToDo
		err = cursor.Decode(&elem)
		if err != nil {
			errs = append(errs, err.Error())
		}
		res = append(res, &elem)
	}
	cursor.Close(context)
	if len(errs) > 0 {
		return res, errors.New(strings.Join(errs, "; "))
	}
	return res, nil
}

//handleResponse check query response
//returns boolean true:success, false:fail and error
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
