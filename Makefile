.DEFAULT_GOAL := all

all: build-go

build-go: build-amusingplutoserv

build-amusingplutoserv:
	go build -o ./bin/amusingplutoserv ./cmd/amusingplutoserv/main.go