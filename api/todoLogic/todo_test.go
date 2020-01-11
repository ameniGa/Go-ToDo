package todoLogic

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"testing"

	"github.com/3almadmoon/ameni-assignment/api/db"
	entity "github.com/3almadmoon/ameni-assignment/api/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection
	todoImp    TodoImp
	mContext   context.Context
)

func init() {
	collection, _ = db.Connect()
	todoImp = TodoImp{ToDoCollection: collection}
}
func GetToDoHash(uniqueField string) string {
	hashWriter := md5.New()
	io.WriteString(hashWriter, uniqueField)
	hashB := hashWriter.Sum(nil)
	return fmt.Sprintf("%x", string(hashB[:]))
}

var list []*entity.ToDo
var err error

func TestAddToDo(t *testing.T) {
	t.Log("***********ADD******************")
	err := todoImp.AddToDo(context.Background(), entity.ToDo{
		Hash:        GetToDoHash("testTitlezzzz"),
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
	list, err = todoImp.GetAllToDo(context.Background())
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
	status, err := todoImp.UpdateToDo(context.Background(), "000000000000000700000000", entity.DONE)
	if err != nil {
		t.Errorf("can't update: %v", err)
	}
	if !status {
		t.Error("item not found")
	}
}
func TestDeleteToDo(t *testing.T) {
	t.Log("***********Delete******************")

	status, err := todoImp.DeleteToDo(context.Background(), list[0].Hash)
	if err != nil {
		t.Errorf("can't delete: %v", err)
	}
	if !status {
		t.Error("item not found")
	}
	t.Log("test success")
}
