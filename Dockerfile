# Stage 1: Build the Go binary
FROM golang:1.24.4-alpine AS builder

# Install required tools (for swag if needed)
RUN apk add --no-cache git

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum before code to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Install swag and generate Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN /go/bin/swag init

# Build the binary
RUN go build -o genomic-api .

# Stage 2: Create lightweight runtime container
FROM alpine:latest

# Install CA certificates (needed for HTTPS/TLS connections to DB)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/genomic-api .

# Copy Swagger docs (if used)
COPY --from=builder /app/docs ./docs

# Expose API port
EXPOSE 8080

# Run the binary
CMD ["./genomic-api"]
