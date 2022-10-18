# syntax=docker/dockerfile:1

# TO BUILD:
# docker build --tag tiny-service:v1.0 .
# TO RUN:
# docker run --publish 9002:9002 -e TINY_SERVICE_HOST=0.0.0.0 -e TINY_SERVICE_INSTANCE_NAME=instance1 -t tiny-service:v1.0

# image available at:
# https://hub.docker.com/repository/docker/stubin87/tiny-service

FROM golang:1.19.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg

RUN go build -o /bin/api ./cmd/service/main.go

EXPOSE 9002

CMD [ "/bin/api" ]
