# Stage 1: Build binary
FROM golang:tip-alpine3.22 AS builder

WORKDIR /app

# Install templ CLI
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate templ files (ini yang penting!)
RUN templ generate

# Build the application
RUN go build -o server .

# Stage 2: Final minimal image
FROM alpine:latest

WORKDIR /app

# Copy binary dan folder static dari stage build
COPY --from=builder /app/server .
COPY --from=builder /app/static ./static

# Jalankan binary
CMD ["./server"]