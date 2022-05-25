.PHONY: build run reflex

build:
	go build -o ./build/api ./src/cmd/api/main.go

build-protobuf:
	sh ./scripts/build-proto.sh

run: build
	./build/api

reflex:
	reflex --start-service -r '\.go$$' make run