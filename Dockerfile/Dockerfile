# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN ls /app

RUN go build -o /docker-gs-ping /app/cmd/amusingplutoserv/main.go

EXPOSE 8080

ENTRYPOINT exec /docker-gs-ping -cf=/app/cmd/amusingplutoserv/service_testing.conf