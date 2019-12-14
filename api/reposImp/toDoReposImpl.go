package reposImp

import (
	"context"
	"sync"

	entity "github.com/3almadmoon/ameni-assignment/api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReposImp struct {
	Mcollection *mongo.Collection
}

var mutex sync.Mutex
var wg sync.WaitGroup

//AddToDo add a todo item to database
//returns error
func (s *ReposImp) AddToDo(context context.Context, item entity.ToDo) error {
	mutex.Lock()
	_, err := s.Mcollection.InsertOne(context, item)
	mutex.Unlock()
	if err != nil {
		return err
	}
	return nil
}

//DeleteToDo delete item by hash from database
//returns boolean true:success, false:fail and error
func (s *ReposImp) DeleteToDo(context context.Context, hash string) (bool, error) {
	filter := bson.D{{Key: "hash", Value: hash}}
	//mutex.Lock()
	res, erro := s.Mcollection.DeleteOne(context, filter)
	//mutex.Unlock()
	return handleResponse(erro, res.DeletedCount)
}

//UpdateToDo update ,by hash, the status of todo item
//returns boolean true:success, false:fail and error
func (s *ReposImp) UpdateToDo(context context.Context, hash string, status entity.EStatus) (bool, error) {
	filter := bson.D{{Key: "hash", Value: hash}}
	update := bson.D{{"$set", bson.D{{Key: "status", Value: status}}}}
	//mutex.Lock()
	res, erro := s.Mcollection.UpdateOne(context, filter, update)
	//mutex.Unlock()
	return handleResponse(erro, res.MatchedCount)
}

//GetAllToDo finds all items in collection todo
//returns array pf ToDo struct and error
func (s *ReposImp) GetAllToDo(context context.Context) ([]*entity.ToDo, error) {
	var res []*entity.ToDo
	mutex.Lock()
	cursor, err := s.Mcollection.Find(context, bson.D{})
	mutex.Unlock()
	if err != nil {
		return nil, err
	}
	for cursor.Next(context) {
		var elem entity.ToDo
		erro := cursor.Decode(&elem)
		if erro != nil {
			return nil, erro
		}
		res = append(res, &elem)
	}
	cursor.Close(context)
	return res, nil

}

//handleResponse check query response
//returns boolean true:success, false:fail and error
func handleResponse(err error, resParam int64) (bool, error) {
	if err != nil {
		return false, err
	}
	if resParam == 0 {
		return false, nil
	}
	return true, nil
}
