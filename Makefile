.PHONY: build vet

BINARY_NAME=gg

build:
	go build -o $(BINARY_NAME) ./cmd/gg

vet:
	go vet ./...
