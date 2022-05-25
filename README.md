# API Gateway

## Generating Cards microsservice pb files
```
protoc --proto_path=proto --go_out=card/src/grpc/pb --go_opt=paths=source_relative --go-grpc_out=./card/src/grpc/pb --go-grpc_opt=paths=source_relative  cards.proto
```