#!/bin/bash

PROTO_PATH=../proto
GENERATION_PATH=./src/adapters/grpc/pb

rm $GENERATION_PATH/*

protoc --proto_path=${PROTO_PATH}  \
  --go_out=${GENERATION_PATH} \
  --go_opt=paths=source_relative \
  --go-grpc_out=${GENERATION_PATH} \
  --go-grpc_opt=paths=source_relative \
  ${PROTO_PATH}/*.proto