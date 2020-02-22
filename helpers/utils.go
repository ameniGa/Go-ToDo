package helpers

import (
	pb "github.com/3almadmoon/ameni-assignment/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//throwError handle failed response
//returns genericResponse message and error
func ThrowError(msg string, err error, code codes.Code) (*pb.GenericResponse, error) {
	resp := &pb.GenericResponse{
		Status:  "fail",
		Message: msg,
	}
	return resp, status.Error(code, err.Error())
}

// CreateResponse returns genericResponse object
func CreateResponse(status, msg string) *pb.GenericResponse {
	res := pb.GenericResponse{
		Status:  status,
		Message: msg,
	}
	return &res
}
