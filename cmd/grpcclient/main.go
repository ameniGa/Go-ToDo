package main

import (
	"context"
	utils "github.com/3almadmoon/ameni-assignment/api"
	pb "github.com/3almadmoon/ameni-assignment/api/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

// runClient
// create grpc client
// make a remote call
func runClient() {
	ctx, cancelTimeoutFunc := context.WithTimeout(context.Background(), 3*time.Second)
	conn, err := grpc.DialContext(ctx, utils.GRPC_BASE_URL, grpc.WithInsecure(), grpc.WithBlock())
	cancelTimeoutFunc()
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	log.Printf("Dialled OK...")
	defer conn.Close()

	client := pb.NewTodoListServiceClient(conn)
	log.Printf("created client")
	log.Println("**********ADD**************")
	res1, er1 := client.AddToDo(context.Background(), &pb.ToDoItem{Title: "learn golang", Description: "bug", Status: pb.Status_TODO})
	log.Printf("|RES|:\n %v, \n |ERROR|: \n %v", res1, er1)
	log.Println("**********GETALL**************")
	res2, er2 := client.GetAllToDo(context.Background(), &pb.EmptyRequest{})
	log.Printf("|RES|:\n %v, \n |ERROR|: \n %v", res2, er2)
	log.Println("**********UPDATE**************")
	res3, er3 := client.UpdateToDo(context.Background(), &pb.UpdateToDoItem{Id: res2.GetToDoItems()[0].GetId(), Status: pb.Status_INPROGRESS})
	log.Printf("|RES|:\n %v, \n |ERROR|: \n %v", res3, er3)
	log.Println("**********DELETE**************")
	res4, er4 := client.DeleteToDo(context.Background(), &pb.DeleteToDoItem{Id: res2.GetToDoItems()[0].GetId()})
	log.Printf("|RES|:\n %v, \n |ERROR|: \n %v", res4, er4)

}

func main() {
	runClient()
}
