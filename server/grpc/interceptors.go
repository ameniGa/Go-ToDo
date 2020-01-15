package grpc

import (
	"context"
	"errors"
	"github.com/3almadmoon/ameni-assignment/helpers"
	pb "github.com/3almadmoon/ameni-assignment/protobuf"
	"google.golang.org/grpc"
	"time"
)

func UnaryRequestValidator(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	validCtx := ctx
	err := helpers.CheckTimeout(ctx)
	if err != nil {
		validCtx,_ = context.WithTimeout(context.Background(),5*time.Second)
	}
	if req == nil {
		return nil,errors.New("request cannot be null")
	}
	switch r := req.(type) {
	case *pb.ToDoItem:
		if helpers.IsEmpty(r.Title) {
			return nil, errors.New("title should not be empty!")
		}
	case *pb.DeleteToDoItem:
		if helpers.IsEmpty(r.Hash) {
			return nil, errors.New("hash should not be empty!")
		}
	case *pb.UpdateToDoItem:
		if helpers.IsEmpty(r.Hash) {
			return nil, errors.New("hash should not be empty!")
		}
	default:
		return nil, errors.New("request type unknown")
	}
	return handler(validCtx,req)
}
