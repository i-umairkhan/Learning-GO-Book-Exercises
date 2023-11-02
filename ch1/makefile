.DEFAULT_GO := build

fmt:
	go fmt hello.go
.PHONY:fmt

lint: fmt
	golint hello.go
.PHONY:lint

vet: fmt
	go vet hello.go
.PHONY:vet

build: vet
	go build hello.go
.PHONY:build
