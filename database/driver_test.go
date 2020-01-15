package database

import (
	"context"
	"github.com/3almadmoon/ameni-assignment/config"
	entity "github.com/3almadmoon/ameni-assignment/entities"
	"testing"
	"time"
)

var (
	mContext context.Context
)

var (
	db  *MongoDBhandler
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
			"tasks",
			"todo-test",
		},
	}
	db, err = NewMongoDBhandler(&conf)
}

func TestAddToDo(t *testing.T) {
	ctx := context.Background()
	var canc context.CancelFunc
	for _, testCase := range TtToDo {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.IsValidCtx {
				ctx, canc = context.WithTimeout(context.Background(), 2*time.Second)
				defer canc()
			}
			item := entity.ToDo{
				Hash:        testCase.Hash,
				Title:       testCase.Title,
				Description: testCase.Description,
				Status:      testCase.Status,
			}
			err := db.AddToDo(ctx, item)
			if testCase.HasErrorOnCreate {
				if err == nil {
					t.Error("expected error got nothing")
				}
			} else {
				if err != nil {
					t.Errorf("expected success got %v", err)
				}
			}
		})
	}
}

func TestGetAllToDo(t *testing.T) {
	t.Run("get all items successfully", func(t *testing.T) {
		ctx, canc := context.WithTimeout(context.Background(), 2*time.Second)
		defer canc()

		list, err := db.GetAllToDo(ctx)
		if err != nil {
			t.Errorf("can't get all: %v", err)
		}
		if len(list) == 0 {
			t.Errorf("expected items got empty list")
		}
	})
}
func TestUpdateToDo(t *testing.T) {
	ctx := context.Background()
	var canc context.CancelFunc
	for _, testCase := range TtToDo {
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
	var canc context.CancelFunc
	for _, testCase := range TtToDo[1:] {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.IsValidCtx {
				ctx, canc = context.WithTimeout(context.Background(), 2*time.Second)
				defer canc()
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
