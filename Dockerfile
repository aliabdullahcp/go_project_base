FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

# ENTRYPOINT go run commands/runserver.go

# For hot reloading
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build commands/runserver.go" --command=./runserver
