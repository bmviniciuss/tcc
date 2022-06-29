.PHONY: build run reflex

build:
	go build -o ./build/api ./src/cmd/api/main.go

build-protobuf:
	sh ./scripts/build-proto.sh

prd: build
	./build/api

dev:
	reflex --start-service -r '\.go$$' go run ./src/cmd/api/main.go
