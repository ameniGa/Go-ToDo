package http

import (
	"context"
	"github.com/3almadmoon/ameni-assignment/config"
	pb "github.com/3almadmoon/ameni-assignment/protobuf"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)
type HttpRunner struct {
	*config.Config
}

func NewHttpRunner(conf *config.Config) HttpRunner{
	return HttpRunner{conf}
}

//serveSwagger : handler function
func serveSwagger(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "www/swagger.json")
	log.Println("serve swagger")
}

func (HTTP HttpRunner) Start() error{
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	conn, err := grpc.Dial(HTTP.Server.Grpc.Host, grpc.WithInsecure())
	defer func() {
		if err := conn.Close(); err != nil{
			log.Printf("cannot close client connection: %s",err)
		}
	}()
	if err != nil {
		log.Fatalf("fail to dial : %v", err)
		return err
	}

	rmux := runtime.NewServeMux()
	client := pb.NewTodoListServiceClient(conn)
	err = pb.RegisterTodoListServiceHandlerClient(ctx, rmux, client)
	if err != nil {
		log.Fatalf("can't register %v", err)
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", rmux)
	mux.HandleFunc("/swagger.json", serveSwagger)
	fs := http.FileServer(http.Dir("www/swagger-ui"))
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", fs))

	log.Println("REST server ready...")
	err = http.ListenAndServe(HTTP.Server.Http.Host, mux)
	if err != nil {
		log.Fatalf("can't serve %v ", err)
		return err
	}
	return nil
}