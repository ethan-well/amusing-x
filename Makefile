.DEFAULT_GOAL := all

all: build-go

build-go: build-amusingapiserv build-amusinguserserv

build-amusingapiserv:
	go build -o ./bin/amusingapiserv cmd/amusingapiserv/main.go
build-amusinguserserv:
	go build -o ./bin/amusinguserserv cmd/amusinguserserv/main.go