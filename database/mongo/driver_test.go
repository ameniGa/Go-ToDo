package mongo

import (
	"context"
	"github.com/3almadmoon/ameni-assignment/config"
	entity "github.com/3almadmoon/ameni-assignment/entities"
	td "github.com/3almadmoon/ameni-assignment/testData"
	"testing"
	"time"
)

var (
	db  *MongoHandler
	err error
)

func init() {
	conf := config.Config{
		Database: struct {
			Type       string
			Uri        string
			Name       string
			Collection string
		}{
			"mongo",
			"mongodb://localhost:27017",
			"tasks-test",
			"todo-test",
		},
	}
	db, err = NewMongoDBhandler(&conf)
}

func TestAddToDo(t *testing.T) {
	ctx := context.Background()
	var cancel context.CancelFunc
	for _, testCase := range td.TTtoDo {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.IsValidCtx {
				ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
				defer cancel()
			}
			item := entity.ToDo{
				Hash:        testCase.Hash,
				Title:       testCase.Title,
				Description: testCase.Description,
				Status:      testCase.Status,
			}
			err := db.AddToDo(ctx, item)
			if err == nil && testCase.HasErrorOnCreate {
				t.Error("expected error got nothing")
			}
			if err != nil && !testCase.HasErrorOnCreate {
				t.Errorf("expected success got %v", err)
			}
		})
	}
}

func TestGetAllToDo(t *testing.T) {
	t.Run("get all items successfully", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		getChan := make(chan entity.ToDoWithError,20)
		db.GetAllToDo(ctx,getChan)

		for output := range getChan {
			if output.Err != nil {
				t.Errorf("expected success got: %s", err)
				break
			}
			if output.ToDo == nil{
				t.Errorf("expected items got empty list")
				break
			}
           t.Logf("%v",output)
		}
	})
}
func TestUpdateToDo(t *testing.T) {
	ctx := context.Background()
	var canc context.CancelFunc
	for _, testCase := range td.TTtoDo {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.IsValidCtx {
				ctx, canc = context.WithTimeout(context.Background(), 2*time.Second)
				defer canc()
			}
			ok, err := db.UpdateToDo(ctx, testCase.Hash, testCase.Status)
			if testCase.HasErrorOnUpdate {
				if err == nil || ok {
					t.Error("expected error got nothing")
				}
			} else {
				if err != nil || !ok {
					t.Errorf("expected success got %v", err)
				}
			}
		})
	}
}
func TestDeleteToDo(t *testing.T) {
	ctx := context.Background()
	var cancel context.CancelFunc
	for _, testCase := range td.TTtoDo[1:] {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.IsValidCtx {
				ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
				defer cancel()
			}
			status, err := db.DeleteToDo(ctx, testCase.Hash)
			if testCase.HasErrorOnCreate {
				if err == nil || status {
					t.Error("expected error got nothing")
				}
			} else {
				if err != nil || !status {
					t.Errorf("expected success got %v", err)
				}
			}
		})
	}
}
