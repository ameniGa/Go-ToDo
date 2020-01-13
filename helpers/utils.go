package helpers

import pb "github.com/3almadmoon/ameni-assignment/protobuf"

//throwError handle failed response
//returns genericResponse message and error
func ThrowError(err error, msg string) (*pb.GenericResponse, error) {
	resp := &pb.GenericResponse{
		Status: "fail",
		Message: msg,
	}
	return resp, err
}
