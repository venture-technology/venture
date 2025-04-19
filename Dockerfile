# Stage 1: Build the Go binary
FROM golang:1.24.2-alpine AS builder

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/api/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app ./

RUN mkdir -p /app/config /app/database

EXPOSE 9999

CMD ["/app/main"]
