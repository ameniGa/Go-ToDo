package grpc

import (
	"context"
	entity "github.com/3almadmoon/ameni-assignment/entities"
	"github.com/3almadmoon/ameni-assignment/helpers"
	pb "github.com/3almadmoon/ameni-assignment/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"time"
)

//AddToDo add ToDoItem message and returns a genericResponse message
//which indicates success or fail and detailed message and error
func (svc *GrpcRunner) AddToDo(ctx context.Context, item *pb.ToDoItem) (*pb.GenericResponse, error) {
	//TODO call db in go routine
	toDo := entity.ToDo{
		Hash:        uuid.New().String(),
		Title:       item.GetTitle(),
		Description: item.GetDescription(),
		Status:      entity.EStatus(item.GetStatus()),
	}
	err := svc.DB.AddToDo(ctx, toDo)
	if err != nil {
		return helpers.ThrowError(err, "item not added")
	}
	resp := pb.GenericResponse{
		Status:  "success",
		Message: "item added",
	}
	return &resp, nil
}

//DeleteToDo delete to do item by hash
//returns Generic response
//which indicates success or fail and detailed message and error
func (svc *GrpcRunner) DeleteToDo(context context.Context, item *pb.DeleteToDoItem) (*pb.GenericResponse, error) {

	status, err := svc.DB.DeleteToDo(context, item.GetHash())
	if err != nil {
		return helpers.ThrowError(err, "item not added")
	}
	if status == false {
		return helpers.ThrowError(nil, "item not found")
	}
	resp := pb.GenericResponse{
		Status:  "success",
		Message: "item deleted",
	}
	return &resp, nil

}

//UpdateToDo update status of to do item by hash
//returns Generic response
//which indicates success or fail and detailed message and error
func (svc *GrpcRunner) UpdateToDo(context context.Context, item *pb.UpdateToDoItem) (*pb.GenericResponse, error) {
	status, err := svc.DB.UpdateToDo(context, item.GetHash(), entity.EStatus(item.GetStatus()))
	if err != nil {
		return helpers.ThrowError(err, "item not updated")
	}
	if status == false {
		return helpers.ThrowError(nil, "item not found")
	}
	resp := pb.GenericResponse{
		Status:  "success",
		Message: "item updated",
	}
	return &resp, nil
}

//GetAllToDo stream all to do items
//returns error
func (svc *GrpcRunner) GetAllToDo(item *empty.Empty, stream pb.TodoListService_GetAllToDoServer) error {
	var streamErr error
	dbCtx,canc := context.WithTimeout(stream.Context(),5*time.Second)
	defer canc()
	res, err := svc.DB.GetAllToDo(dbCtx)
	if err != nil {
		return err
	}
	for _, todo := range res {
		toDoItem := pb.ToDoItem{
			Title:       todo.Title,
			Description: todo.Description,
			Status:      pb.Status(todo.Status),
		}
		elem := pb.GetToDoItem{
			Hash:                 todo.Hash,
			ToDoItem:             &toDoItem,
		}
		streamErr = stream.Send(&elem)
		if streamErr != nil {
			return streamErr
		}
	}
	return nil
}
