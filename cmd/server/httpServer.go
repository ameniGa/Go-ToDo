package main

import (
	"github.com/3almadmoon/ameni-assignment/config"
	"github.com/3almadmoon/ameni-assignment/server"
	"log"
	"sync"
)

var wgp sync.WaitGroup

func startHTTP() {
	conf,err := config.GetConfig()
	if err != nil {
		log.Panicf("cannot parse config file: %v", err)
	}
	srv := server.CreateRunner(conf,"http")
	err = srv.Start()
	if err != nil {
		log.Panicf("cannot start server: %v",err)
	}
}

func main() {
	go startHTTP()
	wgp.Add(1)
	wgp.Wait()

}
