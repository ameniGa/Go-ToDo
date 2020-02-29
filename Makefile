GOPATH=$(shell go env GOPATH)
install:
	go mod download
generate:
	protoc \
  -I$(GOPATH)/src \
  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:protobuf \
  --swagger_out=logtostderr=true:protobuf \
  --grpc-gateway_out=logtostderr=true:protobuf \
  --proto_path protobuf service.proto
	
	cp protobuf/service.swagger.json  www/swagger.json

build:
	go build  cmd/server/grpcServer.go
	go build  cmd/client/grpcClient.go
	go build  cmd/server/httpServer.go

rungrpc:
	go run  cmd/server/grpcServer.go

runhttp:	
	go run  cmd/server/httpServer.go

test:
	go test -v ./...

clientCli:
	./grpcClient.exe help