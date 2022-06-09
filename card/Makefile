.PHONY: build run reflex

build:
	go build -o ./build/api ./src/cmd/api/main.go

build-protobuf:
	sh ./scripts/build-proto.sh

run-prd: build
	./build/api

install-reflex:
	go install github.com/cespare/reflex@latest

run-dev:
	reflex --start-service -r '\.go$$' go run ./src/cmd/api/main.go
