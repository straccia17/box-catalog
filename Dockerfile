# syntax=docker/dockerfile:1

## Build
FROM golang:1.18.3-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /box-catalog-api

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /box-catalog-api /box-catalog-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/box-catalog-api"]