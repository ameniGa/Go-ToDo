package main

import (
	"log"

	"net"

	"google.golang.org/grpc"

	"github.com/3almadmoon/ameni-assignment/api/db"
	"github.com/spf13/viper"

	ri "github.com/3almadmoon/ameni-assignment/api/todoLogic"

	pb "github.com/3almadmoon/ameni-assignment/api/proto"

	service "github.com/3almadmoon/ameni-assignment/api/service"

	config "github.com/3almadmoon/ameni-assignment/config"

	"sync"
)

var wg sync.WaitGroup

// startGRPC
// start a grpc server
// register service to server
func startGRPC() {
	// get file name and line number in case of server crash
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//connection to DB
	collection, err := db.Connect()
	if err != nil {
		log.Fatalf("can't connect to mongoDB %v", err)
	}
	lis, err := net.Listen("tcp", viper.GetString("grpcserver.host"))
	if err != nil {
		log.Fatalf("Failed to listen : %v ", err)
	}
	grpcServer := grpc.NewServer()
	toDoImp := ri.TodoImp{ToDoCollection: collection}
	pb.RegisterTodoListServiceServer(grpcServer, &service.TodoListService{toDoImp})
	if error := grpcServer.Serve(lis); error != nil {
		log.Fatalf("Failed to serve %v", error)
	}
	log.Println("server ready to listen ....")
	grpcServer.Serve(lis)
}

func main() {
	if err := config.SetViper(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	go startGRPC()

	wg.Add(1)
	wg.Wait()
}
