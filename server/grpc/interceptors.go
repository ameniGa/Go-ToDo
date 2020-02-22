package grpc

import (
	"context"
	"github.com/3almadmoon/ameni-assignment/helpers"
	pb "github.com/3almadmoon/ameni-assignment/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

// UnaryRequestValidator is a server side interceptor , it validates the requests
func UnaryRequestValidator(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	validCtx := ctx
	err := helpers.CheckTimeout(ctx)
	if err != nil {
		validCtx,_ = context.WithTimeout(context.Background(),5*time.Second)
	}
	if req == nil {
		return nil, status.Error(codes.InvalidArgument,"request cannot be null")
	}
	switch r := req.(type) {
	case *pb.ToDoItem:
		if helpers.IsEmpty(r.Title) {
			return nil, status.Error(codes.InvalidArgument,"title should not be empty")
		}
	case *pb.DeleteToDoItem:
		if helpers.IsEmpty(r.Hash) {
			return nil, status.Error(codes.InvalidArgument,"hash should not be empty")
		}
	case *pb.UpdateToDoItem:
		if helpers.IsEmpty(r.Hash) {
			return nil, status.Error(codes.InvalidArgument,"hash should not be empty")
		}
	default:
		return nil, status.Error(codes.InvalidArgument,"request type unknown")
	}
	return handler(validCtx,req)
}
