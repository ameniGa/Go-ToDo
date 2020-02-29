# TODO app

- Is a simple service written in **Golang**, provides APIs via **gRPC** and **REST**.
- I used **Mongo** as database, **unary interceptor** to validate requests and **swagger** to test REST APIs.
- This is my first project in Go :D

### Prerequisites
1. [install Go](https://golang.org/doc/install)
2. [install MongoDB](https://docs.mongodb.com/manual/installation/)

### install app dependencies
    make install

### generate protobuf, swagger and grpc gateway
    make generate

### build
    make build
    
### run grpc server 
    make rungrpc
   
### run test
    make test
        
### run client CLI
 to use client cli you have to build the client and run grpc server.
 for more information about how to use the client cli, try help command.

    ./grpcClient help
     
### run http server
    make runhttp

### visit [swagger](http://localhost:8080/swagger-ui) to make REST requests


    



