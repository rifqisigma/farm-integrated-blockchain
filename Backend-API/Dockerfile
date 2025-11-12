# Stage 1: Build binary
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache bash

WORKDIR /app

# Copy go mod dan sum (cache friendly)
COPY go.mod go.sum ./
ENV GOPROXY=https://proxy.golang.org,direct
RUN go mod download

# Copy semua source code
COPY . .

# Build binary
RUN go build -ldflags="-s -w" -o migrate ./cmd/migrate/main.go
RUN go build -ldflags="-s -w" -o main ./main.go

# Stage 2: Clean final image
FROM alpine:latest

# Install bash kalau tetap pakai CMD shell
RUN apk add --no-cache bash

WORKDIR /app

# Copy hanya hasil build
COPY --from=builder /app/migrate .
COPY --from=builder /app/main .


CMD ["sh", "-c", "sleep 20 && ./migrate && ./main"]
