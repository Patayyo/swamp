FROM golang:1.22-alpine

WORKDIR /opt/app/api/

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

run apk add --no-cache make
CMD ["air"]