.PHONY:clean
.DEFAULT_GOAL := all

all: build-amusingplutoserv

build-amusingplutoserv:
	go build -o ./bin/amusingplutoserv ./cmd/amusingplutoserv/main.go

docker-build:
	docker build -t plutoserv:0.0.01 -f Dockerfile/Dockerfile .

clean:
	rm -f ./bin/*