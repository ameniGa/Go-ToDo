package main

import (
	"context"
	"net/http"

	"log"

	"google.golang.org/grpc"

	utils "github.com/3almadmoon/ameni-assignment/api"
	pb "github.com/3almadmoon/ameni-assignment/api/proto"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

var wg sync.WaitGroup

//serveSwagger : handler function
func serveSwagger(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "www/swagger.json")
	log.Println("serve swagger")
}

//startHTTP
//connect to grpc server
//register for grpc gateway
// serve the swagger-ui and swagger
func startHTTP() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn, err := grpc.Dial(utils.GRPC_BASE_URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial : v%", err)
	}
	defer conn.Close()

	rmux := runtime.NewServeMux()
	client := pb.NewTodoListServiceClient(conn)
	err = pb.RegisterTodoListServiceHandlerClient(ctx, rmux, client)
	if err != nil {
		log.Fatalf("can't register %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", rmux)

	mux.HandleFunc("/swagger.json", serveSwagger)
	fs := http.FileServer(http.Dir("www/swagger-ui"))
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", fs))

	log.Println("REST server ready...")
	err = http.ListenAndServe(utils.HTTP_BASE_URL, mux)
	if err != nil {
		log.Fatalf("can't serve %v ", err)
	}

}

func main() {

	go startHTTP()
	wg.Add(1)
	wg.Wait()

}
