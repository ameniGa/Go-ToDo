package serviceimp

import (
	"context"
	"time"

	entity "github.com/3almadmoon/ameni-assignment/api/entities"
	pb "github.com/3almadmoon/ameni-assignment/api/proto"
	ri "github.com/3almadmoon/ameni-assignment/api/reposImp"
	"github.com/golang/protobuf/ptypes/empty"
)

//TodoListServiceServer this is the implementation of server API TodoListServiceServer for TodoListService service
type TodoListServiceServer struct {
	ReposImp ri.ReposImp
}

const (
	//MIN_TIMEOUT minimum timeout acceptes
	MIN_TIMEOUT = 5 * time.Second
	//MAX_TIMEOUT maximum timeout accepted
	MAX_TIMEOUT = 30 * time.Second
)

//AddToDo add ToDoItem message and returns a genericResponse message
//which indicates success or fail and detailed message and error
func (s *TodoListServiceServer) AddToDo(ctx context.Context, item *pb.ToDoItem) (*pb.GenericResponse, error) {
	//check rpc timeout
	checkErr := checkTimeout(ctx)
	if checkErr != nil {
		return throwError(checkErr, "item not added")
	}
	if item.GetTitle() != "" {
		toDo := entity.ToDo{
			Hash:        getToDoHash(item.GetTitle()),
			Title:       item.GetTitle(),
			Description: item.GetDescription(),
			Status:      entity.EStatus(item.GetStatus()),
		}
		err := s.ReposImp.AddToDo(ctx, toDo) // this is goroutine
		if err != nil {
			return throwError(err, "item not added")
		}
		return &pb.GenericResponse{Status: "success", Message: "item added"}, nil
	}
	return &pb.GenericResponse{Status: "fail", Message: "title cannot be empty"}, nil
}

//DeleteToDo delete todo item by hash
//returns Generic response
//which indicates success or fail and detailed message and error
func (s *TodoListServiceServer) DeleteToDo(context context.Context, item *pb.DeleteToDoItem) (*pb.GenericResponse, error) {
	checkErr := checkTimeout(context)
	if checkErr != nil {
		return throwError(checkErr, "item not deleted")
	}
	if (item != &pb.DeleteToDoItem{}) {
		status, err := s.ReposImp.DeleteToDo(context, item.GetHash())
		if err != nil {
			return throwError(err, "item not added")
		}
		if status == false {
			return throwError(nil, "item not found")
		}
		return &pb.GenericResponse{Status: "success", Message: "item deleted"}, nil
	}
	return &pb.GenericResponse{Status: "fail", Message: "hash cannot be empty"}, nil
}

//UpdateToDo update status of todo item by hash
//returns Generic response
//which indicates success or fail and detailed message and error
func (s *TodoListServiceServer) UpdateToDo(context context.Context, item *pb.UpdateToDoItem) (*pb.GenericResponse, error) {
	checkErr := checkTimeout(context)
	if checkErr != nil {
		return throwError(checkErr, "item not deleted")
	}
	if (item != &pb.UpdateToDoItem{}) {
		status, err := s.ReposImp.UpdateToDo(context, item.GetHash(), entity.EStatus(item.GetStatus()))

		if err != nil {
			return throwError(err, "item not updated")
		}
		if status == false {
			return throwError(nil, "item not found")
		}
		return &pb.GenericResponse{Status: "success", Message: "item updated"}, nil
	}
	return &pb.GenericResponse{Status: "fail", Message: "body cannot be empty"}, nil
}

//GetAllToDo stream all todo items
//returns error
func (s *TodoListServiceServer) GetAllToDo(item *empty.Empty, stream pb.TodoListService_GetAllToDoServer) error {
	checkErr := checkTimeout(stream.Context())
	if checkErr != nil {
		return checkErr
	}
	var streamErr error
	res, err := s.ReposImp.GetAllToDo(context.Background())
	if err != nil {
		return err
	}
	for _, todo := range res {

		elem := pb.AllToDoItems{
			ToDoItems: &pb.GetToDoItem{
				Hash: todo.Hash,
				ToDoItem: &pb.ToDoItem{
					Title:       todo.Title,
					Description: todo.Description,
					Status:      pb.Status(todo.Status),
				},
			}, GenericResponse: &pb.GenericResponse{Status: "success", Message: "todo list"},
		}
		streamErr = stream.Send(&elem)
		if streamErr != nil {
			return streamErr
		}
	}
	return nil
}
