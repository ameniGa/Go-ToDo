package serviceimp

import (
	"context"
	"fmt"

	entity "github.com/3almadmoon/ameni-assignment/api/entities"
	pb "github.com/3almadmoon/ameni-assignment/api/proto"
	ri "github.com/3almadmoon/ameni-assignment/api/reposImp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoListServiceServer struct {
	ReposImp ri.ReposImp
}

func (s *TodoListServiceServer) AddToDo(context context.Context, item *pb.ToDoItem) (*pb.GenericResponse, error) {
	toDo := entity.ToDo{
		Id:          primitive.NewObjectID(),
		Title:       item.GetTitle(),
		Description: item.GetDescription(),
		Status:      entity.EStatus(item.GetStatus()),
	}
	err := s.ReposImp.AddToDo(context, toDo)
	if err != nil {
		return &pb.GenericResponse{Status: "fail", Message: "item not added"}, err
	}
	return &pb.GenericResponse{Status: "success", Message: "item added"}, nil

}

func (s *TodoListServiceServer) DeleteToDo(context context.Context, item *pb.DeleteToDoItem) (*pb.GenericResponse, error) {
	status, err := s.ReposImp.DeleteToDo(context, item.GetId())
	if err != nil {
		return &pb.GenericResponse{Status: "fail", Message: "item not deleted"}, err
	}
	if status == false {
		return &pb.GenericResponse{Status: "fail", Message: "item not found"}, nil
	}
	return &pb.GenericResponse{Status: "success", Message: "item deleted"}, nil
}
func (s *TodoListServiceServer) UpdateToDo(context context.Context, item *pb.UpdateToDoItem) (*pb.GenericResponse, error) {
	status, err := s.ReposImp.UpdateToDo(context, item.GetId(), entity.EStatus(item.GetStatus()))
	if err != nil {
		return &pb.GenericResponse{Status: "fail", Message: "item not updated"}, err
	}
	if status == false {
		return &pb.GenericResponse{Status: "fail", Message: "item not found"}, nil
	}
	return &pb.GenericResponse{Status: "success", Message: "item updated"}, nil

}

func (s *TodoListServiceServer) GetAllToDo(context context.Context, item *pb.EmptyRequest) (*pb.AllToDoItems, error) {
	var list []*pb.GetToDoItem
	res, err := s.ReposImp.GetAllToDo(context)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(res); i++ {
		elem := pb.GetToDoItem{
			Id:          res[i].Id.Hex(),
			Title:       res[i].Title,
			Description: res[i].Description,
			Status:      pb.Status(res[i].Status),
		}
		list = append(list, &elem)
	}
	fmt.Printf("list %v", list)
	return &pb.AllToDoItems{ToDoItems: list, GenericResponse: &pb.GenericResponse{Status: "success", Message: "todo list"}}, nil
}
