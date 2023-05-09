.PHONY: init
init:
	go install github.com/golang/protobuf@latest
	go install google.golang.org/genproto@latest
	go install google.golang.org/grpc@latest
	go install google.golang.org/protobuf@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: protoc
protoc:
	#protoc -I=./proto --go_out=. ./*.proto
	protoc -I=./protos  --go_out=plugins=grpc:. ./protos/*.proto

.PHONY: wire
wire:
	wire