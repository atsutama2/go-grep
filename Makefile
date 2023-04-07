.PHONY: build vet

BINARY_NAME=gg

build: vet
	go build -o $(BINARY_NAME)
	sudo mv gg /usr/local/bin

vet:
	go vet ./...
