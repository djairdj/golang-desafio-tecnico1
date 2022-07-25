install-proto-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

proto:
	@protoc --go_out=./pkg --go-grpc_out=require_unimplemented_servers=false:./pkg proto/*.proto
# no windows por definição:
#protoc --go_out=.\pkg --go-grpc_out=.\pkg proto\*.proto

install-grpc-clients:
	go install github.com/ktr0731/evans@latest
#go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

refresh-services:
	docker-compose down
	docker build -t store-server .
	docker-compose up -d

.PHONY: proto