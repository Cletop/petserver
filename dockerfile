FROM golang:1.17.2

WORKDIR /petserver
COPY . .

RUN go mod download
