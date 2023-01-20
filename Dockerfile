FROM golang:alpine

RUN apk add --update --no-cache ca-certificates git

COPY . /go/helloworld
COPY go.mod /go/helloworld
WORKDIR /go/helloworld

RUN go mod download
COPY . .

EXPOSE 8080
CMD ["go", "run", "main.go"]