.DEFAULT_GOAL := all

all: build-plutos

build-plutos:
	go build -o ./bin/pluto ./cmd/pluto/main.go

docker-build:
	docker build -t plutoserv:0.0.01 -f Dockerfile/Dockerfile .

clean:
	rm -f ./bin/*