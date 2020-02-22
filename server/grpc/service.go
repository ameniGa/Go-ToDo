package grpc

import (
	"context"
	entity "github.com/3almadmoon/ameni-assignment/entities"
	"github.com/3almadmoon/ameni-assignment/helpers"
	pb "github.com/3almadmoon/ameni-assignment/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

//AddToDo add ToDoItem message and returns a genericResponse message
//which indicates success or fail and detailed message and error
func (svc *GrpcRunner) AddToDo(ctx context.Context, item *pb.ToDoItem) (*pb.GenericResponse, error) {
	toDo := entity.ToDo{
		Hash:        uuid.New().String(),
		Title:       item.GetTitle(),
		Description: item.GetDescription(),
		Status:      entity.EStatus(item.GetStatus()),
	}

	addChan := make(chan error,1)
    go func(ch chan<- error) {
    	defer close(ch)
		err := svc.DB.AddToDo(ctx, toDo)
		ch <- err
	}(addChan)

	err := <- addChan
	if err != nil {
		return helpers.ThrowError("item not added",err, codes.NotFound)
	}
	resp :=helpers.CreateResponse("success","item added")
	return resp, nil
}

//DeleteToDo delete to do item by hash
//returns Generic response
//which indicates success or fail and detailed message and error
func (svc *GrpcRunner) DeleteToDo(context context.Context, item *pb.DeleteToDoItem) (*pb.GenericResponse, error) {

	deleteChan := make(chan entity.StatusWithError ,1)
	go func(ch chan<- entity.StatusWithError) {
		defer close(ch)
		ok, err := svc.DB.DeleteToDo(context, item.GetHash())
		ch <- entity.StatusWithError{
			Status: ok,
			Err:    err,
		}

	}(deleteChan)
    res := <- deleteChan
	if res.Err != nil {
		return helpers.ThrowError("error occured", res.Err, codes.Internal)
	}
	if res.Status == false {
		return helpers.ThrowError("item not found", nil, codes.NotFound)
	}
	resp := helpers.CreateResponse("success", "item deleted")
	return resp, nil
}

//UpdateToDo update status of to do item by hash
//returns Generic response
//which indicates success or fail and detailed message and error
func (svc *GrpcRunner) UpdateToDo(context context.Context, item *pb.UpdateToDoItem) (*pb.GenericResponse, error) {
	updateChan := make(chan entity.StatusWithError ,1)
	go func(ch chan<- entity.StatusWithError) {
		defer close(ch)
		ok, err := svc.DB.UpdateToDo(context, item.GetHash(), entity.EStatus(item.GetStatus()))
		ch <- entity.StatusWithError{
			Status: ok,
			Err:    err,
		}
	}(updateChan)
	res := <- updateChan
	if res.Err != nil {
		return helpers.ThrowError("error occured",res.Err, codes.Internal)
	}
	if res.Status == false {
		return helpers.ThrowError("item not found", nil, codes.NotFound)
	}
	resp := helpers.CreateResponse("success","item updated")
	return resp, nil
}

//GetAllToDo stream all to do items
//returns error
func (svc *GrpcRunner) GetAllToDo(_ *empty.Empty, stream pb.TodoListService_GetAllToDoServer) error {
	var streamErr error
	dbCtx,cancel := context.WithTimeout(stream.Context(),5*time.Second)
	defer cancel()

	getChan := make(chan entity.ToDoWithError)
	go svc.DB.GetAllToDo(dbCtx, getChan)

	var errorsList []string
	for output := range getChan {
		if output.Err != nil {
			return output.Err
		}
		toDoItem := pb.ToDoItem{
			Title:       output.ToDo.Title,
			Description: output.ToDo.Description,
			Status:      pb.Status(output.ToDo.Status),
		}
		elem := pb.GetToDoItem{
			Hash:                 output.ToDo.Hash,
			ToDoItem:             &toDoItem,
		}
		streamErr = stream.Send(&elem)
		if streamErr != nil {
			errorsList = append(errorsList, streamErr.Error())
		}
	}
	if len(errorsList) > 0 {
		return status.Error(codes.Code(code.Code_INTERNAL), "failed to Get All todos")
	}
	return nil
}
