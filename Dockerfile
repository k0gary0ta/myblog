# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app

EXPOSE 8080

RUN go install github.com/cosmtrek/air@v1.27.3

CMD ["air", "-c", ".air.toml"]
