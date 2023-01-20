FROM golang:1.19.1-alpine
COPY . /go/helloworld
WORKDIR /go/helloworld
