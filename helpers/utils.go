package helpers

import (
	"context"
	"time"

	pb "github.com/3almadmoon/ameni-assignment/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const MAX_TIMEOUT = 30 * time.Second

//checkTimeout check if context of rpc contains timeout
//returns error
func CheckTimeout(context context.Context) error {
	deadline, ok := context.Deadline()
	if !ok {
		return status.Error(codes.InvalidArgument, "you must specify a deadline")
	}
	timeout := deadline.Sub(time.Now())
	if  timeout > MAX_TIMEOUT {
		return status.Error(codes.InvalidArgument, "deadline must be  < 30 seconds")
	}
	return nil
}

//throwError handle failed response
//returns genericResponse message and error
func ThrowError(err error, msg string) (*pb.GenericResponse, error) {
	return &pb.GenericResponse{Status: "fail", Message: msg}, err
}

func IsEmpty(field string) bool {
	if field == "" {
		return true
	}
	return false
}
