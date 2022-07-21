install-proto-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

proto:
	# no windows por definição:
	protoc --go_out=.\pkg --go-grpc_out=.\pkg proto\*.proto

evans:
	go install github.com/ktr0731/evans@latest

.PHONY: proto