.DEFAULT_GOAL := all

all: build-go

build-go: build-amusingapiserv

build-amusingapiserv:
	go build -o ./bin/amusingapiserv cmd/amusingapiserv/main.go
