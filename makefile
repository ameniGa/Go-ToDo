install:
	go get
generate:
	protoc \
  -I$(GOPATH)/src \
  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:api/proto \
  --swagger_out=logtostderr=true:api/proto \
  --grpc-gateway_out=logtostderr=true:api/proto \
  --proto_path api/proto service.proto
	
	cp api/swagger/service.swagger.json  www/swagger.json

build:
	go build  cmd/grpcserver/grpcServer.go
	go build  cmd/grpcclient/grpcClient.go
	go build  cmd/httpserver/httpServer.go


rungrpc:
	go run  cmd/grpcserver/grpcServer.go
runhttp:	
	go run  cmd/httpserver/httpServer.go  


runclient:  
	go run  cmd/grpcclient/grpcClient.go


test:
	go test api/reposImp

