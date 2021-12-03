format:
	go fmt ./...

test: format
	go test ./...

build:
	go build .

up-components:
	docker-compose up

down-components:
	docker-compose down