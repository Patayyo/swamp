FROM golang:1.21-alpine

WORKDIR /opt/app/api

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air"]