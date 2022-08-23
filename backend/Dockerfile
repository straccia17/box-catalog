# syntax=docker/dockerfile:1

## Build
FROM golang:1.18.3-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o box-catalog-api

## Deploy
FROM alpine

WORKDIR /app

COPY --from=build /app/box-catalog-api ./

EXPOSE 8080

CMD [ "/app/box-catalog-api" ]