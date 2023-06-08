.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: protoc
protoc:
	#protoc -I=./proto --go_out=. ./*.proto
	protoc -I=./protos --go_out=.  --go-grpc_out=. ./protos/*.proto

.PHONY: wire
wire:
	wire