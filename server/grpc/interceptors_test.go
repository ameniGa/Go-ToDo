package grpc

import (
	"context"
	pb "github.com/3almadmoon/ameni-assignment/protobuf"
	"github.com/google/uuid"
	"testing"
	"time"
)

var tt= []struct{
	name string
	hash string
	title string
	isAddItem bool
	hasError bool
	isValidCtx bool
}{
	{"invalid context","","golang",true,false,false},
	{"valid Add request","","golang",true,false,true},
    {"empty title in Add request","","",true,true,true},
	{"empty hash in Delete request","","",false,true,true},
	{"valid delete request",uuid.New().String(),"",false,false,true},
}

func TestUnaryRequestValidator(t *testing.T) {
	ctx := context.Background()
	var canc context.CancelFunc
	for _,testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.isValidCtx {
				ctx, canc = context.WithTimeout(context.Background(), 2*time.Second)
				defer canc()
			}

			var req interface{}
			if testCase.isAddItem {
				req = &pb.ToDoItem{
					Title:       testCase.title,
					Description: "",
					Status:      0,
				}
			} else {
				req = &pb.DeleteToDoItem{
					Hash: testCase.hash,
				}
			}

			_, err := UnaryRequestValidator(ctx, req, nil, handler)
			if testCase.hasError {
				if err == nil{
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

func handler(ctx context.Context, req interface{}) (interface{}, error) {
	return req, nil
}