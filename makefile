.DEFAULT_GOAL := all

all: build-plutos

build-plutos:
	go build -o ./bin/callisto ./cmd/callisto/main.go
	go build -o ./bin/charon ./cmd/charon/main.go
	go build -o ./bin/europa ./cmd/europa/main.go
	go build -o ./bin/user ./cmd/user/main.go
	go build -o ./bin/pluto ./cmd/pluto/main.go

docker-build:
	docker build -t plutoserv:0.0.01 -f Dockerfile/Dockerfile .

clean:
	rm -f ./bin/*