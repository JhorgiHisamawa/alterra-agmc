FROM golang:latest

WORKDIR /book-api

COPY go.mod go.sum ./

RUN go mod download

COPY . /book-api/

RUN go build

CMD ["./api"] 