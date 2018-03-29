all: build test testjson

# Single targets
build:
	vgo build
test:
	vgo test -v -cover ./... | ./gorgeous
testjson:
	vgo test -v -cover -json ./... | ./gorgeous -json
