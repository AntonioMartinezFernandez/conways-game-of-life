.PHONY: run test build

run:
	go run ./cmd/main.go

test:
	go test ./...

build:
	go build -o bin/gameoflife ./cmd/main.go
