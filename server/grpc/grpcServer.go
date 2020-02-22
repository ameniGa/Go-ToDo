package grpc

import(
	"github.com/3almadmoon/ameni-assignment/config"
	db"github.com/3almadmoon/ameni-assignment/database"
	pb "github.com/3almadmoon/ameni-assignment/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
)

// GrpcRunner represents the runner of the service
type GrpcRunner struct {
	DB db.DBhandler
	*config.Config
}

// NewGrpcRunner creates a grpc runner
func NewGrpcRunner(conf *config.Config) GrpcRunner{
    dbHandler,err := db.CreateDBhandler(conf)
    if err != nil {
    	log.Panicf("cannot create db handler")
	}
	return GrpcRunner{dbHandler,conf}
}

// Start starts grpc server
func (svc GrpcRunner) Start() error {
	lis, err := net.Listen("tcp", svc.Server.Grpc.Host)
	if err != nil {
		log.Fatalf("Failed to listen : %v ", err)
		return err
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(UnaryRequestValidator))
	pb.RegisterTodoListServiceServer(grpcServer, &svc)
	if error := grpcServer.Serve(lis); error != nil {
		log.Fatalf("Failed to serve %v", error)
		return error
	}
	log.Println("server ready to listen ....")
	return grpcServer.Serve(lis)
}
