# Stage 1: Build binary
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache bash

WORKDIR /app

# Copy go mod dan sum (cache friendly)
COPY go.mod go.sum ./
ENV GOPROXY=https://proxy.golang.org,direct
RUN go mod download

# Copy semua source code (tanpa .env)
COPY . .

# Build binary
RUN go build -ldflags="-s -w" -o main ./main.go

# Stage 2: Clean final image
FROM alpine:latest

RUN apk add --no-cache bash curl

WORKDIR /app

# Copy hasil build
COPY --from=builder /app/main .

# ⚠️ Jangan copy .env ke image
# karena nanti kita mount dari host lewat docker-compose

# Jalankan program dengan delay agar Besu siap dulu
CMD sh -c '\
  echo "⏳ Warming up Besu selama 15 detik..."; \
  sleep 15; \
  echo "🚀 Starting main application..."; \
  ./main'
