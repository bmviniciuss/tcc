#!/bin/bash

PROTO_FILES_PATH=../proto
GENERATION_PATH=./src/adapters/grpc/pb

# Generate JavaScript code
npx grpc_tools_node_protoc \
	--js_out=import_style=commonjs,binary:${GENERATION_PATH} \
	--grpc_out=${GENERATION_PATH} \
	--plugin=proto-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin \
	--proto_path=${PROTO_FILES_PATH} \
	${PROTO_FILES_PATH}/*.proto

# Generate TypeScript typins (d.ts)
npx grpc_tools_node_protoc \
	--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
	--ts_out=${GENERATION_PATH} \
	--proto_path=${PROTO_FILES_PATH}  \
	${PROTO_FILES_PATH}/*.proto