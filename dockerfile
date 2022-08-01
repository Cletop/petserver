FROM golang:1.18.2

WORKDIR /pet-server
COPY . .

RUN go mod download
