# Stage 1: Build the Go binary
FROM golang:1.22.5-alpine AS builder

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY . .

# Build the binary from the specific main.go path
RUN go mod download && go build -o main ./cmd/api/server

# Stage 2: Run the application
FROM alpine:latest

WORKDIR /app

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/main .

# Create directories for config and database
RUN mkdir -p /app/config /app/database

EXPOSE 9999

CMD ["/app/main"]
