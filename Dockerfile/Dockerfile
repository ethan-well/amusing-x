FROM golang:1.17

#RUN apk add build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /amusingplutoserv /app/cmd/amusingplutoserv/main.go

ENTRYPOINT exec /amusingplutoserv -cf=/app/cmd/amusingplutoserv/service_testing.conf