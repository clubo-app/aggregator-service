FROM golang:1.18.1-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY ./services/aggregator ./services/aggregator
COPY ./packages ./packages

EXPOSE 8081

WORKDIR /app/services/aggregator

RUN go build -o aggregator

ENTRYPOINT ./aggregator
