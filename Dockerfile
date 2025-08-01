# Stage 1: Build binary
FROM golang:tip-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server .

# Stage 2: Final minimal image
FROM alpine:latest

WORKDIR /app

# Salin binary dan folder static dari stage build
COPY --from=builder /app/server .
COPY --from=builder /app/static ./static

# Jalankan binary
CMD ["./server"]
