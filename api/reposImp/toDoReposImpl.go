package reposImp

import (
	"context"
	"fmt"
	"log"

	entity "github.com/3almadmoon/ameni-assignment/api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReposImp struct {
	Mcollection *mongo.Collection
}

func NewToDoRepos(collection *mongo.Collection) ReposImp {
	return ReposImp{collection}
}

func (s *ReposImp) AddToDo(context context.Context, item entity.ToDo) error {

	_, err := s.Mcollection.InsertOne(context, item)
	if err != nil {
		return err
	}
	return nil
}
func (s *ReposImp) DeleteToDo(context context.Context, id string) (bool, error) {
	objID, err := primitive.ObjectIDFromHex(id) // convert string id to objectId
	if err != nil {
		return false, err
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	res, erro := s.Mcollection.DeleteOne(context, filter)
	if erro != nil {
		return false, erro
	}
	if res.DeletedCount == 0 {
		return false, nil
	}
	return true, nil
}
func (s *ReposImp) UpdateToDo(context context.Context, id string, status entity.EStatus) (bool, error) {
	objID, err := primitive.ObjectIDFromHex(id) // convert string id to objectId
	if err != nil {
		return false, err
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{{"$set", bson.D{{Key: "status", Value: status}}}}
	res, erro := s.Mcollection.UpdateOne(context, filter, update)
	if erro != nil {
		return false, erro
	}
	if res.MatchedCount == 0 {
		return false, nil
	}
	fmt.Printf("update res %v", res)
	return true, nil

}
func (s *ReposImp) GetAllToDo(context context.Context) ([]*entity.ToDo, error) {
	var res []*entity.ToDo
	cursor, err := s.Mcollection.Find(context, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context) {
		var elem entity.ToDo
		erro := cursor.Decode(&elem)
		if erro != nil {
			log.Fatal(erro)
		}
		res = append(res, &elem)
	}
	cursor.Close(context)
	return res, nil

}
