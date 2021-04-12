generate-proto:
	protoc --proto_path=./proto/ --proto_path=third_party --go_out=plugins=grpc:pkg/proto t-service.proto

run-service:
	go run cmd/service/main.go 

run-rest-server:
	go run cmd/rest-server/main.go

run-ws-server:
	go run cmd/ws-server/main.go
