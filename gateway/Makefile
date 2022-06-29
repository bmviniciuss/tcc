.PHONY: build run reflex

clean:
	rm -rf ./dist

build:
	npm run build

build-protobuf:
	sh ./scripts/build-proto.sh

run-prd: build
	node ./dist/application/index.js

dev:
	npm run local:start
