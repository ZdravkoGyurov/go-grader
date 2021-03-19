FROM golang:alpine3.13

WORKDIR /app

COPY . .

RUN apk update && apk upgrade && apk add --no-cache git docker-cli

EXPOSE 8080

ENTRYPOINT go run cmd/go-grader/main.go
