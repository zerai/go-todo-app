FROM golang:1.13

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go build

ENTRYPOINT ./go-todo-app