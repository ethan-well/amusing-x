FROM golang:1.17 AS allFile

#RUN apk add build-base

WORKDIR /src
ARG SERVER_NAME
ENV SERVER_NAME ${SERVER_NAME}
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/${SERVER_NAME} /src/cmd/${SERVER_NAME}/main.go

FROM golang:1.17
WORKDIR /app
ARG SERVER_NAME
ENV SERVER_NAME ${SERVER_NAME}
COPY --from=allFile /app/${SERVER_NAME} /app/${SERVER_NAME}
COPY --from=allFile /src/cmd/${SERVER_NAME}/service_testing.conf /app/config/${SERVER_NAME}/service_testing.conf

ENTRYPOINT exec /app/${SERVER_NAME} -cf=/app/config/${SERVER_NAME}/service_testing.conf