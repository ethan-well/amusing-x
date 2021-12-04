FROM golang:1.17

#RUN apk add build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /pluto /app/cmd/pluto/main.go

ENTRYPOINT exec /pluto -cf=/app/cmd/pluto/config.yaml