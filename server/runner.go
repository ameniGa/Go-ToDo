package server

import (
	"github.com/3almadmoon/ameni-assignment/config"
	"github.com/3almadmoon/ameni-assignment/server/grpc"
	"github.com/3almadmoon/ameni-assignment/server/http"
	"log"
)

type Runner interface {
	Start() error
}
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