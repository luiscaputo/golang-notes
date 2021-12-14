FROM golang:alpine as builders

WORKDIR /go/src/

COPY . .

RUN ["/go/src/server"]