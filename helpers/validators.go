package helpers

import (
	"context"
	"time"

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

// IsEmpty checks if a string is empty
func IsEmpty(field string) bool {
	return field == ""
}
