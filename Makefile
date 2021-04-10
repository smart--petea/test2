generate-proto:
	protoc --proto_path=./proto/ --proto_path=third_party --go_out=plugins=grpc:pkg/proto t-service.proto

run-service:
	go run cmd/service/main.go -grpc-port=9090 -db-host=127.0.0.1:3306 -db-user=user -db-password=password -db-schema=rpc

run-rest-server:
	go run cmd/rest-server/main.go -server=127.0.0.1:9090
