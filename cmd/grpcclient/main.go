package main

import (
	"context"
	pb "github.com/ameniGa/TODO/api/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

// runClient
// create grpc client
// make a remote call
func runClient() {
	ctx, cancelTimeoutFunc := context.WithTimeout(context.Background(), 3*time.Second)
	conn, err := grpc.DialContext(ctx, "localhost:5566", grpc.WithInsecure(), grpc.WithBlock())
	cancelTimeoutFunc()
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	log.Printf("Dialled OK...")
	defer conn.Close()

	client := pb.NewTodoListServiceClient(conn)
	log.Printf("created client")

	item, er := client.AddToDo(context.Background(), &pb.ToDoItem{Title: "learn golang", Tag: "bug"})
	if er != nil {
		log.Fatalf("failed to add an item : %v", er)
	}
	log.Printf("Got a new item %v", item.Msg)
}

func main() {
	runClient()
}
