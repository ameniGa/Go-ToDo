package server

import (
	"github.com/3almadmoon/ameni-assignment/config"
	"github.com/3almadmoon/ameni-assignment/server/grpc"
	"github.com/3almadmoon/ameni-assignment/server/http"
	"log"
)

// Runner holds the function to implement by each server
type Runner interface {
	Start() error
}

// CreateRunner creates a runner from the given config
func CreateRunner(conf *config.Config,serverType string) Runner {
	var runner Runner
	switch serverType {
	case "grpc":
		runner = grpc.NewGrpcRunner(conf)
	case "http":
		runner = http.NewHttpRunner(conf)
	default:
		log.Panicf("%v not supported as server type ",serverType)
	}
	return runner
}