FROM golang:alpine3.13

ARG gitUser=ZdravkoGyurov
ARG gitRepo=docker-tests
ARG assignment=assignment1
ENV ASSIGNMENT ${assignment}

WORKDIR /app

COPY ./go.mod /app/go.mod
COPY ./tests/${assignment}/* /app/${assignment}/

RUN apk update && apk upgrade && apk add --no-cache git

RUN git clone https://github.com/${gitUser}/${gitRepo}.git
RUN mv ${gitRepo}/${assignment}/* /app/${assignment}/

ENTRYPOINT go test /app/$ASSIGNMENT/calc_test.go
