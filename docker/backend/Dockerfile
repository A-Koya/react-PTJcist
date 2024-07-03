FROM golang:1.22.4-alpine3.20

RUN apk update && apk add bash

WORKDIR /app

COPY ./backend/ .

RUN go install
