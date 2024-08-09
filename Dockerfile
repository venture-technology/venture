# Stage 1: Build the Go binary
FROM golang:1.22.5-alpine AS builder

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY . .

RUN go mod download && go build -o main .

# Stage 2: Run the application
FROM alpine:latest

WORKDIR /app

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/main .

# Copy config.yaml from /config to /app/config/
COPY config/config.yaml /app/config/config.yaml

# Copy all .sql files from /database to /app/database/
COPY database/*.sql /app/database/

EXPOSE 8888

CMD ["./main"]
