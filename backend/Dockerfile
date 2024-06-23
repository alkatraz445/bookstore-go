FROM golang:1.22 AS build-stage

WORKDIR /app
COPY / ./
COPY go.mod go.sum ./

RUN go mod download

RUN go build -o /bookstore-go

# Koniec kompilacji

FROM ubuntu:24.04 AS run-stage

WORKDIR /
COPY --from=build-stage /bookstore-go /bookstore-go

EXPOSE 8080

ENTRYPOINT [ "/bookstore-go" ]