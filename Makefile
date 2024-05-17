.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY: fmt

vet: fmt
	go vet ./...
.PHONY: vet

build: vet
	go build -o ./dist ./cmd/chip8
.PHONY: build

run: vet
	go run ./cmd/chip8
.PHONY: run
