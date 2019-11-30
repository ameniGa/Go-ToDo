package reposImp

import (
	"context"
	"testing"

	"github.com/3almadmoon/ameni-assignment/api/db"
	entity "github.com/3almadmoon/ameni-assignment/api/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var reposImp ReposImp
var mContext context.Context

func init() {
	collection, _ = db.Connect()
	reposImp = ReposImp{collection}
}

var list []*entity.ToDo
var err error

func TestAddToDo(t *testing.T) {
	t.Log("***********ADD******************")
	err := reposImp.AddToDo(context.Background(), entity.ToDo{
		Id:          primitive.NewObjectID(),
		Title:       "testTitle",
		Description: "testDesc",
		Status:      entity.TODO})
	if err != nil {
		t.Errorf("add fail: %v", err)
	}
	t.Log("test success")
}

func TestGetAllToDo(t *testing.T) {
	t.Log("***********GET******************")
	list, err = reposImp.GetAllToDo(context.Background())
	if err != nil {
		t.Errorf("can't get all: %v", err)
	}
	if len(list) == 0 {
		t.Errorf("test failed")
	}
	t.Logf("all todo %v", list)
}
func TestUpdateToDo(t *testing.T) {
	t.Log("***********Update******************")
	status, err := reposImp.UpdateToDo(context.Background(), "000000000000000700000000", entity.DONE)
	if err != nil {
		t.Errorf("can't update: %v", err)
	}
	if !status {
		t.Error("item not found")
	}
}
func TestDeleteToDo(t *testing.T) {
	t.Log("***********Delete******************")

	status, err := reposImp.DeleteToDo(context.Background(), list[0].Id.Hex())
	if err != nil {
		t.Errorf("can't delete: %v", err)
	}
	if !status {
		t.Error("item not found")
	}
	t.Log("test success")
}
