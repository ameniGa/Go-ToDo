package main

import (
	"context"
	"net/http"

	"log"

	"google.golang.org/grpc"

	"sync"

	pb "github.com/3almadmoon/ameni-assignment/api/proto"
	config "github.com/3almadmoon/ameni-assignment/config"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
)

var wg sync.WaitGroup

type Adapter func(http.Handler) http.Handler

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

	conn, err := grpc.Dial(viper.GetString("grpcserver.host"), grpc.WithInsecure())
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
	err = http.ListenAndServe(viper.GetString("httpserver.host"), mux)
	if err != nil {
		log.Fatalf("can't serve %v ", err)
	}

}
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

// WithHeader is an Adapter that sets an HTTP handler.
func WithHeader(key, value string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(key, value)
			h.ServeHTTP(w, r)
		})
	}
}

func main() {
	if err := config.SetViper(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	go startHTTP()
	wg.Add(1)
	wg.Wait()

}
