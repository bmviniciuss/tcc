build:
	npm run build

build-protobuf:
	sh ./scripts/build-proto.sh

prd: build
	node ./dist/application/index.js

dev:
	npm run local:start
