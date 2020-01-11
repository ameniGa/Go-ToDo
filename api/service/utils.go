package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/3almadmoon/ameni-assignment/api/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//GetToDoHash calculate hash of toDo ite
//take title as input
//returns string
func GetToDoHash(uniqueField string) string {
	hashWriter := md5.New()
	io.WriteString(hashWriter, uniqueField)
	hashB := hashWriter.Sum(nil)
	return fmt.Sprintf("%x", string(hashB[:]))
}

//checkTimeout check if context of rpc contains timeout
//returns error
func checkTimeout(context context.Context) error {
	deadline, ok := context.Deadline()
	if !ok {
		return status.Error(codes.InvalidArgument, "you must specify a deadline")
	}
	timeout := deadline.Sub(time.Now())
	log.Printf("timeouut %v \n", timeout)
	if timeout < MIN_TIMEOUT || timeout > MAX_TIMEOUT {
		return status.Error(codes.InvalidArgument, "deadline must be 5-30 seconds")
	}
	return nil
}

//throwError handle failed response
//returns genericResponse message and error
func throwError(err error, msg string) (*pb.GenericResponse, error) {
	return &pb.GenericResponse{Status: "fail", Message: msg}, err
}
