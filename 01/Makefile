.DEFAULT_GOAL := build

.PHONY: fmt vet build clean

fmt:
	go fmt main.go

vet: fmt
	go vet main.go

build: vet
	go build

clean:
	go clean