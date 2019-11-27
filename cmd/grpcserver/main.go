package main

import (
	"log"

	"net"

	"google.golang.org/grpc"

	"github.com/3almadmoon/ameni-assignment/api/db"

	ri "github.com/3almadmoon/ameni-assignment/api/reposImp"

	utils "github.com/3almadmoon/ameni-assignment/api"

	pb "github.com/3almadmoon/ameni-assignment/api/proto"

	service "github.com/3almadmoon/ameni-assignment/api/serviceimp"

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
	lis, err := net.Listen("tcp", utils.GRPC_BASE_URL)
	if err != nil {
		log.Fatalf("Failed to listen : %v ", err)
	}
	grpcServer := grpc.NewServer()
	myReposImp := ri.NewToDoRepos(collection)
	pb.RegisterTodoListServiceServer(grpcServer, &service.TodoListServiceServer{myReposImp})
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
