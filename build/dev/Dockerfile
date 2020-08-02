FROM golang:1.13

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

#ENTRYPOINT CompileDaemon --build="go build commands/runserver.go" --command=./runserver
ENTRYPOINT CompileDaemon --build="go build -o ./server" --command=./server