format:
	go fmt ./...

test: format
	go test ./...

build:
	go build .
