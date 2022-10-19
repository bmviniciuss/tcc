.PHONY: build run-prd

build:
	go build -o ./build/api ./src/cmd/api/main.go

compile-protobuf:
	sh ./scripts/build-proto.sh

prd: build
	./build/api

install-reflex:
	go install github.com/cespare/reflex@latest

dev:
	reflex --start-service -r '\.go$$' go run ./src/cmd/api/main.go

rm-vol:
	docker volume rm card_db_data
