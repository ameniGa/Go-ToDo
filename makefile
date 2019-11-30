install:
	go get
generate:
	protoc \
  -I/usr/local/include \
  -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:proto \
  --swagger_out=logtostderr=true:api/proto \
  --grpc-gateway_out=logtostderr=true:api/proto \
  --proto_path api/proto service.proto
	
	cp api/swagger/service.swagger.json  www/swagger.json

build:
	go build  cmd/grpcserver/main.go
	go build  cmd/grpcclient/main.go
	go build  cmd/httpserver/main.go


rungrpc:
	go run  cmd/grpcserver/main.go
runhttp:	
	go run  cmd/httpserver/main.go  


runclient:  
	go run  cmd/grpcclient/main.go


test:
	go test api/reposImp

