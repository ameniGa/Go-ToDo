package grpc

import (
	"context"
	pb "github.com/3almadmoon/ameni-assignment/protobuf"
	td "github.com/3almadmoon/ameni-assignment/testData"
	"testing"
	"time"
)


func TestUnaryRequestValidator(t *testing.T) {
	ctx := context.Background()
	var canc context.CancelFunc
	for _, testCase := range td.TTReqValidation {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.IsValidCtx {
				ctx, canc = context.WithTimeout(context.Background(), 2*time.Second)
				defer canc()
			}
			var req interface{}
			switch testCase.ReqType {
			case "add":
				req = &pb.ToDoItem{
					Title:       testCase.Title,
					Description: "",
					Status:      0,
				}
			case "update":
				req = &pb.UpdateToDoItem{
					Hash: testCase.Hash,
				}
			case "delete":
				req = &pb.DeleteToDoItem{
					Hash: testCase.Hash,
				}
			case "unkown":
				req = pb.ToDoItem{}
			}

			_, err := UnaryRequestValidator(ctx, req, nil, handler)
			if testCase.HasError {
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

func handler(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}
