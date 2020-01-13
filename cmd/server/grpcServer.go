package main

import (
	srv "github.com/3almadmoon/ameni-assignment/server"
	"log"
	"github.com/3almadmoon/ameni-assignment/config"
	"sync"
)

var wg sync.WaitGroup

// startGRPC
// start a grpc server
// register service to server
func startGRPC() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Panicf("cannot parse config file: %v", err)
	}
	srv := srv.CreateRunner(conf, "grpc")
    err = srv.Start()
    if err != nil {
    	log.Panicf("cannot start server: %v",err)
	}
}

func main() {
	go startGRPC()
	wg.Add(1)
	wg.Wait()
}
