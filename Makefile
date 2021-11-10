.PHONY:clean
.DEFAULT_GOAL := all

all: build-amusingplutoserv build-amusingcharonserv
docker: docker-build

build-amusingplutoserv:
	CGO_ENABLED=0 go build -o ./bin/amusingplutoserv ./cmd/amusingplutoserv/main.go

build-amusingcharonserv:
	CGO_ENABLED=0 go build -o ./bin/amusingcharonserv ./cmd/amusingcharonserv/main.go

clean:
	rm -f ./bin/*