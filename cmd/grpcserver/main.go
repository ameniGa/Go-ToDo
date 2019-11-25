package main

import (
	"log"

	"net"

	"google.golang.org/grpc"

	pb "github.com/ameniGa/TODO/api/proto"

	service "github.com/ameniGa/TODO/api/serviceimp"

	"sync"
)

var wg sync.WaitGroup

// startGRPC
// start a grpc server
// register service to server
func startGRPC() {
	lis, err := net.Listen("tcp", "localhost:5566")
	if err != nil {
		log.Fatalf("Failed to listen : %v ", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterTodoListServiceServer(grpcServer, &service.TodoListServiceServer{})
	if error := grpcServer.Serve(lis); error != nil {
		log.Fatalf("Failed to serve %v", error)
	}
	log.Println("server ready to listen ....")
	grpcServer.Serve(lis)
}

func main() {

	go startGRPC()

	wg.Add(1)
	wg.Wait()
}
