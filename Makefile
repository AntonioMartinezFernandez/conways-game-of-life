.PHONY: test build-new-board build play

test:
	go test ./...

build-new-board:
	go build -o bin/newboard ./cmd/newboard/main.go

build:
	go build -o bin/gameoflife ./cmd/play/main.go

play:
	go run ./cmd/play/main.go
