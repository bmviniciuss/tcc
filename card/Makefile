.PHONY: build run reflex

build:
	go build -o ./build/api ./src/cmd/api/main.go

run: build
	./build/api

reflex:
	reflex --start-service -r '\.go$$' make run